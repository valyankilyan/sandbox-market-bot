package myinvest_client

import (
	"context"

	pb "github.com/valyankilyan/sandbox-market-bot/pkg/myinvestapi"
)

type InvestClient struct {
	client pb.MyInvestServiceClient
	ctx    context.Context
}

func New(ctx context.Context, client pb.MyInvestServiceClient) *InvestClient {
	return &InvestClient{
		client: client,
		ctx:    ctx,
	}
}
