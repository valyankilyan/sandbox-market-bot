package config

type telegram struct {
	Token      string
	GetUpdates struct {
		Offset         int64
		Limit          int64
		Timeout        int
		AllowedUpdates []string
	}
}

type database struct {
	Host     string
	Port     string
	Username string
	Password string
}

type rpc struct {
	Host string
	Port string
}

type myinvest struct {
	Host string
	Port string
}

type tinkoff struct {
	Endpoint     string
	DefaultToken string
	UpdateTime   int32
}

// type Config struct {
// 	Telegram Telegram
// 	Database Database
// 	Rpc      Rpc
// 	Tinkoff  Tinkoff
// }
