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

package cmd

import (
	"io"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/color"
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/config"
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/schema"
	schemautil "github.com/GoogleContainerTools/skaffold/pkg/skaffold/schema/util"
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/util"
)

func NewCmdFix(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fix",
		Short: "Converts old skaffold.yaml to newest schema version",
		Run: func(cmd *cobra.Command, args []string) {
			contents, err := util.ReadConfiguration(filename)
			if err != nil {
				logrus.Errorf("fix: %s", err)
			}
			cfg, err := config.GetConfig(contents, false)
			if err != nil {
				logrus.Error(err)
				return
			}
			if cfg.GetVersion() == config.LatestVersion {
				color.Default.Fprintln(out, "config is already latest version")
				return
			}
			if err := runFix(out, cfg); err != nil {
				logrus.Errorf("fix: %s", err)
			}
		},
		Args: cobra.NoArgs,
	}
	AddFixFlags(cmd)
	return cmd
}

func runFix(out io.Writer, cfg schemautil.VersionedConfig) error {
	cfg, err := schema.RunTransform(cfg)
	if err != nil {
		return err
	}
	newCfg, err := yaml.Marshal(cfg)
	if err != nil {
		return errors.Wrap(err, "marshaling new config")
	}
	if overwrite {
		if err := ioutil.WriteFile(filename, newCfg, 0644); err != nil {
			return errors.Wrap(err, "writing config file")
		}
		color.Default.Fprintf(out, "New config at version %s generated and written to %s\n", cfg.GetVersion(), filename)
	} else {
		out.Write(newCfg)
	}
	return nil
}
