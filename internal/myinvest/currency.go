package myinvest

import (
	"context"
	"log"

	pb "github.com/valyankilyan/sandbox-market-bot/pkg/myinvestapi"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (srv *myinvestServer) GetCurrencies(ctx context.Context,
	req *pb.CurrencyRequest) (*pb.CurrencyList, error) {

	currencies := srv.curProcessor.Get(req.ShortName)

	list := make([]*pb.Currency, len(currencies))
	for i, c := range currencies {
		list[i] = &pb.Currency{
			ShortName: c.ShortName,
			Name:      c.Name,
			Price:     &pb.Quotation{Units: c.Price.Units, Nano: c.Price.Nano},
			Lot:       c.Lot,
		}
	}
	log.Println("getcurrencies", req.ShortName)
	return &pb.CurrencyList{Currencies: list}, nil
}

func (srv *myinvestServer) GetAllCurrencies(ctx context.Context, empty *emptypb.Empty) (*pb.CurrencyList, error) {
	// invest_client.
	currencies := srv.curProcessor.All()
	list := make([]*pb.Currency, len(currencies))
	for i, c := range currencies {
		list[i] = &pb.Currency{
			ShortName: c.ShortName,
			Name:      c.Name,
			Price:     &pb.Quotation{Units: c.Price.Units, Nano: c.Price.Nano},
			Lot:       c.Lot,
		}
	}
	log.Println("getallcurrencies")

	return &pb.CurrencyList{Currencies: list}, nil
}
