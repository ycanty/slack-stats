package cmd

import (
	"github.com/spf13/cobra"
	"github.com/ycanty/slack-stats/json"
	"log"
)

func newFindChannelIDCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "find-channel-id",
		Short: "Find the channel ID from a channel name, or part of it",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			channels, err := slackApi().GetChannels(cmd.Flag("name").Value.String())
			if err != nil {
				return err
			}
			return json.PrintJSON(cmd.OutOrStdout(), channels)
		},
	}

	command.Flags().StringP("name", "n", "", "Channel name")
	if err := command.MarkFlagRequired("name"); err != nil {
		log.Fatal(err)
	}

	return command
}
