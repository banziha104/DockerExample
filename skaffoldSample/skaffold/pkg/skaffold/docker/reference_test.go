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

	"github.com/GoogleContainerTools/skaffold/testutil"
)

func TestParseReference(t *testing.T) {
	var tests = []struct {
		description            string
		image                  string
		expectedName           string
		expectedFullyQualified bool
	}{
		{
			description:            "port and tag",
			image:                  "host:1234/user/container:tag",
			expectedName:           "host:1234/user/container",
			expectedFullyQualified: true,
		},
		{
			description:            "port",
			image:                  "host:1234/user/container",
			expectedName:           "host:1234/user/container",
			expectedFullyQualified: false,
		},
		{
			description:            "tag",
			image:                  "host/user/container:tag",
			expectedName:           "host/user/container",
			expectedFullyQualified: true,
		},
		{
			description:            "latest",
			image:                  "host/user/container:latest",
			expectedName:           "host/user/container",
			expectedFullyQualified: false,
		},
		{
			description:            "digest",
			image:                  "gcr.io/k8s-skaffold/example@sha256:81daf011d63b68cfa514ddab7741a1adddd59d3264118dfb0fd9266328bb8883",
			expectedName:           "gcr.io/k8s-skaffold/example",
			expectedFullyQualified: true,
		},
		{
			description:            "docker library",
			image:                  "nginx:latest",
			expectedName:           "nginx",
			expectedFullyQualified: false,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			parsed, err := ParseReference(test.image)

			testutil.CheckErrorAndDeepEqual(t, false, err, test.expectedName, parsed.BaseName)
			testutil.CheckErrorAndDeepEqual(t, false, err, test.expectedFullyQualified, parsed.FullyQualified)
		})
	}
}
