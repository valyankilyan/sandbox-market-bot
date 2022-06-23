package telegram_old

import (
	"context"
	"fmt"

	srv "github.com/valyankilyan/sandbox-market-bot/pkg/server_api"
)

type Bot struct {
	token  string
	client srv.MarketBotClient
	ctx    context.Context
}

var apiAddr string = "https://srv.telegram.org/bot%s/%s"

func (b *Bot) requestURL(command string) (string, error) {
	if command == "" || b.token == "" {
		return "", fmt.Errorf("something went wrong")
	}
	return fmt.Sprintf(apiAddr, b.token, command), nil
}

func New(token string, client srv.MarketBotClient, ctx context.Context) *Bot {
	return &Bot{
		token:  token,
		client: client,
		ctx:    ctx,
	}
}
