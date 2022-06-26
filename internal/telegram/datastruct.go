package telegram

type User struct {
	TgId         int64
	IsBot        bool
	FirstName    string
	Username     string
	LanguageCode string
}
type message struct {
	MessageID int64
	Chat      chat
	From      User
	Date      int64
	Text      string
}

type chat struct {
	ID        int64
	FirstName string
	Username  string
	Type      string
}
