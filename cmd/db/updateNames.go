package db

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	cmdslack "github.com/ycanty/slack-stats/cmd/slack"
	"github.com/ycanty/slack-stats/slack"
)

func newUpdateNamesCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "update-names",
		Short: "Query slack to fill user and channel names missing from the database",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			dbClient, err := dbApi()

			if err != nil {
				return err
			}

			slackApi := slack.NewApi(viper.GetString(cmdslack.ConfigSlackToken))

			users, err := dbClient.GetUsersWithMissingNames()

			if err != nil {
				return err
			}

			for _, user := range users {
				userInfo, err := slackApi.GetUserInfo(user.ID)
				if err != nil {
					return nil
				}
				user.Name = userInfo.Name
				user.RealName = userInfo.RealName
				if err := dbClient.SaveUser(&user); err != nil {
					return err
				}
			}

			channels, err := dbClient.GetChannelsWithMissingNames()

			if err != nil {
				return err
			}

			for _, channel := range channels {
				channelInfo, err := slackApi.GetChannelInfo(channel.ID)
				if err != nil {
					return nil
				}
				channel.Name = channelInfo.Name
				if err := dbClient.SaveChannel(&channel); err != nil {
					return err
				}
			}

			return nil
		},
	}

	return command
}
