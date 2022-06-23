package telegram_old

import (
	"context"
	"fmt"
	"log"

	srv "github.com/valyankilyan/sandbox-market-bot/pkg/server_api"
)

type Bot struct {
	token  string
	client srv.MarketBotClient
	ctx    context.Context
}

var apiAddr string = "https://api.telegram.org/bot%s/%s"

func (b *Bot) requestURL(command string) (string, error) {
	if command == "" {
		return "", fmt.Errorf("no command was given in requestURL")
	}
	return fmt.Sprintf(apiAddr, b.token, command), nil
}

func New(token string, client srv.MarketBotClient, ctx context.Context) *Bot {
	if token == "" {
		log.Fatal("No telegram token was given.")
	}

	return &Bot{
		token:  token,
		client: client,
		ctx:    ctx,
	}
}
