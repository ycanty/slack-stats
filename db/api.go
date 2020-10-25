package db

import (
	"github.com/ycanty/slack-stats/slack"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Api struct {
	db *gorm.DB
}

func Open(file string) (*Api, error) {
	db, err := gorm.Open(sqlite.Open(file), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	api := &Api{
		db: db,
	}

	if err = db.AutoMigrate(&Message{}, &Channel{}, &User{}, &Reaction{}, &MessageReaction{}); err != nil {
		return nil, err
	}

	return api, nil
}

func (a *Api) Save(ch *slack.ConversationHistory) error {
	dbMessages := make([]Message, 0)
	for _, msg := range ch.Messages {
		dbMessages = append(dbMessages, convertMessage(ch.Channel, msg))
	}

	tx := a.db.Save(dbMessages)
	return tx.Error
}

func convertMessage(channel slack.Channel, message slack.Message) Message {
	msg := Message{
		Channel:    convertChannel(channel),
		User:       convertUsers([]string{message.User})[0],
		Text:       message.Text,
		Timestamp:  message.Timestamp,
		IsStarred:  message.IsStarred,
		ReplyCount: message.ReplyCount,
	}
	msg.Reactions = convertReactions(msg, message.Reactions)
	return msg
}

func convertChannel(channel slack.Channel) Channel {
	return Channel{
		ID:   channel.ID,
		Name: channel.Name,
	}
}

func convertUsers(userIDs []string) []User {
	dbUsers := make([]User, 0, len(userIDs))

	for _, id := range userIDs {
		dbUsers = append(dbUsers, User{
			ID: id,
		})
	}

	return dbUsers
}

func convertReactions(message Message, reactions []slack.ItemReaction) []MessageReaction {
	msgReactions := make([]MessageReaction, 0, len(reactions))

	for _, reaction := range reactions {
		msgReaction := MessageReaction{
			Reaction: Reaction{
				Name: reaction.Name,
			},
			Message: message,
			Count:   reaction.Count,
			Users:   convertUsers(reaction.Users),
		}
		msgReactions = append(msgReactions, msgReaction)
	}

	return msgReactions
}
