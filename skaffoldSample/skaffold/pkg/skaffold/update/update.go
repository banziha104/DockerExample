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

package update

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/constants"
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/version"
	"github.com/blang/semver"
	"github.com/pkg/errors"
)

// IsUpdateCheckEnabled returns whether or not the update check is enabled
// It is true by default, but setting it to any other value than true will disable the check
func IsUpdateCheckEnabled() bool {
	// Don't perform a version check on dirty trees
	if version.Get().GitTreeState == "dirty" {
		return false
	}
	v := os.Getenv(constants.UpdateCheckEnvironmentVariable)
	if v == "" || strings.ToLower(v) == "true" {
		return true
	}
	return false
}

var latestVersionURL = fmt.Sprintf("https://storage.googleapis.com/skaffold/releases/latest/VERSION")

// GetLatestVersion uses a VERSION file stored on GCS to determine the latest released version of skaffold
func GetLatestVersion(ctx context.Context) (semver.Version, error) {
	resp, err := http.Get(latestVersionURL)
	if err != nil {
		return semver.Version{}, errors.Wrap(err, "getting latest version info from GCS")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return semver.Version{}, errors.Wrapf(err, "http %d, error: %s", resp.StatusCode, resp.Status)
	}
	versionBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return semver.Version{}, errors.Wrap(err, "reading version file from GCS")
	}
	v, err := version.ParseVersion(string(versionBytes))
	if err != nil {
		return semver.Version{}, errors.Wrap(err, "parsing latest version from GCS")
	}
	return v, nil
}
