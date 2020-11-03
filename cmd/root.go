/*
Copyright © 2020 Yves Canty <46539628+ycanty@users.noreply.github.com>

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
	"fmt"
	"github.com/spf13/cobra"
	"github.com/ycanty/slack-stats/json"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string
var jsonPath string

func NewRootCmd() *cobra.Command {
	command := &cobra.Command{
		Short:        "Slack statistics",
		Long:         `Generate statistics reports of slack conversation histories`,
		SilenceUsage: true,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			json.SetJSONPath(jsonPath)
		},
	}

	command.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.slack-stats.yaml)")
	command.PersistentFlags().StringVar(&jsonPath, "jsonpath", "", "Filter json output with a JSONPath expression")

	command.AddCommand(NewSlackCommand())
	command.AddCommand(NewDBCommand())

	cobra.OnInitialize(initConfig)

	return command
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in current dir, then home dir
		viper.SetConfigName(".slack-stats")
		viper.AddConfigPath(".")
		viper.AddConfigPath(home)
	}

	viper.SetEnvPrefix("SLACK_STATS")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		// fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
