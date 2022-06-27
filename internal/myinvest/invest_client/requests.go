package invest_client

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/valyankilyan/sandbox-market-bot/config"
	"github.com/valyankilyan/sandbox-market-bot/pkg/investapi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

func getContext(token string) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	ctx = metadata.NewOutgoingContext(
		ctx,
		metadata.New(map[string]string{"Authorization": fmt.Sprintf("Bearer %v", token)}),
	)

	return ctx, cancel
}

func getClientConn() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(
		config.Tinkoff.Endpoint,
		grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")),
	)
	return conn, err
}

func sandboxAccountId(ctx context.Context, client investapi.SandboxServiceClient) (string, error) {
	resp, err := client.GetSandboxAccounts(ctx, &investapi.GetAccountsRequest{})
	if err != nil {
		log.Println("sandboxAccountId", err.Error())
		return "", err
	}
	// fmt.Println(resp)

	return resp.Accounts[0].Id, err
}
