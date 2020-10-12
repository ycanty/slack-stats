package slack

import (
	"fmt"
	"github.com/slack-go/slack"
	"strings"
	"time"
)

type Api struct {
	client *slack.Client
}

func NewApi(token string) *Api {
	return &Api{
		client: slack.New(token),
	}
}

func (a *Api) GetChannels(name string) ([]Channel, error) {
	channels, err := a.client.GetChannels(true)
	if err != nil {
		return nil, err
	}

	responses := make([]Channel, 0)

	for _, channel := range channels {
		if strings.Contains(channel.Name, name) {
			responses = append(responses, Channel{
				Name: channel.Name,
				ID:   channel.ID,
			})
		}
	}

	return responses, nil
}

func (a *Api) GetConversationHistory(channel Channel) (*ConversationHistory, error) {
	messages := make([]Message, 0)
	cursor := ""
	// TODO Get the start date from a param
	thetime := time.Now().AddDate(0, 0, -7).Unix() // since 7 days ago
	oldest := fmt.Sprintf("%d", thetime)
	//fmt.Fprintf(os.Stdout, "Epoch: %s\n", time.Unix(thetime, 0))
	for {
		history, err := a.client.GetConversationHistory(&slack.GetConversationHistoryParameters{
			ChannelID: channel.ID,
			Cursor:    cursor,
			Inclusive: false,
			Latest:    "",
			Limit:     100,
			Oldest:    oldest,
		})
		if err != nil {
			return nil, err
		}

		for _, msg := range history.Messages {
			reactions := make([]ItemReaction, 0, len(msg.Reactions))

			for _, reaction := range msg.Reactions {
				reactions = append(reactions, ItemReaction{
					Name:  reaction.Name,
					Count: reaction.Count,
					Users: reaction.Users,
				})
			}
			messages = append(messages, Message{
				User:       msg.User, // TODO query slack for real user name
				Text:       msg.Text,
				Timestamp:  msg.Timestamp, // TODO Convert to human-readable time
				IsStarred:  msg.IsStarred,
				ReplyCount: msg.ReplyCount,
				Reactions:  reactions,
			})
		}
		if len(history.ResponseMetaData.NextCursor) == 0 {
			break
		}
		cursor = history.ResponseMetaData.NextCursor
	}

	return &ConversationHistory{
		Channel:  channel,
		Messages: messages,
	}, nil
}
