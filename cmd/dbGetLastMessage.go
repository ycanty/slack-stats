package cmd

import (
	"github.com/spf13/cobra"
	"github.com/ycanty/slack-stats/json"
)

func newGetLastMessageCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "get-last-message",
		Short: "Returns the last conversation stored in the database",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			dbClient, err := dbApi()

			if err != nil {
				return err
			}

			msg, err := dbClient.GetLastMessage()

			if err != nil {
				return err
			}

			if err := json.PrintJSON(cmd.OutOrStdout(), msg); err != nil {
				return err
			}
			return nil
		},
	}

	return command
}
