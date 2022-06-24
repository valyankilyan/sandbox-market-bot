package myinvest

import (
	"context"
	"log"

	pb "github.com/valyankilyan/sandbox-market-bot/pkg/myinvestapi"
)

func (srv *myinvestServer) GetCurrencies(ctx context.Context,
	req *pb.CurrencyRequest) (*pb.CurrencyList, error) {
	// invest_client.
	log.Println("getcurrencies", req.ShortName)

	return &pb.CurrencyList{}, nil
}
