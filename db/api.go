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

func (a *Api) SaveUser(user *User) error {
	tx := a.db.Save(user)
	return tx.Error
}

func (a *Api) SaveChannel(channel *Channel) error {
	tx := a.db.Save(channel)
	return tx.Error
}

func (a *Api) GetFirstMessage() (*Message, error) {
	msg := &Message{}
	if result := a.db.First(msg); result.Error != nil {
		return nil, result.Error
	}
	return msg, nil
}

func (a *Api) GetLastMessage() (*Message, error) {
	msg := &Message{}
	if result := a.db.Last(msg); result.Error != nil {
		return nil, result.Error
	}
	return msg, nil
}

func (a *Api) GetMessageCount() (int64, error) {
	return a.getCount(&Message{})
}

func (a *Api) GetMessageCountOnDay(day int) (int64, error) {
	// TODO
	var count int64
	if result := a.db.Model(&Message{}).Count(&count); result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}

func (a *Api) GetUserCount() (int64, error) {
	return a.getCount(&User{})
}

func (a *Api) getCount(object interface{}) (int64, error) {
	var count int64
	if result := a.db.Model(object).Count(&count); result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}

func (a *Api) GetUsersWithMissingNames() ([]User, error) {
	var users []User
	if result := a.db.Where("name = ''").Find(&users); result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (a *Api) GetChannelsWithMissingNames() ([]Channel, error) {
	var channels []Channel
	if result := a.db.Where("name = ''").Find(&channels); result.Error != nil {
		return nil, result.Error
	}
	return channels, nil
}

func convertMessage(channel slack.Channel, message slack.Message) Message {
	msg := Message{
		Channel:    convertChannel(channel),
		User:       convertUsers([]string{message.User})[0],
		Text:       message.Text,
		Timestamp:  message.Timestamp,
		IsStarred:  message.IsStarred,
		ReplyCount: message.ReplyCount,
		Permalink:  message.Permalink,
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
