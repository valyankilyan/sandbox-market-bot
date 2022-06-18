package tinkoff

import (
	"fmt"

	"gitlab.ozon.dev/valyankilyan/homework-2-market-bot/internal/tinkoff/investapi"
)

func GetLastPrice(figi string) Quotation {
	ctx, cancel := defTink.getContext()
	defer cancel()

	conn, err := defTink.getClientConn()
	if err != nil {
		fmt.Println(err.Error())
		return Quotation{}
	}
	client := investapi.NewMarketDataServiceClient(conn)
	lpreq := investapi.GetLastPricesRequest{
		Figi: []string{figi},
	}

	resp, err := client.GetLastPrices(ctx, &lpreq)
	if err != nil {
		fmt.Println(err.Error())
		return Quotation{}
	}

	return Quotation{
		Units: resp.LastPrices[0].Price.Units,
		Nano:  resp.LastPrices[0].Price.Nano,
	}
}
