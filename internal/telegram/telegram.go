package telegram

import (
	"context"
	"fmt"
	"log"

	srv "github.com/valyankilyan/sandbox-market-bot/pkg/server_api"
)

type Bot interface {
	New(token string, client srv.UserServiceClient, ctx context.Context)
}

type TBot struct {
	token  string
	client srv.UserServiceClient
	ctx    context.Context
}

var apiAddr string = "https://api.telegram.org/bot%s/%s"

func New(token string, client srv.UserServiceClient, ctx context.Context) *TBot {
	if token == "" {
		log.Fatal("No telegram token was given.")
	}

	return &TBot{
		token:  token,
		client: client,
		ctx:    ctx,
	}
}

func (b *TBot) requestURL(command string) (string, error) {
	if command == "" {
		return "", fmt.Errorf("no command was given in requestURL")
	}
	return fmt.Sprintf(apiAddr, b.token, command), nil
}
