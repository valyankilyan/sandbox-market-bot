package myinvest

import (
	pb "github.com/valyankilyan/sandbox-market-bot/pkg/myinvestapi"
)

func New(cp CurrencyProcessor, ic InvestClientProcessor) *myinvestServer {
	server := &myinvestServer{
		curProcessor: cp,
	}
	return server
}

type myinvestServer struct {
	curProcessor CurrencyProcessor
	client       InvestClientProcessor
	pb.MyInvestServiceClient
	pb.UnimplementedMyInvestServiceServer
}

type CurrencyProcessor interface {
	All() []Currency
	Get([]string) []Currency
}
type InvestClientProcessor interface {
	PayIn(token string, quantity Quotation) error
}

type Currency struct {
	Name      string
	Shortname string
	Figi      string
	Lot       int32
	Price     Quotation
}
type Quotation struct {
	Units int64
	Nano  int32
}
