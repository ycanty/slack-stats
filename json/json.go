package json

import (
	"encoding/json"
	"fmt"
	colorjson "github.com/nwidger/jsoncolor"
	"github.com/ycanty/slack-stats/slack"
	"io"
	"io/ioutil"
)

func PrintJSON(obj interface{}) error {
	jsonBytes, err := colorjson.MarshalIndent(obj, "", "   ")

	if err != nil {
		return err
	}

	fmt.Println(string(jsonBytes))

	return nil
}

func ReadConversationHistory(reader io.Reader) (*slack.ConversationHistory, error) {
	var messages *slack.ConversationHistory
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &messages); err != nil {
		return nil, err
	}

	return messages, nil
}
