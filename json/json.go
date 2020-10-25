package json

import (
	"fmt"
	colorjson "github.com/nwidger/jsoncolor"
)

func PrintJSON(obj interface{}) error {
	jsonBytes, err := colorjson.MarshalIndent(obj, "", "   ")

	if err != nil {
		return err
	}

	fmt.Println(string(jsonBytes))

	return nil
}
