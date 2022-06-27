package myinvest_client

import (
	"log"

	"github.com/valyankilyan/sandbox-market-bot/internal/telegram"
	pb "github.com/valyankilyan/sandbox-market-bot/pkg/myinvestapi"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (cl *InvestClient) AllCurrencies() ([]telegram.Currency, error) {
	curs, err := cl.client.GetAllCurrencies(cl.ctx, &emptypb.Empty{})
	if err != nil {
		log.Println("error while trying to get all currencies", err)
		return []telegram.Currency{}, err
	}

	return pbCurToTgCur(curs), nil
}

func (cl *InvestClient) Currencies(shortnames []string) ([]telegram.Currency, error) {
	req := pb.CurrencyRequest{ShortName: shortnames}
	curs, err := cl.client.GetCurrencies(cl.ctx, &req)

	if err != nil {
		log.Printf("error while trying to get %v currencies: %v", shortnames, err)
		return []telegram.Currency{}, err
	}

	return pbCurToTgCur(curs), nil
}

func pbCurToTgCur(curs *pb.CurrencyList) []telegram.Currency {
	tcurs := make([]telegram.Currency, len(curs.Currencies))
	for i, c := range curs.Currencies {
		tcurs[i] = telegram.Currency{
			Price:     telegram.Quotation{Units: c.Price.Units, Nano: c.Price.Nano},
			ShortName: c.ShortName,
			Name:      c.Name,
			Lot:       c.Lot,
		}
	}
	return tcurs
}
