package file

import (
	"io"
	"os"
)

// OpenFile reads the filename and returns an io.Reader opened on that file.
// It returns os.Stdin if the filename is "-".
func OpenFile(filename string) (io.Reader, error) {
	fileHandle := os.Stdin
	var err error
	if filename != "-" {
		if fileHandle, err = os.Open(filename); err != nil {
			return nil, err
		}
	}

	return fileHandle, nil
}
