package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ycanty/slack-stats/json"
	"github.com/ycanty/slack-stats/slack"
	"log"
	"time"
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
			afterMessage := cmd.Flags().Lookup("after").Value.String()
			if afterMessage == "" {
				thetime := time.Now().AddDate(0, 0, -7).Unix() // since 7 days ago
				afterMessage = fmt.Sprintf("%d", thetime)
			}
			ch, err := slackApi().GetConversationHistory(slack.Channel{ID: channelID}, afterMessage)
			if err != nil {
				return err
			}

			return json.PrintJSON(cmd.OutOrStdout(), ch)
		},
	}

	command.Flags().StringP("channel-id", "c", "", "Channel ID")
	command.Flags().StringP("after", "a", "", "Message ID")

	if err := viper.BindPFlag(configSlackChannelId, command.Flag("channel-id")); err != nil {
		log.Fatal(err)
	}
	return command
}