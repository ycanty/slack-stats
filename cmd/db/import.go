package db

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ycanty/slack-stats/db"
	"github.com/ycanty/slack-stats/file"
	"github.com/ycanty/slack-stats/json"
	"github.com/ycanty/slack-stats/slack"
	"log"
)

const (
	configDBSqliteFilename = "db.sqlite.filename"
)

func newImportCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "import",
		Short: "Import data into the statistics database",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			fileHandle, err := file.OpenFile(cmd.Flag("from").Value.String())

			if err != nil {
				return err
			}

			dbFilename := viper.GetString(configDBSqliteFilename)

			ch, err := slack.NewConversationHistoryFromJSON(fileHandle)

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

	command.Flags().StringP("from", "f", "", "Input File name or - for stdin")
	if err := command.MarkFlagRequired("from"); err != nil {
		log.Fatal(err) // err here means programming error on name param of MarkFlagRequired
	}

	command.Flags().StringP("into", "i", "", "Database file name")
	if err := viper.BindPFlag(configDBSqliteFilename, command.Flags().Lookup("into")); err != nil {
		log.Fatal(err)
	}

	return command
}
