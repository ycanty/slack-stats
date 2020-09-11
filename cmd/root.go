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
	"log"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/slack-go/slack"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:          "slack",
	Short:        "Slack CLI utilities",
	Long:         `Interact with Slack through the command line`,
	SilenceUsage: true,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.slack-cli.yaml)")
	rootCmd.PersistentFlags().String("token", "xyz", "Slack authentication token")

	if err := viper.BindPFlag("slack.token", rootCmd.PersistentFlags().Lookup("token")); err != nil {
		log.Fatal(err)
	}

	rootCmd.AddCommand(newFindChannelIDCommand())
	rootCmd.AddCommand(newGetConversationHistoryCommand())
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
		viper.SetConfigName(".slack-cli")
		viper.AddConfigPath(".")
		viper.AddConfigPath(home)
	}

	viper.SetEnvPrefix("SLACK_CLI")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		// fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

var __api *slack.Client

func Api() *slack.Client {
	if __api != nil {
		return __api
	}
	return slack.New(viper.GetString("slack.token"))
}