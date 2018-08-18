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
	"testing"

	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/schema/v1alpha2"
	"github.com/GoogleContainerTools/skaffold/testutil"
)

func TestApplyProfiles(t *testing.T) {
	tests := []struct {
		description string
		config      SkaffoldConfig
		profile     string
		expected    SkaffoldConfig
		shouldErr   bool
	}{
		{
			description: "unknown profile",
			config:      SkaffoldConfig{},
			profile:     "profile",
			shouldErr:   true,
		},
		{
			description: "build type",
			profile:     "profile",
			config: SkaffoldConfig{
				Build: v1alpha2.BuildConfig{
					Artifacts: []*v1alpha2.Artifact{
						{ImageName: "image"},
					},
					BuildType: v1alpha2.BuildType{
						LocalBuild: &v1alpha2.LocalBuild{},
					},
				},
				Deploy: v1alpha2.DeployConfig{},
				Profiles: []v1alpha2.Profile{
					{
						Name: "profile",
						Build: v1alpha2.BuildConfig{
							BuildType: v1alpha2.BuildType{
								GoogleCloudBuild: &v1alpha2.GoogleCloudBuild{},
							},
						},
					},
				},
			},
			expected: SkaffoldConfig{
				Build: v1alpha2.BuildConfig{
					Artifacts: []*v1alpha2.Artifact{
						{
							ImageName: "image",
							Workspace: ".",
							ArtifactType: v1alpha2.ArtifactType{
								DockerArtifact: &v1alpha2.DockerArtifact{
									DockerfilePath: "Dockerfile",
								},
							},
						},
					},
					BuildType: v1alpha2.BuildType{
						GoogleCloudBuild: &v1alpha2.GoogleCloudBuild{},
					},
					TagPolicy: v1alpha2.TagPolicy{
						GitTagger: &v1alpha2.GitTagger{},
					},
				},
				Deploy: v1alpha2.DeployConfig{},
			},
		},
		{
			description: "tag policy",
			profile:     "dev",
			config: SkaffoldConfig{
				Build: v1alpha2.BuildConfig{
					Artifacts: []*v1alpha2.Artifact{
						{ImageName: "image"},
					},
					TagPolicy: v1alpha2.TagPolicy{GitTagger: &v1alpha2.GitTagger{}},
				},
				Deploy: v1alpha2.DeployConfig{},
				Profiles: []v1alpha2.Profile{
					{
						Name: "dev",
						Build: v1alpha2.BuildConfig{
							TagPolicy: v1alpha2.TagPolicy{ShaTagger: &v1alpha2.ShaTagger{}},
						},
					},
				},
			},
			expected: SkaffoldConfig{
				Build: v1alpha2.BuildConfig{
					Artifacts: []*v1alpha2.Artifact{
						{
							ImageName: "image",
							Workspace: ".",
							ArtifactType: v1alpha2.ArtifactType{
								DockerArtifact: &v1alpha2.DockerArtifact{
									DockerfilePath: "Dockerfile",
								},
							},
						},
					},
					TagPolicy: v1alpha2.TagPolicy{ShaTagger: &v1alpha2.ShaTagger{}},
					BuildType: v1alpha2.BuildType{
						LocalBuild: &v1alpha2.LocalBuild{},
					},
				},
				Deploy: v1alpha2.DeployConfig{},
			},
		},
		{
			description: "artifacts",
			profile:     "profile",
			config: SkaffoldConfig{
				Build: v1alpha2.BuildConfig{
					Artifacts: []*v1alpha2.Artifact{
						{ImageName: "image"},
					},
					TagPolicy: v1alpha2.TagPolicy{GitTagger: &v1alpha2.GitTagger{}},
				},
				Deploy: v1alpha2.DeployConfig{},
				Profiles: []v1alpha2.Profile{
					{
						Name: "profile",
						Build: v1alpha2.BuildConfig{
							Artifacts: []*v1alpha2.Artifact{
								{ImageName: "image"},
								{ImageName: "imageProd"},
							},
						},
					},
				},
			},
			expected: SkaffoldConfig{
				Build: v1alpha2.BuildConfig{
					Artifacts: []*v1alpha2.Artifact{
						{
							ImageName: "image",
							Workspace: ".",
							ArtifactType: v1alpha2.ArtifactType{
								DockerArtifact: &v1alpha2.DockerArtifact{
									DockerfilePath: "Dockerfile",
								},
							},
						},
						{
							ImageName: "imageProd",
							Workspace: ".",
							ArtifactType: v1alpha2.ArtifactType{
								DockerArtifact: &v1alpha2.DockerArtifact{
									DockerfilePath: "Dockerfile",
								},
							},
						},
					},
					TagPolicy: v1alpha2.TagPolicy{GitTagger: &v1alpha2.GitTagger{}},
					BuildType: v1alpha2.BuildType{
						LocalBuild: &v1alpha2.LocalBuild{},
					},
				},
				Deploy: v1alpha2.DeployConfig{},
			},
		},
		{
			description: "deploy",
			profile:     "profile",
			config: SkaffoldConfig{
				Build: v1alpha2.BuildConfig{},
				Deploy: v1alpha2.DeployConfig{
					DeployType: v1alpha2.DeployType{
						KubectlDeploy: &v1alpha2.KubectlDeploy{},
					},
				},
				Profiles: []v1alpha2.Profile{
					{
						Name: "profile",
						Deploy: v1alpha2.DeployConfig{
							DeployType: v1alpha2.DeployType{
								HelmDeploy: &v1alpha2.HelmDeploy{},
							},
						},
					},
				},
			},
			expected: SkaffoldConfig{
				Build: v1alpha2.BuildConfig{
					TagPolicy: v1alpha2.TagPolicy{
						GitTagger: &v1alpha2.GitTagger{},
					},
					BuildType: v1alpha2.BuildType{
						LocalBuild: &v1alpha2.LocalBuild{},
					},
				},
				Deploy: v1alpha2.DeployConfig{
					DeployType: v1alpha2.DeployType{
						HelmDeploy: &v1alpha2.HelmDeploy{},
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			err := test.config.ApplyProfiles([]string{test.profile})

			testutil.CheckErrorAndDeepEqual(t, test.shouldErr, err, test.expected, test.config)
		})
	}
}
