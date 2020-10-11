package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ycanty/go-cli/console"
	"github.com/ycanty/go-cli/slack"
)

func newImportCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "import",
		Short: "Import data into the statistics database",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			fileHandle, err := GetFileFlag(cmd, "file")

			if err != nil {
				return err
			}

			api := slack.NewApi(viper.GetString("token"))

			messages, err := api.ReadConversationHistory(fileHandle)

			if err != nil {
				return err
			}

			if err := console.PrintJSON(messages); err != nil {
				return err
			}

			// TODO Store to sqlite DB (https://gorm.io/docs/)

			return nil
		},
	}

	command.Flags().StringP("file", "f", "", "File name or - for stdin")
	if err := command.MarkFlagRequired("file"); err != nil {
		panic(err) // err here means programming error on name param of MarkFlagRequired
	}

	return command
}
