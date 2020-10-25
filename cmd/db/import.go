package db

import (
	"github.com/spf13/cobra"
	"github.com/ycanty/go-cli/argparse"
	"github.com/ycanty/go-cli/db"
	"github.com/ycanty/go-cli/json"
)

func newImportCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "import",
		Short: "Import data into the statistics database",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			fileHandle, err := argparse.GetFileFlag(cmd, "file")

			if err != nil {
				return err
			}

			dbFilename, err := cmd.Flags().GetString("into")

			if err != nil {
				return err
			}

			ch, err := json.ReadConversationHistory(fileHandle)

			if err != nil {
				return err
			}

			if err := json.PrintJSON(ch); err != nil {
				return err
			}

			dbClient, err := db.Open(dbFilename)

			if err != nil {
				return err
			}

			err = dbClient.Save(ch)

			return err
		},
	}

	command.Flags().StringP("file", "f", "", "File name or - for stdin")
	if err := command.MarkFlagRequired("file"); err != nil {
		panic(err) // err here means programming error on name param of MarkFlagRequired
	}

	command.Flags().StringP("into", "i", "", "Database file name")
	if err := command.MarkFlagRequired("into"); err != nil {
		panic(err) // err here means programming error on name param of MarkFlagRequired
	}

	return command
}
