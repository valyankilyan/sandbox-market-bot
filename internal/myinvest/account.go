package myinvest

import (
	"context"
	"log"

	pb "github.com/valyankilyan/sandbox-market-bot/pkg/myinvestapi"
)

func (srv *myinvestServer) PayIn(ctx context.Context,
	req *pb.PayinRequest) (*pb.Quotation, error) {
	quantity, _ := srv.client.PayIn(req.Token.Token, Quotation{req.Quantity.Units, req.Quantity.Nano})
	log.Println("payin req", req.Token, req.Quantity)
	return &pb.Quotation{Units: quantity.Units, Nano: quantity.Nano}, nil
}
