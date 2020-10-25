package db

import (
	"github.com/spf13/cobra"
	"github.com/ycanty/slack-stats/file"
	"github.com/ycanty/slack-stats/json"
	"github.com/ycanty/slack-stats/slack"
	"log"
)

func newImportCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "import",
		Short: "Import data into the statistics database",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			fileHandle, err := file.OpenFile(cmd.Flag("file").Value.String())

			if err != nil {
				return err
			}

			ch, err := slack.NewConversationHistoryFromJSON(fileHandle)

			if err != nil {
				return err
			}

			if err := json.PrintJSON(ch); err != nil {
				return err
			}

			dbClient, err := dbApi()

			if err != nil {
				return err
			}

			err = dbClient.Save(ch)

			return err
		},
	}

	command.Flags().StringP("file", "f", "", "Input File name or - for stdin")
	if err := command.MarkFlagRequired("file"); err != nil {
		log.Fatal(err) // err here means programming error on name param of MarkFlagRequired
	}

	return command
}
