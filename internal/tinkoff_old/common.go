package tinkoff_old

import (
	"context"
	"fmt"
	"time"

	"github.com/valyankilyan/sandbox-market-bot/config"
	"github.com/valyankilyan/sandbox-market-bot/internal/tinkoff/investapi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

func (t *Tinkoff) getContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	ctx = metadata.NewOutgoingContext(
		ctx,
		metadata.New(map[string]string{"Authorization": fmt.Sprintf("Bearer %v", t.Token)}),
	)

	return ctx, cancel
}

func (t *Tinkoff) getClientConn() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(
		config.Tinkoff.Endpoint,
		grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")),
	)
	return conn, err
}

func (t *Tinkoff) sandboxAccountId(ctx context.Context, client investapi.SandboxServiceClient) (string, error) {
	resp, err := client.GetSandboxAccounts(ctx, &investapi.GetAccountsRequest{})
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	// fmt.Println(resp)

	return resp.Accounts[0].Id, err
}
