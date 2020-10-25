package slack

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

func NewConversationHistoryFromJSON(reader io.Reader) (*ConversationHistory, error) {
	var messages *ConversationHistory
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &messages); err != nil {
		return nil, err
	}

	return messages, nil
}
