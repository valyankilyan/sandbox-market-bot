package tinkoff_old

import (
	"fmt"

	"github.com/valyankilyan/sandbox-market-bot/pkg/investapi"
)

type Balance struct {
	TotalAmount Quotation
	Currencies  []MoneyValue
}

func (t *Tinkoff) GetSandboxPortfolio(blc *Balance) {
	ctx, cancel := t.getContext()
	defer cancel()

	conn, err := t.getClientConn()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()

	client := investapi.NewSandboxServiceClient(conn)
	account_id, err := t.sandboxAccountId(ctx, client)

	pfreq := investapi.PortfolioRequest{
		AccountId: account_id,
	}

	pfresp, err := client.GetSandboxPortfolio(ctx, &pfreq)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Portfolio response:")
	fmt.Printf("\tCurrencies_Amount %v\n", pfresp.TotalAmountCurrencies)
	fmt.Println("\tPositions:")
	for _, pos := range pfresp.Positions {
		fmt.Printf("\t\t %v\n", pos)
	}

	blc.TotalAmount.Units = pfresp.TotalAmountCurrencies.Units
	blc.TotalAmount.Nano = pfresp.TotalAmountCurrencies.Nano
}

func (t *Tinkoff) GetSandboxPosition(blc *Balance) {
	ctx, cancel := t.getContext()
	defer cancel()

	conn, err := t.getClientConn()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()

	client := investapi.NewSandboxServiceClient(conn)
	account_id, err := t.sandboxAccountId(ctx, client)

	posreq := investapi.PositionsRequest{
		AccountId: account_id,
	}

	posresp, err := client.GetSandboxPositions(ctx, &posreq)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Positions response:")
	for _, m := range posresp.Money {
		fmt.Printf("\t%v: %v\n", m.Currency, SCurPrice(Quotation{Units: m.Units, Nano: m.Nano}))
	}
	for _, mon := range posresp.Money {
		blc.Currencies = append(blc.Currencies, MoneyValue{
			Currency: mon.Currency,
			Units:    mon.Units,
			Nano:     mon.Nano,
		})
	}
}
