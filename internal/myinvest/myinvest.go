package myinvest

import (
	"github.com/valyankilyan/sandbox-market-bot/internal/myinvest/invest_client"
	pb "github.com/valyankilyan/sandbox-market-bot/pkg/myinvestapi"
)

type myinvestServer struct {
	defClient *invest_client.InvestClient
	pb.MyInvestServiceClient
	pb.UnimplementedMyInvestServiceServer
}

func New(ic *invest_client.InvestClient) *myinvestServer {
	return &myinvestServer{defClient: ic}
}
