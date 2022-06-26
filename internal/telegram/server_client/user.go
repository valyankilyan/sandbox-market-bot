package server_client

import (
	"log"

	"github.com/valyankilyan/sandbox-market-bot/internal/telegram"
	srv "github.com/valyankilyan/sandbox-market-bot/pkg/server_api"
)

func (s *Server) CreateUser(usr telegram.User) {
	nwusr := srv.User{
		Name:       usr.FirstName,
		TgUserName: usr.Username,
		TgId:       usr.TgId,
	}

	s.client.CreateUser(s.ctx, &nwusr)
}

func (s *Server) UpdateTinkoffToken(user telegram.User, token string) error {
	usr := s.getUserByTg(user)
	usr.TinkoffToken = token
	_, err := s.client.UpdateUser(s.ctx, usr)
	log.Println("couldn't update User tinkoff token:", err)
	return err
}

func (s *Server) getUserByTg(user telegram.User) *srv.User {
	usr, err := s.client.ReadUser(s.ctx, &srv.TgId{TgId: user.TgId})
	if err != nil {
		log.Println("error in updateTinkofToken:", err)
		return &srv.User{}
	}
	if usr == nil {
		s.CreateUser(user)
		// silly but i will change it
		return s.getUserByTg(user)
	}

	return usr
}

func (s *Server) readUserToken(user telegram.User) (token string, err error) {
	usr, err := s.client.ReadUser(s.ctx, &srv.TgId{TgId: user.TgId})
	if err != nil {
		log.Println("user don't have tinkoff token:", err)
		return "", err
	}

	return usr.TinkoffToken, nil
}
