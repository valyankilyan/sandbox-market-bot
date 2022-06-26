package invest_client

import (
	"log"

	"github.com/valyankilyan/sandbox-market-bot/internal/myinvest"
)

// "github.com/ofen/tinkoff-invest-example/tinkoff/investapi"

func (ic *InvestClient) PayIn(token string, quantity myinvest.Quotation) (myinvest.Quotation, error) {
	log.Println("Payin from invest_client")
	return quantity, nil
}
