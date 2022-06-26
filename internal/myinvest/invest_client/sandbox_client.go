package invest_client

import (
	"fmt"
	"log"

	"github.com/valyankilyan/sandbox-market-bot/internal/myinvest"
	"github.com/valyankilyan/sandbox-market-bot/pkg/investapi"
)

// "github.com/ofen/tinkoff-invest-example/tinkoff/investapi"

func (ic *InvestClient) PayIn(token string, quantity myinvest.Quotation) (myinvest.Quotation, error) {
	ctx, cancel := getContext(token)
	defer cancel()

	conn, err := getClientConn()
	if err != nil {
		log.Println("payin clientconn", err)
		return myinvest.Quotation{}, err
	}

	client := investapi.NewSandboxServiceClient(conn)

	account_id, err := sandboxAccountId(ctx, client)
	if err != nil {
		log.Println("acc id", err)
		return myinvest.Quotation{}, err
	}

	spireq := investapi.SandboxPayInRequest{
		AccountId: account_id,
		Amount: &investapi.MoneyValue{
			Units: quantity.Units,
			Nano:  quantity.Nano,
		},
	}

	spiresp, err := client.SandboxPayIn(ctx, &spireq)
	if err != nil {
		fmt.Printf("sandbox payin %v\n", err)
		return myinvest.Quotation{}, err
	}

	log.Println("Payin from invest_client")
	return myinvest.Quotation{
		Units: spiresp.Balance.Units,
		Nano:  spiresp.Balance.Nano,
	}, nil
}
