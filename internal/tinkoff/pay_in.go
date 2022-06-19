package tinkoff

import (
	"fmt"
	"os"

	"gitlab.ozon.dev/valyankilyan/homework-2-market-bot/internal/tinkoff/investapi"
)

func (t *Tinkoff) PayIn(u int64, n int32) (units int64, nano int32, err error) {
	ctx, cancel := t.getContext()
	defer cancel()

	conn, err := t.getClientConn()
	if err != nil {
		fmt.Println(err.Error())
		return 0, 0, err
	}
	defer conn.Close()

	client := investapi.NewSandboxServiceClient(conn)

	account_id, err := t.sandboxAccountId(ctx, client)

	spireq := investapi.SandboxPayInRequest{
		AccountId: account_id,
		Amount: &investapi.MoneyValue{
			Units: u,
			Nano:  n,
		},
	}

	spiresp, err := client.SandboxPayIn(ctx, &spireq)
	if err != nil {
		fmt.Printf("sandbox payin %v\n", err.Error())
		os.Exit(3)
	}

	return spiresp.Balance.Units, spiresp.Balance.Nano, err
}
