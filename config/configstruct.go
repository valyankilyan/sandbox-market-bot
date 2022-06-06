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
}

type Tinkoff struct {
	Address string
}

type Config struct {
	Telegram Telegram
	Database Database
	Rpc      Rpc
	Tinkoff  Tinkoff
}
