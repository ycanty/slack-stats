package slack

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
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

	if err := viper.BindPFlag("slack.token", command.PersistentFlags().Lookup("token")); err != nil {
		log.Fatal(err)
	}

	return command
}
