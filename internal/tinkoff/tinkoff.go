package tinkoff

import (
	"context"
	"fmt"
	"os"
	"time"

	"gitlab.ozon.dev/valyankilyan/homework-2-market-bot/internal/app/tinkoff/investapi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

type Tinkoff struct {
	Token string
	// Conn  *grpc.ClientConn
}

func New(token string) *Tinkoff {
	return &Tinkoff{
		Token: token,
	}
}

const endpoint = "invest-public-api.tinkoff.ru:443"

func (t *Tinkoff) GetAccounts() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	ctx = metadata.NewOutgoingContext(
		ctx,
		metadata.New(map[string]string{"Authorization": fmt.Sprintf("Bearer %v", t.Token)}),
	)

	conn, err := grpc.Dial(
		endpoint,
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

	spireq := investapi.SandboxPayInRequest{
		AccountId: resp.Accounts[0].Id,
		Amount: &investapi.MoneyValue{
			Units: 100,
			Nano:  123456,
		},
	}

	spiresp, err := client.SandboxPayIn(ctx, &spireq)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(3)
	}

	fmt.Printf("spiresp: %v\n", spiresp)

	return nil
}
