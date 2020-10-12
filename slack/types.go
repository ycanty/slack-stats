package slack

type Channel struct {
	Name string
	ID   string
}

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

type ConversationHistory struct {
	Channel  Channel   `json:"channel"`
	Messages []Message `json:"messages"`
}
