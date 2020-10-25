package db

import (
	"github.com/spf13/cobra"
)

func NewDBCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "db",
		Short: "Interact with the statistics database",
		Long:  ``,
	}

	command.AddCommand(newImportCommand())

	return command
}
