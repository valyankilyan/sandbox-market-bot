package myinvest

import "github.com/valyankilyan/sandbox-market-bot/internal/myinvest/invest_client"

type myinvestServer struct {
	defClient *invest_client.InvestClient
}

func New(ic *invest_client.InvestClient) *myinvestServer {
	return &myinvestServer{defClient: ic}
}
