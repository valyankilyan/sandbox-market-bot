package server_client

import (
	"context"

	srv "github.com/valyankilyan/sandbox-market-bot/pkg/server_api"
)

type Server struct {
	client srv.UserServiceClient
	ctx    context.Context
}

func New(client srv.UserServiceClient, ctx context.Context) *Server {
	return &Server{
		client: client,
		ctx:    ctx,
	}
}
