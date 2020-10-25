package argparse

import (
	"github.com/spf13/cobra"
	"io"
	"os"
)

// GetFileFlag reads the given flag as a filename and returns an io.Reader opened on that file.
// It returns os.Stdin if the filename is "-".
func GetFileFlag(cmd *cobra.Command, flagName string) (io.Reader, error) {
	file, err := cmd.Flags().GetString(flagName)

	if err != nil {
		return nil, err
	}

	fileHandle := os.Stdin
	if file != "-" {
		if fileHandle, err = os.Open(file); err != nil {
			return nil, err
		}
	}

	return fileHandle, nil
}
