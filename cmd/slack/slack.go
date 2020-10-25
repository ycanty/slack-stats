package slack

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ycanty/slack-stats/slack"
	"log"
)

const (
	configSlackToken = "slack.token"
)

func NewSlackCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "slack",
		Short: "Interact with slack",
		Long:  ``,
	}

	command.AddCommand(newGetConversationHistoryCommand())
	command.AddCommand(newFindChannelIDCommand())
	command.AddCommand(newGetUserInfoCommand())

	command.PersistentFlags().String("token", "", "Slack authentication token")

	if err := viper.BindPFlag(configSlackToken, command.PersistentFlags().Lookup("token")); err != nil {
		log.Fatal(err)
	}

	return command
}

var cachedSlackApi *slack.Api

// slackApi is to be used by subcommands to get the slack API
func slackApi() *slack.Api {
	if cachedSlackApi != nil {
		cachedSlackApi = slack.NewApi(viper.GetString(configSlackToken))
	}
	return cachedSlackApi
}
