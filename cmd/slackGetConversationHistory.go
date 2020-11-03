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
			afterMessage := cmd.Flags().Lookup("after-message").Value.String()
			since_date := cmd.Flags().Lookup("since").Value.String()
			if since_date != "" {
				thetime, err := time.Parse("2006-01-02", since_date)
				if err != nil {
					return err
				}
				afterMessage = fmt.Sprintf("%d", thetime.Unix())
			} else if afterMessage == "" && since_date == "" {
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
	command.Flags().StringP("after-message", "a", "", "Message ID")
	command.Flags().StringP("since", "s", "", "Since YYYY-MM-DD")

	if err := viper.BindPFlag(configSlackChannelId, command.Flag("channel-id")); err != nil {
		log.Fatal(err)
	}
	return command
}

func sinceDefault(days_ago int) int64 {
	thetime := time.Now().AddDate(0, 0, -days_ago).Unix() // since x days ago
	return thetime
}
