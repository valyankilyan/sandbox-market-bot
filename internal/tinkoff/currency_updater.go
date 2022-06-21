package tinkoff

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/valyankilyan/sandbox-market-bot/config"
	"github.com/valyankilyan/sandbox-market-bot/internal/tinkoff/investapi"
)

var Currencies CurrencyList

type CurrencyList struct {
	List []Currency
	figi []string
	lock sync.Mutex
}

// func New() *

type Currency struct {
	Name      string
	Shortname string
	Figi      string
	Units     int64
	Nano      int32
	Lot       int32
}

func (c *CurrencyList) InitCurrencies() {
	ctx, cancel := defTink.getContext()
	defer cancel()

	conn, err := defTink.getClientConn()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	client := investapi.NewInstrumentsServiceClient(conn)

	ir := investapi.InstrumentsRequest{
		InstrumentStatus: investapi.InstrumentStatus_INSTRUMENT_STATUS_UNSPECIFIED,
	}

	resp, err := client.Currencies(ctx, &ir)

	if err != nil {
		fmt.Printf("currency init %v\n", err.Error())
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
					Lot:       cur.Lot,
				},
			)
			c.figi = append(c.figi, cur.Figi)
		}
	}
	go c.updateCurrencies()
}

func (c *CurrencyList) updateCurrencies() {
	for {
		ctx, cancel := defTink.getContext()

		conn, err := defTink.getClientConn()
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
			fmt.Printf("currency init %v\n", err.Error())
			os.Exit(3)
		}

		c.lock.Lock()
		for i := range resp.LastPrices {
			if resp.LastPrices[i].Price != nil {
				c.List[i].Units = resp.LastPrices[i].Price.Units
				c.List[i].Nano = resp.LastPrices[i].Price.Nano
			}
		}
		c.lock.Unlock()

		cancel()
		time.Sleep(time.Duration(config.Conf.Tinkoff.UpdateTime) * time.Second)
	}
}

func SCurPrice(q Quotation) string {
	if q.Nano != 0 {
		return fmt.Sprintf("%v.%v", q.Units, SCurNano(q.Nano))
	} else {
		return fmt.Sprintf("%v", q.Units)
	}
}

func SCurNano(nano int32) string {
	ans := fmt.Sprintf("%d", nano)
	ans = strings.TrimRight(ans, "0")
	return ans
}

func GetCurrency(shortname string) (Currency, bool) {
	shortname = strings.ToLower(shortname)
	for _, cur := range Currencies.List {
		fmt.Println(cur.Shortname)
	}
	for _, cur := range Currencies.List {
		if cur.Shortname == shortname {
			return cur, true
		}
	}
	return Currency{}, false
}
