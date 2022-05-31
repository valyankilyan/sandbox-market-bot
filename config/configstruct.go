package config

type Telegram struct {
	Token      string
	GetUpdates struct {
		Offset         int64
		Limit          int64
		Timeout        int
		AllowedUpdates []string
	}
}

type User struct {
	Name         string
	TgUserName   string
	TgId         int64
	TinkoffToken string
}

type Config struct {
	Telegram Telegram
	Users    map[int64]User
}
