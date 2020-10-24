package cmd

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
			channelID := viper.GetString("channel-id")
			api := slack.NewApi(viper.GetString("token"))
			ch, err := api.GetConversationHistory(slack.Channel{ID: channelID})
			if err != nil {
				return err
			}
			_ = json.PrintJSON(ch)
			return nil
		},
	}

	command.Flags().StringP("channel-id", "c", "", "Channel ID")
	//err := command.MarkFlagRequired("channel-id")

	//if err != nil {
	//	// Ok to panic: MarkFlagRequired() returns an error only if we use an invalid flag name
	//	panic(err)
	//}

	if err := viper.BindPFlag("channel-id", command.Flag("channel-id")); err != nil {
		log.Fatal(err)
	}
	return command
}
