package argparse

import (
	"github.com/spf13/cobra"
	"github.com/ycanty/slack-stats/file"
	"io"
)

// GetFileFlag reads the given flag as a filename and returns an io.Reader opened on that file.
// It returns os.Stdin if the filename is "-".
func GetFileFlag(cmd *cobra.Command, flagName string) (io.Reader, error) {
	filename, err := cmd.Flags().GetString(flagName)

	if err != nil {
		return nil, err
	}
	return file.OpenFile(filename)
}
