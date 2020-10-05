package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ycanty/go-cli/console"
	"github.com/ycanty/go-cli/slack"
)

func newFindChannelIDCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "find-channel-id",
		Short: "Find the channel ID from a channel name",
		Long:  ``,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			api := slack.NewApi(viper.GetString("token"))
			channels, err := api.GetChannels(args[0])
			if err != nil {
				return err
			}
			_ = console.PrintJSON(channels)
			return nil
		},
	}

	return command
}
