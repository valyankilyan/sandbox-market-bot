package tinkoff

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"gitlab.ozon.dev/valyankilyan/homework-2-market-bot/config"
	"gitlab.ozon.dev/valyankilyan/homework-2-market-bot/internal/tinkoff/investapi"
)

var Currencies CurrencyList

type CurrencyList struct {
	List []Currency
	figi []string
	lock sync.Mutex
	tink *Tinkoff
}

// func New() *

type Currency struct {
	Name      string
	Shortname string
	Figi      string
	Units     int64
	Nano      int32
}

// func init() {
// 	fmt.Println("CurrenciesInited")
// 	fmt.Println("tink added")
// 	Currencies.initCurrencies()
// 	fmt.Println("Currencies inited")
// }

func (c *CurrencyList) InitCurrencies() {
	c.tink = New(config.Conf.Tinkoff.DefaultToken)
	ctx, cancel := c.tink.getContext()
	defer cancel()

	conn, err := c.tink.getClientConn()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	client := investapi.NewInstrumentsServiceClient(conn)

	// is = investapi.InstrumentStatus()
	ir := investapi.InstrumentsRequest{
		InstrumentStatus: investapi.InstrumentStatus_INSTRUMENT_STATUS_UNSPECIFIED,
	}

	resp, err := client.Currencies(ctx, &ir)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(3)
	}

	c.List = make([]Currency, 0)
	c.figi = make([]string, 0)

	for _, cur := range resp.Instruments {
		if cur.IsoCurrencyName != "rub" {
			c.List = append(
				c.List,
				Currency{
					Name:      cur.Name,
					Shortname: cur.IsoCurrencyName,
					Figi:      cur.Figi,
				},
			)

			c.figi = append(c.figi, cur.Figi)
		}
	}
	go c.updateCurrencies()
}

func (c *CurrencyList) updateCurrencies() {
	for {
		ctx, cancel := c.tink.getContext()

		conn, err := c.tink.getClientConn()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		client := investapi.NewMarketDataServiceClient(conn)
		lpreq := investapi.GetLastPricesRequest{
			Figi: c.figi,
		}

		resp, err := client.GetLastPrices(ctx, &lpreq)

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(3)
		}

		fmt.Printf("cur = %v, resp = %v", len(c.List), len(resp.LastPrices))

		c.lock.Lock()
		for i := range resp.LastPrices {
			if resp.LastPrices[i].Price != nil {
				c.List[i].Units = resp.LastPrices[i].Price.Units
				c.List[i].Nano = resp.LastPrices[i].Price.Nano
			}
			fmt.Printf("%v: %v.%v rub\n", c.List[i].Name, c.List[i].Units, SCurNano(c.List[i].Nano))
		}
		c.lock.Unlock()

		cancel()
		time.Sleep(time.Duration(config.Conf.Tinkoff.UpdateTime) * time.Second)
	}
}

func SCurNano(nano int32) string {
	ans := fmt.Sprintf("%d", nano)
	ans = strings.TrimRight(ans, "0")
	return ans
}
