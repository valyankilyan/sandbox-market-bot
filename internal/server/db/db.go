package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	"gitlab.ozon.dev/valyankilyan/homework-2-market-bot/internal/server/models"
	"gorm.io/gorm"
)

const dsn = "host=127.0.0.1 port=5432 user=postgres password=postgres dbname=marketbotdb sslmode=disable"

// func New() (*gorm.DB, error) {
// 	adp, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Println(err)
// 		return nil, err
// 	}
// 	adp.AutoMigrate(&models.User{})

// 	return adp, nil
// }

func New(ctx context.Context) (*pgxpool.Pool, error) {
	return pgxpool.Connect(ctx, dsn)
}

func AddUser(adp *gorm.DB, usr *models.User) {
	log.Printf("adding new user: %v\n", *usr)
	res := adp.Create(usr)
	log.Printf("res = %v\n", res)
}
