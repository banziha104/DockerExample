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

package docker

import (
	"testing"

	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/util"
	"github.com/GoogleContainerTools/skaffold/testutil"
)

func TestNewEnvClient(t *testing.T) {
	var tests = []struct {
		description string
		envs        map[string]string
		shouldErr   bool
	}{
		{
			description: "get client",
			envs: map[string]string{
				"DOCKER_HOST": "http://127.0.0.1:8080",
			},
		},
		{
			description: "invalid cert path",
			envs: map[string]string{
				"DOCKER_HOST":      "http://127.0.0.1:8080",
				"DOCKER_CERT_PATH": "invalid/cert/path",
			},
			shouldErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			unsetEnvs := testutil.SetEnvs(t, test.envs)
			_, err := newEnvAPIClient()
			testutil.CheckError(t, test.shouldErr, err)
			unsetEnvs(t)
		})
	}

}

func TestNewMinikubeImageAPIClient(t *testing.T) {
	var tests = []struct {
		description string
		cmd         util.Command

		expected  APIClient
		shouldErr bool
	}{
		{
			description: "correct client",
			cmd: testutil.NewFakeCmdOut("minikube docker-env --shell none", `DOCKER_TLS_VERIFY=1
DOCKER_HOST=http://127.0.0.1:8080
DOCKER_CERT_PATH=testdata
DOCKER_API_VERSION=1.23`, nil),
		},
		{
			description: "correct client",
			cmd: testutil.NewFakeCmdOut("minikube docker-env --shell none", `DOCKER_TLS_VERIFY=1
DOCKER_HOST=http://127.0.0.1:8080
DOCKER_CERT_PATH=bad/cert/path
DOCKER_API_VERSION=1.23`, nil),
			shouldErr: true,
		},
		{
			description: "missing host env, no error",
			cmd: testutil.NewFakeCmdOut("minikube docker-env --shell none", `DOCKER_TLS_VERIFY=1
DOCKER_CERT_PATH=testdata
DOCKER_API_VERSION=1.23`, nil),
		},
		{
			description: "missing version env, no error",
			cmd: testutil.NewFakeCmdOut("minikube docker-env --shell none", `DOCKER_TLS_VERIFY=1
DOCKER_HOST=http://127.0.0.1:8080
DOCKER_CERT_PATH=testdata`, nil),
		},
		{
			description: "missing version env, no error",
			cmd: testutil.NewFakeCmdOut("minikube docker-env --shell none", `DOCKER_TLS_VERIFY=1
DOCKER_HOST=badurl
DOCKER_CERT_PATH=testdata
DOCKER_API_VERSION=1.23`, nil),
			shouldErr: true,
		},
		{
			description: "bad env output, should fallback to host docker",
			cmd: testutil.NewFakeCmdOut("minikube docker-env --shell none", `DOCKER_TLS_VERIFY=1
DOCKER_HOST=http://127.0.0.1:8080=toomanyvalues
DOCKER_CERT_PATH=testdata
DOCKER_API_VERSION=1.23`, nil),
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			defer func(c util.Command) { util.DefaultExecCommand = c }(util.DefaultExecCommand)
			util.DefaultExecCommand = test.cmd

			_, err := newMinikubeAPIClient()
			testutil.CheckError(t, test.shouldErr, err)
		})
	}
}
