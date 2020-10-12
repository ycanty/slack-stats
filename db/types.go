package db

// Channel is a slack channel
type Channel struct {
	ID   string `gorm:"primaryKey"`
	Name string
}

// User is a slack user
type User struct {
	ID   string `gorm:"primaryKey"`
	Name string
}

// Reaction is a reaction icon (thumbs up, etc)
type Reaction struct {
	Name string `gorm:"primaryKey"`
}

// Message is a slack message
type Message struct {
	Timestamp string `gorm:"primaryKey"`

	Text       string
	IsStarred  bool
	ReplyCount int

	Channel   Channel `gorm:"foreignKey:ChannelID"`
	ChannelID string

	User   User `gorm:"foreignKey:UserID"`
	UserID string

	Reactions []MessageReaction `gorm:"foreignKey:MessageID;foreignKey:ReactionID"`
}

// MessageReaction defines user reactions to a given message
type MessageReaction struct {
	Message   Message `gorm:"foreignKey:MessageID"`
	MessageID string  `gorm:"primaryKey"`

	Reaction   Reaction `gorm:"foreignKey:ReactionID"`
	ReactionID string   `gorm:"primaryKey"`

	Count int
	Users []User `gorm:"many2many:users_reactions"`
}
