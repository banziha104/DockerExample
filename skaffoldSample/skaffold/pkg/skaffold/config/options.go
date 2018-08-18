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

package config

import (
	"strings"
)

// SkaffoldOptions are options that are set by command line arguments not included
// in the config file itself
type SkaffoldOptions struct {
	Cleanup      bool
	Notification bool
	Profiles     []string
	CustomTag    string
	Namespace    string
}

// Labels returns a map of labels to be applied to all deployed
// k8s objects during the duration of the run
func (opts *SkaffoldOptions) Labels() map[string]string {
	labels := map[string]string{}

	if opts.Cleanup {
		labels["cleanup"] = "true"
	}
	if opts.Namespace != "" {
		labels["namespace"] = opts.Namespace
	}
	if len(opts.Profiles) > 0 {
		labels["profiles"] = strings.Join(opts.Profiles, ",")
	}
	return labels
}
