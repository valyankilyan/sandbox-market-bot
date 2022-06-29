package telegram

import (
	"fmt"
	"log"

	"github.com/valyankilyan/sandbox-market-bot/internal/bot"
)

type TBot struct {
	token  string
	server bot.Server
	invest bot.Invest
}

var apiAddr string = "https://api.telegram.org/bot%s/%s"

func New(token string, server bot.Server, invest bot.Invest) *TBot {
	if token == "" {
		log.Fatal("No telegram token was given.")
	}

	return &TBot{
		token:  token,
		server: server,
		invest: invest,
	}
}

func (b *TBot) requestURL(command string) (string, error) {
	if command == "" {
		return "", fmt.Errorf("no command was given in requestURL")
	}
	return fmt.Sprintf(apiAddr, b.token, command), nil
}
