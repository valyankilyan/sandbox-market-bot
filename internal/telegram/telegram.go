package telegram

import (
	"context"
	"fmt"

	"gitlab.ozon.dev/valyankilyan/homework-2-market-bot/pkg/api"
)

type Bot struct {
	token  string
	client api.MarketBotClient
	ctx    context.Context
}

var apiAddr string = "https://api.telegram.org/bot%s/%s"

func (b *Bot) requestURL(command string) (string, error) {
	if command == "" || b.token == "" {
		return "", fmt.Errorf("something went wrong")
	}
	return fmt.Sprintf(apiAddr, b.token, command), nil
}

func New(token string, client api.MarketBotClient, ctx context.Context) *Bot {
	return &Bot{
		token:  token,
		client: client,
		ctx:    ctx,
	}
}
