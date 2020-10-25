package slack

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ycanty/go-cli/json"
	"github.com/ycanty/go-cli/slack"
	"log"
)

func newGetConversationHistoryCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "get-conversation-history",
		Short: "Get the conversation history in a channel",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			channelID := viper.GetString("slack.channel-id")
			api := slack.NewApi(viper.GetString("slack.token"))
			ch, err := api.GetConversationHistory(slack.Channel{ID: channelID})
			if err != nil {
				return err
			}
			_ = json.PrintJSON(ch)
			return nil
		},
	}

	command.Flags().StringP("channel-id", "c", "", "Channel ID")

	if err := viper.BindPFlag("slack.channel-id", command.Flag("channel-id")); err != nil {
		log.Fatal(err)
	}
	return command
}
