package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/slack-go/slack"
	"github.com/spf13/cobra"
)

func newGetConversationHistoryCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "get-conversation-history",
		Short: "Get the conversation history in a channel",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			channelID, err := cmd.Flags().GetString("channel-id")
			if err != nil {
				return err
			}
			history, err := Api().GetConversationHistory(&slack.GetConversationHistoryParameters{
				ChannelID: channelID,
				Cursor:    "",
				Inclusive: false,
				Latest:    "",
				Limit:     5,
				Oldest:    "",
			})

			if err != nil {
				return err
			}

			jsonBytes, err := json.Marshal(history)

			if err != nil {
				return err
			}
			fmt.Println(string(jsonBytes))

			return nil
		},
	}

	command.Flags().StringP("channel-id", "c", "", "Channel ID")
	err := command.MarkFlagRequired("channel-id")

	if err != nil {
		// Ok to panic: MarkFlagRequired() returns an error only if we use an invalid flag name
		panic(err)
	}

	return command
}
