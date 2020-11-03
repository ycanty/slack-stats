package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ycanty/slack-stats/db"
	"log"
)

const (
	configDBSqliteFilename = "db.sqlite.filename"
)

func NewDBCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "db",
		Short: "Interact with the statistics database",
		Long:  ``,
	}

	command.AddCommand(newImportCommand(), newGetLastMessageCommand(), newUpdateNamesCommand())

	command.PersistentFlags().StringP("dbfile", "d", "", "Database file name")
	if err := viper.BindPFlag(configDBSqliteFilename, command.PersistentFlags().Lookup("dbfile")); err != nil {
		log.Fatal(err)
	}

	return command
}

func dbApi() (*db.Api, error) {
	dbFilename := viper.GetString(configDBSqliteFilename)
	return db.Open(dbFilename)
}
