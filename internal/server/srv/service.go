package srv

import (
	pb "gitlab.ozon.dev/valyankilyan/homework-2-market-bot/pkg/api"
	"gorm.io/gorm"
)

type tserver struct {
	// lastId int64
	// repo   Repository
	db *gorm.DB
	pb.MarketBotClient
	pb.UnimplementedMarketBotServer
}

func New(db *gorm.DB) *tserver {
	return &tserver{db: db}
}

// func (t *tserver) mustEmbedUnimplementedMarketBotServer() {
// 	log.Println("unimplemented")
// }
