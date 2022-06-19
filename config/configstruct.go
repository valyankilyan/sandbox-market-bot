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

type Rpc struct {
	Host string
	Port string
}

type Tinkoff struct {
	Endpoint     string
	DefaultToken string
	UpdateTime   int32
}

type Config struct {
	Telegram Telegram
	Database Database
	Rpc      Rpc
	Tinkoff  Tinkoff
}
