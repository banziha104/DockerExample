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

package deploy

import (
	"context"
	"fmt"
	"io"

	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/build"
	"k8s.io/apimachinery/pkg/runtime"
)

// Artifact contains all information about a completed deployment
type Artifact struct {
	Obj       *runtime.Object
	Namespace string
}

// Deployer is the Deploy API of skaffold and responsible for deploying
// the build results to a Kubernetes cluster
type Deployer interface {
	Labels() map[string]string

	// Deploy should ensure that the build results are deployed to the Kubernetes
	// cluster.
	Deploy(context.Context, io.Writer, []build.Artifact) ([]Artifact, error)

	// Dependencies returns a list of files that the deployer depends on.
	// In dev mode, a redeploy will be triggered
	Dependencies() ([]string, error)

	// Cleanup deletes what was deployed by calling Deploy.
	Cleanup(context.Context, io.Writer) error
}

func joinTagsToBuildResult(builds []build.Artifact, params map[string]string) (map[string]build.Artifact, error) {
	imageToBuildResult := map[string]build.Artifact{}
	for _, build := range builds {
		imageToBuildResult[build.ImageName] = build
	}

	paramToBuildResult := map[string]build.Artifact{}
	for param, imageName := range params {
		build, ok := imageToBuildResult[imageName]
		if !ok {
			return nil, fmt.Errorf("No build present for %s", imageName)
		}
		paramToBuildResult[param] = build
	}
	return paramToBuildResult, nil
}
