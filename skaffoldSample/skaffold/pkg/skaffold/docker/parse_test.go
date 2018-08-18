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
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/schema/v1alpha2"
	"github.com/GoogleContainerTools/skaffold/testutil"
	"github.com/google/go-containerregistry/pkg/v1"
)

const copyServerGo = `
FROM ubuntu:14.04
COPY server.go .
CMD server.go
`

const addNginx = `
FROM nginx
ADD nginx.conf /etc/nginx
CMD nginx
`

const multiCopy = `
FROM nginx
ADD test.conf /etc/test1
COPY test.conf /etc/test2
CMD nginx
`

const wildcards = `
FROM nginx
ADD *.go /tmp
`

const wildcardsMatchesNone = `
FROM nginx
ADD *.none /tmp
`

const oneWilcardMatchesNone = `
FROM nginx
ADD *.go *.none /tmp
`

const multiStageDockerfile = `
FROM golang:1.9.2
WORKDIR /go/src/github.com/r2d4/leeroy/
COPY worker.go .
RUN go build -o worker .

FROM gcr.io/distroless/base
WORKDIR /root/
COPY --from=0 /go/src/github.com/r2d4/leeroy .
CMD ["./worker"]
`

const envTest = `
FROM busybox
ENV foo bar
WORKDIR ${foo}   # WORKDIR /bar
COPY $foo /quux # COPY bar /quux
`

const copyDirectory = `
FROM nginx
ADD . /etc/
COPY ./file /etc/file
CMD nginx
`
const multiFileCopy = `
FROM ubuntu:14.04
COPY server.go file .
`

const remoteFileAdd = `
FROM ubuntu:14.04
ADD https://example.com/test /test
`

const contextDockerfile = `
FROM nginx
ADD nginx.conf /etc/nginx
COPY . /files
CMD nginx
`

// This has an ONBUILD instruction of "COPY . /go/src/app"
const onbuild = `
FROM golang:onbuild
`

const onbuildError = `
FROM noimage:latest
ADD ./file /etc/file
`

const copyServerGoBuildArg = `
FROM ubuntu:14.04
ARG FOO
COPY $FOO .
CMD $FOO
`

const copyServerGoBuildArgCurlyBraces = `
FROM ubuntu:14.04
ARG FOO
COPY ${FOO} .
CMD ${FOO}
`

const copyServerGoBuildArgExtraWhitespace = `
FROM ubuntu:14.04
ARG  FOO
COPY $FOO .
CMD $FOO
`

const copyServerGoBuildArgDefaultValue = `
FROM ubuntu:14.04
ARG FOO=server.go
COPY $FOO .
CMD $FOO
`

var fooArg = "server.go" // used for build args

var ImageConfigs = map[string]*v1.ConfigFile{
	"golang:onbuild": {
		Config: v1.Config{
			OnBuild: []string{
				"COPY . /go/src/app",
			},
		},
	},
	"golang:1.9.2":           {Config: v1.Config{}},
	"gcr.io/distroless/base": {Config: v1.Config{}},
	"ubuntu:14.04":           {Config: v1.Config{}},
	"nginx":                  {Config: v1.Config{}},
	"busybox":                {Config: v1.Config{}},
	"oneport": {
		Config: v1.Config{
			ExposedPorts: map[string]struct{}{
				"8000": {},
			},
		}},
	"severalports": {
		Config: v1.Config{
			ExposedPorts: map[string]struct{}{
				"8000":     {},
				"8001/tcp": {},
			},
		}},
}

func mockRetrieveImage(image string) (*v1.ConfigFile, error) {
	if cfg, ok := ImageConfigs[image]; ok {
		return cfg, nil
	}
	return nil, fmt.Errorf("No image found for %s", image)
}

func TestGetDependencies(t *testing.T) {
	var tests = []struct {
		description string
		dockerfile  string
		workspace   string
		ignore      string
		buildArgs   map[string]*string

		expected  []string
		badReader bool
		shouldErr bool
	}{
		{
			description: "copy dependency",
			dockerfile:  copyServerGo,
			workspace:   ".",
			expected:    []string{"Dockerfile", "server.go"},
		},
		{
			description: "add dependency",
			dockerfile:  addNginx,
			workspace:   "docker",
			expected:    []string{"Dockerfile", "nginx.conf"},
		},
		{
			description: "wildcards",
			dockerfile:  wildcards,
			workspace:   ".",
			expected:    []string{"Dockerfile", "server.go", "worker.go"},
		},
		{
			description: "wildcards matches none",
			dockerfile:  wildcardsMatchesNone,
			workspace:   ".",
			shouldErr:   true,
		},
		{
			description: "one wilcard matches none",
			dockerfile:  oneWilcardMatchesNone,
			workspace:   ".",
			expected:    []string{"Dockerfile", "server.go", "worker.go"},
		},
		{
			description: "bad read",
			badReader:   true,
			shouldErr:   true,
		},
		{
			// https://github.com/GoogleContainerTools/skaffold/issues/158
			description: "no dependencies on remote files",
			dockerfile:  remoteFileAdd,
			expected:    []string{"Dockerfile"},
		},
		{
			description: "multistage dockerfile",
			dockerfile:  multiStageDockerfile,
			workspace:   "",
			expected:    []string{"Dockerfile", "worker.go"},
		},
		{
			description: "copy twice",
			dockerfile:  multiCopy,
			workspace:   ".",
			expected:    []string{"Dockerfile", "test.conf"},
		},
		{
			description: "env test",
			dockerfile:  envTest,
			workspace:   ".",
			expected:    []string{"Dockerfile", "bar"},
		},
		{
			description: "multi file copy",
			dockerfile:  multiFileCopy,
			workspace:   ".",
			expected:    []string{"Dockerfile", "file", "server.go"},
		},
		{
			description: "dockerignore test",
			dockerfile:  copyDirectory,
			ignore:      "bar\ndocker/*",
			workspace:   ".",
			expected:    []string{"Dockerfile", "file", "server.go", "test.conf", "worker.go"},
		},
		{
			description: "dockerignore dockerfile",
			dockerfile:  copyServerGo,
			ignore:      "Dockerfile",
			workspace:   ".",
			expected:    []string{"Dockerfile", "server.go"},
		},
		{
			description: "dockerignore with non canonical workspace",
			dockerfile:  contextDockerfile,
			workspace:   "docker/../docker",
			ignore:      "bar\ndocker/*",
			expected:    []string{"Dockerfile", "nginx.conf"},
		},
		{
			description: "dockerignore with context in parent directory",
			dockerfile:  copyDirectory,
			workspace:   "docker/..",
			ignore:      "bar\ndocker/*\n*.go",
			expected:    []string{"Dockerfile", "file", "test.conf"},
		},
		{
			description: "onbuild test",
			dockerfile:  onbuild,
			workspace:   ".",
			expected:    []string{"Dockerfile", "bar", filepath.Join("docker", "bar"), filepath.Join("docker", "nginx.conf"), "file", "server.go", "test.conf", "worker.go"},
		},
		{
			description: "onbuild error",
			dockerfile:  onbuildError,
			workspace:   ".",
			expected:    []string{"Dockerfile", "file"},
		},
		{
			description: "build args",
			dockerfile:  copyServerGoBuildArg,
			workspace:   ".",
			buildArgs:   map[string]*string{"FOO": &fooArg},
			expected:    []string{"Dockerfile", "server.go"},
		},
		{
			description: "build args with curly braces",
			dockerfile:  copyServerGoBuildArgCurlyBraces,
			workspace:   ".",
			buildArgs:   map[string]*string{"FOO": &fooArg},
			expected:    []string{"Dockerfile", "server.go"},
		},
		{
			description: "build args with extra whitespace",
			dockerfile:  copyServerGoBuildArgExtraWhitespace,
			workspace:   ".",
			buildArgs:   map[string]*string{"FOO": &fooArg},
			expected:    []string{"Dockerfile", "server.go"},
		},
		{
			description: "build args with default value and buildArgs unset",
			dockerfile:  copyServerGoBuildArgDefaultValue,
			workspace:   ".",
			expected:    []string{"Dockerfile", "server.go"},
		},
		{
			description: "build args with default value and buildArgs set",
			dockerfile:  copyServerGoBuildArgDefaultValue,
			workspace:   ".",
			buildArgs:   map[string]*string{"FOO": &fooArg},
			expected:    []string{"Dockerfile", "server.go"},
		},
	}

	RetrieveImage = mockRetrieveImage
	defer func() {
		RetrieveImage = retrieveImage
	}()

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			tmpDir, cleanup := testutil.TempDir(t)
			defer cleanup()

			os.MkdirAll(filepath.Join(tmpDir, "docker"), 0750)
			for _, file := range []string{"docker/nginx.conf", "docker/bar", "server.go", "test.conf", "worker.go", "bar", "file"} {
				ioutil.WriteFile(filepath.Join(tmpDir, file), []byte(""), 0644)
			}

			workspace := filepath.Join(tmpDir, test.workspace)
			if !test.badReader {
				ioutil.WriteFile(filepath.Join(workspace, "Dockerfile"), []byte(test.dockerfile), 0644)
			}

			if test.ignore != "" {
				ioutil.WriteFile(filepath.Join(workspace, ".dockerignore"), []byte(test.ignore), 0644)
			}

			deps, err := GetDependencies(workspace, &v1alpha2.DockerArtifact{
				BuildArgs:      test.buildArgs,
				DockerfilePath: "Dockerfile",
			})

			testutil.CheckErrorAndDeepEqual(t, test.shouldErr, err, test.expected, deps)
		})
	}
}
