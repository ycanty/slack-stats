package cmd

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/ycanty/go-cli/console"
	"github.com/ycanty/go-cli/slack"
	"io/ioutil"
	"os"
)

func newImportCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "import",
		Short: "Import data into the statistics database",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			file, err := cmd.Flags().GetString("file")

			if err != nil {
				return err
			}

			fileHandle := os.Stdin
			if file != "-" {
				if fileHandle, err = os.Open(file); err != nil {
					return err
				}
			}
			var messages []slack.Message
			data, err := ioutil.ReadAll(fileHandle)
			if err != nil {
				return err
			}
			if err := json.Unmarshal(data, &messages); err != nil {
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
