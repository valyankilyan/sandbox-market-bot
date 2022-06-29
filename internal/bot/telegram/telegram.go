package telegram

import (
	"fmt"
	"log"
)

type TBot struct {
	token string
}

var apiAddr string = "https://api.telegram.org/bot%s/%s"

func New(token string) *TBot {
	if token == "" {
		log.Fatal("No telegram token was given.")
	}

	return &TBot{
		token: token,
	}
}

func (b *TBot) requestURL(command string) (string, error) {
	if command == "" {
		return "", fmt.Errorf("no command was given in requestURL")
	}
	return fmt.Sprintf(apiAddr, b.token, command), nil
}
