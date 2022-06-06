package telegram

import (
	"gitlab.ozon.dev/valyankilyan/homework-2-market-bot/pkg/api"
)

func (b *Bot) createUser(usr User) {
	nwusr := api.User{
		Name:       usr.FirstName,
		TgUserName: usr.Username,
		TgId:       usr.ID,
	}

	b.client.CreateUser(b.ctx, &nwusr)
}
