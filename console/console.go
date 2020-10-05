package console

import (
	"fmt"
	json "github.com/nwidger/jsoncolor"
)

func PrintJSON(obj interface{}) error {
	jsonBytes, err := json.MarshalIndent(obj, "", "   ")

	if err != nil {
		return err
	}

	fmt.Println(string(jsonBytes))

	return nil
}
