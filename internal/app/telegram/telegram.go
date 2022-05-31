package telegram

import (
	"fmt"
)

type Bot struct {
	token string
}

var apiAddr string = "https://api.telegram.org/bot%s/%s"

func (b *Bot) requestURL(command string) (string, error) {
	if command == "" || b.token == "" {
		return "", fmt.Errorf("something went wrong")
	}
	return fmt.Sprintf(apiAddr, b.token, command), nil
}

func New(token string) *Bot {
	return &Bot{
		token: token,
	}
}
