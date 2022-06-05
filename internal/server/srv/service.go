package srv

import (
	pb "gitlab.ozon.dev/valyankilyan/homework-2-market-bot/pkg/api"
)

type tserver struct {
	lastId int64
	dbHand DBHandler
	pb.MarketBotClient
	pb.UnimplementedMarketBotServer
}

func New(dbHand DBHandler) *tserver {
	return &tserver{dbHand: dbHand}
}

// func (t *tserver) mustEmbedUnimplementedMarketBotServer() {
// 	log.Println("unimplemented")
// }
