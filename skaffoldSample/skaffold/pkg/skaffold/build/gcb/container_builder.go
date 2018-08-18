/*
Copyright 2018 The Skaffold Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package gcb

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"time"

	cstorage "cloud.google.com/go/storage"
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/build"
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/build/tag"
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/color"
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/constants"
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/docker"
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/schema/v1alpha2"
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/util"
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/version"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2/google"
	cloudbuild "google.golang.org/api/cloudbuild/v1"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/iterator"
)

// Build builds a list of artifacts with GCB.
func (b *Builder) Build(ctx context.Context, out io.Writer, tagger tag.Tagger, artifacts []*v1alpha2.Artifact) ([]build.Artifact, error) {
	return build.InParallel(ctx, out, tagger, artifacts, b.buildArtifact)
}

func (b *Builder) buildArtifact(ctx context.Context, out io.Writer, tagger tag.Tagger, artifact *v1alpha2.Artifact) (string, error) {
	client, err := google.DefaultClient(ctx, cloudbuild.CloudPlatformScope)
	if err != nil {
		return "", errors.Wrap(err, "getting google client")
	}

	cbclient, err := cloudbuild.New(client)
	if err != nil {
		return "", errors.Wrap(err, "getting builder")
	}
	cbclient.UserAgent = version.UserAgent()

	c, err := cstorage.NewClient(ctx)
	if err != nil {
		return "", errors.Wrap(err, "getting cloud storage client")
	}
	defer c.Close()

	// need to format build args as strings to pass to container builder docker
	var buildArgs []string
	for k, v := range artifact.DockerArtifact.BuildArgs {
		if v != nil {
			buildArgs = append(buildArgs, []string{"--build-arg", fmt.Sprintf("%s=%s", k, *v)}...)
		}
	}
	logrus.Debugf("Build args: %s", buildArgs)

	projectID, err := b.guessProjectID(artifact)
	if err != nil {
		return "", errors.Wrap(err, "getting projectID")
	}

	cbBucket := fmt.Sprintf("%s%s", projectID, constants.GCSBucketSuffix)
	buildObject := fmt.Sprintf("source/%s-%s.tar.gz", projectID, util.RandomID())

	if err := b.createBucketIfNotExists(ctx, projectID, cbBucket); err != nil {
		return "", errors.Wrap(err, "creating bucket if not exists")
	}
	if err := b.checkBucketProjectCorrect(ctx, projectID, cbBucket); err != nil {
		return "", errors.Wrap(err, "checking bucket is in correct project")
	}

	color.Default.Fprintf(out, "Pushing code to gs://%s/%s\n", cbBucket, buildObject)
	if err := docker.UploadContextToGCS(ctx, artifact.Workspace, artifact.DockerArtifact, cbBucket, buildObject); err != nil {
		return "", errors.Wrap(err, "uploading source tarball")
	}

	args := append([]string{"build", "--tag", artifact.ImageName, "-f", artifact.DockerArtifact.DockerfilePath}, buildArgs...)
	args = append(args, ".")
	call := cbclient.Projects.Builds.Create(projectID, &cloudbuild.Build{
		LogsBucket: cbBucket,
		Source: &cloudbuild.Source{
			StorageSource: &cloudbuild.StorageSource{
				Bucket: cbBucket,
				Object: buildObject,
			},
		},
		Steps: []*cloudbuild.BuildStep{
			{
				Name: "gcr.io/cloud-builders/docker",
				Args: args,
			},
		},
		Images: []string{artifact.ImageName},
		Options: &cloudbuild.BuildOptions{
			DiskSizeGb:  b.DiskSizeGb,
			MachineType: b.MachineType,
		},
	})
	op, err := call.Context(ctx).Do()
	if err != nil {
		return "", errors.Wrap(err, "could not create build")
	}

	remoteID, err := getBuildID(op)
	if err != nil {
		return "", errors.Wrapf(err, "getting build ID from op")
	}
	logsObject := fmt.Sprintf("log-%s.txt", remoteID)
	color.Default.Fprintf(out, "Logs at available at \nhttps://console.cloud.google.com/m/cloudstorage/b/%s/o/%s\n", cbBucket, logsObject)
	var imageID string
	offset := int64(0)
watch:
	for {
		logrus.Debugf("current offset %d", offset)
		cb, err := cbclient.Projects.Builds.Get(projectID, remoteID).Do()
		if err != nil {
			return "", errors.Wrap(err, "getting build status")
		}

		r, err := b.getLogs(ctx, offset, cbBucket, logsObject)
		if err != nil {
			return "", errors.Wrap(err, "getting logs")
		}
		if r != nil {
			written, err := io.Copy(out, r)
			if err != nil {
				return "", errors.Wrap(err, "copying logs to stdout")
			}
			offset += written
			r.Close()
		}
		switch cb.Status {
		case StatusQueued, StatusWorking, StatusUnknown:
		case StatusSuccess:
			imageID, err = getImageID(cb)
			if err != nil {
				return "", errors.Wrap(err, "getting image id from finished build")
			}
			break watch
		case StatusFailure, StatusInternalError, StatusTimeout, StatusCancelled:
			return "", fmt.Errorf("cloud build failed: %s", cb.Status)
		default:
			return "", fmt.Errorf("unknown status: %s", cb.Status)
		}

		time.Sleep(RetryDelay)
	}

	if err := c.Bucket(cbBucket).Object(buildObject).Delete(ctx); err != nil {
		return "", errors.Wrap(err, "cleaning up source tar after build")
	}
	logrus.Infof("Deleted object %s", buildObject)
	builtTag := fmt.Sprintf("%s@%s", artifact.ImageName, imageID)
	logrus.Infof("Image built at %s", builtTag)

	newTag, err := tagger.GenerateFullyQualifiedImageName(artifact.Workspace, &tag.Options{
		ImageName: artifact.ImageName,
		Digest:    imageID,
	})

	if err != nil {
		return "", errors.Wrap(err, "generating tag")
	}

	if err := docker.AddTag(builtTag, newTag); err != nil {
		return "", errors.Wrap(err, "tagging image")
	}

	return newTag, nil
}

func getBuildID(op *cloudbuild.Operation) (string, error) {
	if op.Metadata == nil {
		return "", errors.New("missing Metadata in operation")
	}
	var buildMeta cloudbuild.BuildOperationMetadata
	if err := json.Unmarshal([]byte(op.Metadata), &buildMeta); err != nil {
		return "", err
	}
	if buildMeta.Build == nil {
		return "", errors.New("missing Build in operation metadata")
	}
	return buildMeta.Build.Id, nil
}

func getImageID(b *cloudbuild.Build) (string, error) {
	if b.Results == nil || len(b.Results.Images) == 0 {
		return "", errors.New("build failed")
	}
	return b.Results.Images[0].Digest, nil
}

func (b *Builder) getLogs(ctx context.Context, offset int64, bucket, objectName string) (io.ReadCloser, error) {
	c, err := cstorage.NewClient(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "getting storage client")
	}
	defer c.Close()

	r, err := c.Bucket(bucket).Object(objectName).NewRangeReader(ctx, offset, -1)
	if err != nil {
		if gerr, ok := err.(*googleapi.Error); ok {
			switch gerr.Code {
			case 404, 416, 429, 503:
				logrus.Debugf("Status Code: %d, %s", gerr.Code, gerr.Body)
				return nil, nil
			}
		}
		if err == cstorage.ErrObjectNotExist {
			logrus.Debugf("Logs for %s %s not uploaded yet...", bucket, objectName)
			return nil, nil
		}
		return nil, errors.Wrap(err, "unknown error")
	}
	return r, nil
}

func (b *Builder) checkBucketProjectCorrect(ctx context.Context, projectID, bucket string) error {
	c, err := cstorage.NewClient(ctx)
	if err != nil {
		return errors.Wrap(err, "getting storage client")
	}
	it := c.Buckets(ctx, projectID)
	// Set the prefix to the bucket we're looking for to only return that bucket and buckets with that prefix
	// that we'll filter further later on
	it.Prefix = bucket
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			return errors.Wrap(err, "bucket not found")
		}
		if err != nil {
			return errors.Wrap(err, "iterating over buckets")
		}
		// Since we can't filter on bucket name specifically, only prefix, we need to check equality here and not just prefix
		if attrs.Name == bucket {
			return nil
		}
	}
}

func (b *Builder) createBucketIfNotExists(ctx context.Context, projectID, bucket string) error {
	c, err := cstorage.NewClient(ctx)
	if err != nil {
		return errors.Wrap(err, "getting storage client")
	}
	defer c.Close()

	_, err = c.Bucket(bucket).Attrs(ctx)

	if err == nil {
		// Bucket exists
		return nil
	}

	if err != cstorage.ErrBucketNotExist {
		return errors.Wrapf(err, "getting bucket %s", bucket)
	}

	if err := c.Bucket(bucket).Create(ctx, projectID, &cstorage.BucketAttrs{
		Name: bucket,
	}); err != nil {
		return err
	}
	logrus.Debugf("Created bucket %s in %s", bucket, projectID)
	return nil
}
