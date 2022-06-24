package myinvest

import (
	"context"
	"log"

	pb "github.com/valyankilyan/sandbox-market-bot/pkg/myinvestapi"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (srv *myinvestServer) UpdateTinkoffToken(ctx context.Context,
	req *pb.TinkoffToken) (*emptypb.Empty, error) {
	// invest_client.
	log.Println("update tinkoff token", req.Token)

	return &emptypb.Empty{}, nil
}

func (srv *myinvestServer) PayIn(ctx context.Context,
	req *pb.PayinRequest) (*emptypb.Empty, error) {
	log.Println("payin req", req.Token, req.Quantity)

	return &emptypb.Empty{}, nil
}
