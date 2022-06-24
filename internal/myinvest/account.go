package myinvest

import (
	"context"
	"log"

	pb "github.com/valyankilyan/sandbox-market-bot/pkg/myinvestapi"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (srv *myinvestServer) PayIn(ctx context.Context,
	req *pb.PayinRequest) (*emptypb.Empty, error) {
	log.Println("payin req", req.Token, req.Quantity)

	return &emptypb.Empty{}, nil
}
