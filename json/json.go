package json

import (
	"fmt"
	colorjson "github.com/nwidger/jsoncolor"
	"io"
	"k8s.io/client-go/util/jsonpath"
)

var jsonPath string

// SetJSONPath sets the json path expression PrintJSON will use to filter out its output.
func SetJSONPath(path string) {
	jsonPath = path
}

func PrintJSON(writer io.Writer, obj interface{}) error {
	if jsonPath == "" {
		jsonBytes, err := colorjson.MarshalIndent(obj, "", "   ")
		if err != nil {
			return err
		}

		_, err = fmt.Fprintln(writer, string(jsonBytes))
		return err
	}

	parser := jsonpath.New("filter")
	if err := parser.Parse(jsonPath); err != nil {
		return err
	}

	return parser.Execute(writer, obj)
}
