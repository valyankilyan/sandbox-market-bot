package config

type TgToken string

type User struct {
	Name         string
	TgUserName   string
	TgId         int64
	TinkoffToken string
}

type Config struct {
	TgToken TgToken
	Users   map[int64]User
}
