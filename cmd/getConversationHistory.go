package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/slack-go/slack"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"time"
)

type Message struct {
	// Basic Message
	User       string `json:"user,omitempty"`
	Text       string `json:"text,omitempty"`
	Timestamp  string `json:"ts,omitempty"`
	IsStarred  bool   `json:"is_starred,omitempty"`
	ReplyCount int    `json:"reply_count,omitempty"`

	// reactions
	Reactions []ItemReaction `json:"reactions,omitempty"`
}

type ItemReaction struct {
	Name  string   `json:"name"`
	Count int      `json:"count"`
	Users []string `json:"users"`
}

func newGetConversationHistoryCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "get-conversation-history",
		Short: "Get the conversation history in a channel",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			channelID := viper.GetString("channel-id")

			messages := make([]Message, 0)
			cursor := ""
			// TODO Get the start date from a CLI param
			thetime := time.Now().AddDate(0, 0, -3).Unix() // since 3 days ago
			oldest := fmt.Sprintf("%d", thetime)
			//fmt.Fprintf(os.Stdout, "Epoch: %s\n", time.Unix(thetime, 0))
			for {
				history, err := Api().GetConversationHistory(&slack.GetConversationHistoryParameters{
					ChannelID: channelID,
					Cursor:    cursor,
					Inclusive: false,
					Latest:    "",
					Limit:     100,
					Oldest:    oldest,
				})
				if err != nil {
					return err
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

			jsonBytes, err := json.Marshal(messages)

			if err != nil {
				return err
			}

			fmt.Println(string(jsonBytes))

			// TODO Store to sqlite DB (https://gorm.io/docs/)
			return nil
		},
	}

	command.Flags().StringP("channel-id", "c", "", "Channel ID")
	//err := command.MarkFlagRequired("channel-id")

	//if err != nil {
	//	// Ok to panic: MarkFlagRequired() returns an error only if we use an invalid flag name
	//	panic(err)
	//}

	if err := viper.BindPFlag("channel-id", command.Flags().Lookup("channel-id")); err != nil {
		log.Fatal(err)
	}
	return command
}
