package slack

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ycanty/go-cli/json"
	"github.com/ycanty/go-cli/slack"
	"log"
)

const (
	configSlackChannelId = "slack.channel-id"
)

func newGetConversationHistoryCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "get-conversation-history",
		Short: "Get the conversation history in a channel",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			channelID := viper.GetString(configSlackChannelId)
			ch, err := slackApi().GetConversationHistory(slack.Channel{ID: channelID})
			if err != nil {
				return err
			}
			_ = json.PrintJSON(ch)
			return nil
		},
	}

	command.Flags().StringP("channel-id", "c", "", "Channel ID")

	if err := viper.BindPFlag(configSlackChannelId, command.Flag("channel-id")); err != nil {
		log.Fatal(err)
	}
	return command
}