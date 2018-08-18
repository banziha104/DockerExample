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

package build

import (
	"context"
	"io"

	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/build/tag"
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/color"
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/schema/v1alpha2"
	"github.com/pkg/errors"
)

// InSequence builds a list of artifacts in sequence.
func InSequence(ctx context.Context, out io.Writer, tagger tag.Tagger, artifacts []*v1alpha2.Artifact, buildArtifact artifactBuilder) ([]Artifact, error) {
	var builds []Artifact

	for _, artifact := range artifacts {
		color.Default.Fprintf(out, "Building [%s]...\n", artifact.ImageName)

		tag, err := buildArtifact(ctx, out, tagger, artifact)
		if err != nil {
			return nil, errors.Wrapf(err, "building [%s]", artifact.ImageName)
		}

		builds = append(builds, Artifact{
			ImageName: artifact.ImageName,
			Tag:       tag,
		})
	}

	return builds, nil
}
