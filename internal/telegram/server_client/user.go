package server_client

import (
	"log"

	srv "github.com/valyankilyan/sandbox-market-bot/pkg/server_api"
)

func (s *Server) CreateUser(usr User) {
	nwusr := srv.User{
		Name:       usr.FirstName,
		TgUserName: usr.Username,
		TgId:       usr.ID,
	}

	s.client.CreateUser(s.ctx, &nwusr)
}

func (s *Server) getUserByTg(user User) *srv.User {
	usr, err := s.client.ReadUser(s.ctx, &srv.TgId{TgId: user.ID})
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

func (s *Server) UpdateTinkoffToken(user User, token string) error {
	usr := s.getUserByTg(user)
	usr.TinkoffToken = token
	_, err := s.client.UpdateUser(s.ctx, usr)
	log.Println("couldn't update User tinkoff token:", err)
	return err
}

func (s *Server) ReadUserToken(user User) (token string, err error) {
	usr, err := s.client.ReadUser(s.ctx, &srv.TgId{TgId: user.ID})
	if err != nil {
		log.Println("user don't have tinkoff token:", err)
		return "", err
	}

	return usr.TinkoffToken, nil
}
