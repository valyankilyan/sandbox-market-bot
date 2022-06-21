package tinkoff

import (

	// "github.com/ofen/tinkoff-invest-example/tinkoff/investapi"

	"github.com/valyankilyan/sandbox-market-bot/config"
	"github.com/valyankilyan/sandbox-market-bot/internal/tinkoff/investapi"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
)

func NewSandboxServiceClient(token string) investapi.SandboxServiceClient {
	conn, err := grpc.Dial(
		config.Conf.Tinkoff.Endpoint,
		grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")),
		grpc.WithPerRPCCredentials(oauth.NewOauthAccess(&oauth2.Token{AccessToken: token})),
	)
	if err != nil {
		panic(err)
	}

	return investapi.NewSandboxServiceClient(conn)
}
