package slack

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ycanty/go-cli/json"
	"github.com/ycanty/go-cli/slack"
)

func newGetUserInfoCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "get-user-info",
		Short: "Get information about a user",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			api := slack.NewApi(viper.GetString("slack.token"))
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
				user, err = api.GetUserInfo(id)
			} else if email != "" {
				user, err = api.GetUserInfo(email)
			}

			if err != nil {
				return err
			}
			_ = json.PrintJSON(user)
			return nil
		},
	}

	command.Flags().StringP("id", "i", "", "User ID")
	command.Flags().StringP("email", "e", "", "User E-Mail")
	return command
}
