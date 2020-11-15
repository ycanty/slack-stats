package report

type Count struct {
	Value    float64
	ValuePct float64
}

func newCount(value float64, total float64) Count {
	return Count{
		Value:    value,
		ValuePct: value / total,
	}
}

type AverageMessagesPerDay struct {
	AllDays   float64
	Monday    float64
	Tuesday   float64
	Wednesday float64
	Thursday  float64
	Friday    float64
	Saturday  float64
	Sunday    float64
}

type MessagePerUser struct {
	UserName string
	Messages Count
}

type MessageThreadCount struct {
	MessageText string
	ThreadCount Count
}

type Overview struct {
	TotalMessages         int64
	TotalUsers            int64
	TotalDays             int64
	AverageMessagesPerDay AverageMessagesPerDay
	MessagesPerUser       []MessagePerUser
	MessagesThreadCount   []MessageThreadCount
}

func (a *Api) Overview() (*Overview, error) {
	totalMessages, err := a.db.GetMessageCount()
	if err != nil {
		return nil, err
	}
	totalUsers, err := a.db.GetUserCount()
	if err != nil {
		return nil, err
	}

	totalDays, err := a.totalDays()
	if err != nil {
		return nil, err
	}

	return &Overview{
		TotalMessages: totalMessages,
		TotalUsers:    totalUsers,
		TotalDays:     totalDays,
		AverageMessagesPerDay: AverageMessagesPerDay{
			AllDays:   float64(totalMessages) / float64(totalDays),
			Monday:    0,
			Tuesday:   0,
			Wednesday: 0,
			Thursday:  0,
			Friday:    0,
			Saturday:  0,
			Sunday:    0,
		},
		MessagesPerUser:     nil,
		MessagesThreadCount: nil,
	}, nil
}

func (a *Api) totalDays() (int64, error) {
	firstMessage, err := a.db.GetFirstMessage()
	if err != nil {
		return 0, err
	}

	lastMessage, err := a.db.GetLastMessage()
	if err != nil {
		return 0, err
	}

	firstDate, err := firstMessage.Time()
	if err != nil {
		return 0, err
	}
	lastDate, err := lastMessage.Time()
	if err != nil {
		return 0, err
	}

	totalDays := int64(lastDate.Sub(firstDate).Hours() / 24)

	return totalDays, nil
}
