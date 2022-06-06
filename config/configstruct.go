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

type Database struct {
	Host     string
	Port     string
	Username string
	Password string
}

type Config struct {
	Telegram Telegram
	Database Database
}
