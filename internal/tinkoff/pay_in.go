package tinkoff

import (
	"context"
	"fmt"
	"os"
	"time"

	"gitlab.ozon.dev/valyankilyan/homework-2-market-bot/config"
	"gitlab.ozon.dev/valyankilyan/homework-2-market-bot/internal/tinkoff/investapi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

func (t *Tinkoff) PayIn(u int64, n int32) (units int64, nano int32, err error) {
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
		// os.Exit(1)
		return 0, 0, err
	}
	defer conn.Close()

	client := investapi.NewSandboxServiceClient(conn)
	resp, err := client.GetSandboxAccounts(ctx, &investapi.GetAccountsRequest{})
	if err != nil {
		fmt.Println(err.Error())
		return 0, 0, err
	}
	fmt.Println(resp)

	spireq := investapi.SandboxPayInRequest{
		AccountId: resp.Accounts[0].Id,
		Amount: &investapi.MoneyValue{
			Units: u,
			Nano:  n,
		},
	}

	spiresp, err := client.SandboxPayIn(ctx, &spireq)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(3)
	}

	fmt.Printf("spiresp: %v\n", spiresp)

	return spiresp.Balance.Units, spireq.Amount.Nano, err
}
