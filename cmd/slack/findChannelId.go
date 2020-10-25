package slack

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ycanty/go-cli/json"
	"github.com/ycanty/go-cli/slack"
)

func newFindChannelIDCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "find-channel-id",
		Short: "Find the channel ID from a channel name, or part of it",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			api := slack.NewApi(viper.GetString("slack.token"))
			channels, err := api.GetChannels(cmd.Flag("name").Value.String())
			if err != nil {
				return err
			}
			_ = json.PrintJSON(channels)
			return nil
		},
	}

	command.Flags().StringP("name", "n", "", "Channel name")
	if err := command.MarkFlagRequired("name"); err != nil {
		panic(err)
	}

	return command
}
