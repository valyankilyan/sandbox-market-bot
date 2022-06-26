package invest_client

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/valyankilyan/sandbox-market-bot/config"
	"github.com/valyankilyan/sandbox-market-bot/internal/myinvest"
	"github.com/valyankilyan/sandbox-market-bot/pkg/investapi"
)

type InvestCurrencyProcessor struct {
	token      string
	curlist    []myinvest.Currency
	curfigi    []string
	lastUpdate time.Time
	lock       sync.Mutex
}

func NewInvestCurrencyProcessor(token string) *InvestCurrencyProcessor {
	cp := InvestCurrencyProcessor{token: token}
	cp.initCurrencies()
	log.Println("New investcurrencyproccessor")
	return &cp
}

func (cp *InvestCurrencyProcessor) All() []myinvest.Currency {
	cp.updateCurrencies()
	return cp.curlist
}

func (cp *InvestCurrencyProcessor) Get(curnames []string) []myinvest.Currency {
	cp.updateCurrencies()
	resp := []myinvest.Currency{}

	for _, name := range curnames {
		for _, cur := range cp.curlist {
			if name == cur.ShortName {
				resp = append(resp, cur)
			}
		}
	}
	return resp
}

func (cp *InvestCurrencyProcessor) initCurrencies() {
	log.Println("All currencies from icp")
	ctx, cancel := getContext(cp.token)
	defer cancel()

	conn, err := getClientConn()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(3)
	}

	client := investapi.NewInstrumentsServiceClient(conn)

	ir := investapi.InstrumentsRequest{
		InstrumentStatus: investapi.InstrumentStatus_INSTRUMENT_STATUS_UNSPECIFIED,
	}

	resp, err := client.Currencies(ctx, &ir)

	if err != nil {
		fmt.Printf("curProcessor All %v\n", err.Error())
		os.Exit(3)
	}

	// cp.curlist = make([]myinvest.Currency, len(resp.Instruments))
	// cp.curfigi = make([]string, len(resp.Instruments))
	for _, c := range resp.Instruments {
		if c.Currency != "rub" || c.IsoCurrencyName == "rub" {
			continue
		}
		cp.curlist = append(cp.curlist, myinvest.Currency{
			Name:      c.Name,
			ShortName: c.IsoCurrencyName,
			Figi:      c.Figi,
			Lot:       c.Lot,
		})
		cp.curfigi = append(cp.curfigi, c.Figi)
	}
	cp.updateCurrencies()
}

func (cp *InvestCurrencyProcessor) updateCurrencies() {
	fmt.Println("now - last update", time.Now().Sub(cp.lastUpdate))
	fmt.Println("updatetime", time.Duration(config.Tinkoff.UpdateTime)*time.Second)
	fmt.Println(time.Now().Sub(cp.lastUpdate) > time.Duration(config.Tinkoff.UpdateTime)*time.Second)
	if !(time.Now().Sub(cp.lastUpdate) > time.Duration(config.Tinkoff.UpdateTime)*time.Second) {
		return
	}
	log.Println("updateCurrencies")
	cp.lastUpdate = time.Now()

	ctx, cancel := getContext(cp.token)
	defer cancel()

	conn, err := getClientConn()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(3)
	}
	client := investapi.NewMarketDataServiceClient(conn)
	lpreq := investapi.GetLastPricesRequest{
		Figi: cp.curfigi,
	}

	resp, err := client.GetLastPrices(ctx, &lpreq)

	if err != nil {
		fmt.Printf("update currencys %v\n", err.Error())
		return
	}

	cp.lock.Lock()
	for i, lp := range resp.LastPrices {
		if resp.LastPrices[i].Price != nil {
			cp.curlist[i].Price = myinvest.Quotation{
				Units: lp.Price.Units,
				Nano:  lp.Price.Nano,
			}
		}
	}
	cp.lock.Unlock()

	cancel()
}
