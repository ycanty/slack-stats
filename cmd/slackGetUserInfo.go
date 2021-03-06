package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/ycanty/slack-stats/json"
	"github.com/ycanty/slack-stats/slack"
)

func newGetUserInfoCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "get-user-info",
		Short: "Get information about a user",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			id := cmd.Flag("id").Value.String()
			email := cmd.Flag("email").Value.String()

			if id == "" && email == "" {
				return fmt.Errorf("please specify --id or --email option")
			}
			if id != "" && email != "" {
				return fmt.Errorf("id and email options are exclusive of each other")
			}

			var user *slack.User
			var err error

			if id != "" {
				user, err = slackApi().GetUserInfo(id)
			} else if email != "" {
				user, err = slackApi().GetUserInfo(email)
			}

			if err != nil {
				return err
			}
			return json.PrintJSON(cmd.OutOrStdout(), user)
		},
	}

	command.Flags().StringP("id", "i", "", "User ID")
	command.Flags().StringP("email", "e", "", "User E-Mail")
	return command
}
