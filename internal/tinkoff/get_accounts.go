package tinkoff

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/valyankilyan/sandbox-market-bot/config"
	"github.com/valyankilyan/sandbox-market-bot/internal/tinkoff/investapi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

func (t *Tinkoff) GetAccounts() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	ctx = metadata.NewOutgoingContext(
		ctx,
		metadata.New(map[string]string{"Authorization": fmt.Sprintf("Bearer %v", t.Token)}),
	)

	conn, err := grpc.Dial(
		config.Conf.Tinkoff.Endpoint,
		grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")),
	)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	client := investapi.NewSandboxServiceClient(conn)
	resp, err := client.GetSandboxAccounts(ctx, &investapi.GetAccountsRequest{})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
	fmt.Println(resp)

	return nil
}
