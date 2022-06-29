package myinvest_client

import (
	"log"

	"github.com/valyankilyan/sandbox-market-bot/internal/bot"
	pb "github.com/valyankilyan/sandbox-market-bot/pkg/myinvestapi"
)

func (cl *InvestClient) PayIn(token string, quantity bot.Quotation) (bot.Quotation, error) {
	req := pb.PayinRequest{
		Token: &pb.TinkoffToken{
			Token: token,
		},
		Quantity: &pb.Quotation{
			Units: quantity.Units,
			Nano:  quantity.Nano,
		},
	}
	balance, err := cl.client.PayIn(cl.ctx, &req)
	if err != nil {
		log.Println("Something went wrong trying to payin", err)
		return bot.Quotation{}, err
	}

	return bot.Quotation{
		Units: balance.Units,
		Nano:  balance.Nano,
	}, nil
}
