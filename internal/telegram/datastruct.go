package telegram

import "github.com/valyankilyan/sandbox-market-bot/internal/telegram/server_client"

type message struct {
	MessageID int64
	Chat      chat
	From      server_client.User
	Date      int64
	Text      string
}

type chat struct {
	ID        int64
	FirstName string
	Username  string
	Type      string
}
