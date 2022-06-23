package telegram

import (
	"fmt"
	"log"

	sc "github.com/valyankilyan/sandbox-market-bot/internal/telegram/server_client"
)

type Bot interface {
	New(token string)
}

type TBot struct {
	token  string
	server *sc.Server
}

var apiAddr string = "https://api.telegram.org/bot%s/%s"

func New(token string, server *sc.Server) *TBot {
	if token == "" {
		log.Fatal("No telegram token was given.")
	}

	return &TBot{
		token:  token,
		server: server,
	}
}

func (b *TBot) requestURL(command string) (string, error) {
	if command == "" {
		return "", fmt.Errorf("no command was given in requestURL")
	}
	return fmt.Sprintf(apiAddr, b.token, command), nil
}
