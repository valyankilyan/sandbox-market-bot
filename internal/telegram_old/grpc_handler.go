package telegram_old

import (
	"log"

	srv "github.com/valyankilyan/sandbox-market-bot/pkg/server_api"
)

func (b *Bot) createUser(usr User) {
	nwusr := srv.User{
		Name:       usr.FirstName,
		TgUserName: usr.Username,
		TgId:       usr.ID,
	}

	b.client.CreateUser(b.ctx, &nwusr)
}

func (b *Bot) updateTinkoffToken(tgid int64, token string) error {
	usr, err := b.client.ReadUser(b.ctx, &srv.TgId{TgId: tgid})
	if err != nil {
		log.Printf("error in updateTinkofToken: %v", err)
		return err
	}

	usr.TinkoffToken = token
	_, err = b.client.UpdateUser(b.ctx, usr)
	return err
}

func (b *Bot) readUserToken(tgid int64) (token string, err error) {
	usr, err := b.client.ReadUser(b.ctx, &srv.TgId{TgId: tgid})
	if err != nil {
		log.Printf("error in updateTinkofToken: %v", err)
		return "", err
	}

	return usr.TinkoffToken, nil
}
