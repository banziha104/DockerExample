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

package local

import (
	"context"
	"fmt"

	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/constants"
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/docker"
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/schema/v1alpha2"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Builder uses the host docker daemon to build and tag the image.
type Builder struct {
	cfg *v1alpha2.LocalBuild

	api          docker.APIClient
	localCluster bool
	pushImages   bool
	kubeContext  string

	alreadyTagged map[string]string
}

// NewBuilder returns an new instance of a local Builder.
func NewBuilder(cfg *v1alpha2.LocalBuild, kubeContext string) (*Builder, error) {
	api, err := docker.NewAPIClient()
	if err != nil {
		return nil, errors.Wrap(err, "getting docker client")
	}

	localCluster := kubeContext == constants.DefaultMinikubeContext || kubeContext == constants.DefaultDockerForDesktopContext
	var pushImages bool
	if cfg.SkipPush == nil {
		logrus.Debugf("skipPush value not present. defaulting to cluster default %t (minikube=true, d4d=true, gke=false)", localCluster)
		pushImages = !localCluster
	} else {
		pushImages = !*cfg.SkipPush
	}

	return &Builder{
		cfg:          cfg,
		kubeContext:  kubeContext,
		api:          api,
		localCluster: localCluster,
		pushImages:   pushImages,
	}, nil
}

// Labels are labels specific to local builder.
func (b *Builder) Labels() map[string]string {
	labels := map[string]string{
		constants.Labels.Builder: "local",
	}

	v, err := b.api.ServerVersion(context.Background())
	if err == nil {
		labels[constants.Labels.DockerAPIVersion] = fmt.Sprintf("%v", v.APIVersion)
	}

	return labels
}
