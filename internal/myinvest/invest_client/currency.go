package invest_client

import (
	"log"
	"sync"

	"github.com/valyankilyan/sandbox-market-bot/internal/myinvest"
)

type InvestCurrencyProcessor struct {
	token string
	list  []myinvest.Currency
	figi  []string
	lock  sync.Mutex
}

func NewInvestCurrencyProcessor(token string) *InvestCurrencyProcessor {
	log.Println("New investcurrencyproccessor")
	return &InvestCurrencyProcessor{token: token}
	// ctx, cancel := cp.client.getContext()
	// defer cancel()

	// conn, err := defTink.getClientConn()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// client := investapi.NewInstrumentsServiceClient(conn)

	// ir := investapi.InstrumentsRequest{
	// 	InstrumentStatus: investapi.InstrumentStatus_INSTRUMENT_STATUS_UNSPECIFIED,
	// }

	// resp, err := client.Currencies(ctx, &ir)

	// if err != nil {
	// 	fmt.Printf("currency init %v\n", err.Error())
	// 	os.Exit(3)
	// }

	// c.List = make([]currency, 0)
	// c.figi = make([]string, 0)

	// for _, cur := range resp.Instruments {
	// 	if cur.IsoCurrencyName != "rub" {
	// 		c.List = append(
	// 			c.List,
	// 			currency{
	// 				Name:      cur.Name,
	// 				Shortname: cur.IsoCurrencyName,
	// 				Figi:      cur.Figi,
	// 				Lot:       cur.Lot,
	// 			},
	// 		)
	// 		c.figi = append(c.figi, cur.Figi)
	// 	}
	// }
	// go c.updateCurrencies()
}

func (cp *InvestCurrencyProcessor) All() []myinvest.Currency {
	log.Println("All currencies from icp")
	return []myinvest.Currency{}
}
func (cp *InvestCurrencyProcessor) Get(curlist []string) []myinvest.Currency {
	log.Println("Currencies", curlist, "from icp")
	return []myinvest.Currency{}
}

func (cp *InvestCurrencyProcessor) updateCurrencies() {
	// for {
	// 	ctx, cancel := defTink.getContext()

	// 	conn, err := defTink.getClientConn()
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 		return
	// 	}
	// 	client := investapi.NewMarketDataServiceClient(conn)
	// 	lpreq := investapi.GetLastPricesRequest{
	// 		Figi: c.figi,
	// 	}

	// 	resp, err := client.GetLastPrices(ctx, &lpreq)

	// 	if err != nil {
	// 		fmt.Printf("currency init %v\n", err.Error())
	// 		os.Exit(3)
	// 	}

	// 	c.lock.Lock()
	// 	for i := range resp.LastPrices {
	// 		if resp.LastPrices[i].Price != nil {
	// 			c.List[i].Units = resp.LastPrices[i].Price.Units
	// 			c.List[i].Nano = resp.LastPrices[i].Price.Nano
	// 		}
	// 	}
	// 	c.lock.Unlock()

	// 	cancel()
	// 	time.Sleep(time.Duration(config.Tinkoff.UpdateTime) * time.Second)
	// }
}
