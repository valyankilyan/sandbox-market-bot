package telegram

import (
	"fmt"
	"log"
)

type Bot interface {
}

type TBot struct {
	token  string
	server Server
}

var apiAddr string = "https://api.telegram.org/bot%s/%s"

func New(token string, server Server) *TBot {
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

type Server interface {
	CreateUser(User)
	UpdateTinkoffToken(user User, token string) error
}
