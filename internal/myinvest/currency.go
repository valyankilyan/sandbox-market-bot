package myinvest

import (
	"context"
	"log"

	pb "github.com/valyankilyan/sandbox-market-bot/pkg/myinvestapi"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (srv *myinvestServer) GetCurrencies(ctx context.Context,
	req *pb.CurrencyRequest) (*pb.CurrencyList, error) {
	// invest_client.

	log.Println("getcurrencies", req.ShortName)

	return &pb.CurrencyList{}, nil
}

func (srv *myinvestServer) GetAllCurrencies(ctx context.Context, empty *emptypb.Empty) (*pb.CurrencyList, error) {
	// invest_client.

	log.Println("getallcurrencies")

	return &pb.CurrencyList{}, nil
}
