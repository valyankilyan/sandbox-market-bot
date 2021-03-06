package db

import (
	"fmt"
	"log"

	"github.com/valyankilyan/sandbox-market-bot/config"
	"github.com/valyankilyan/sandbox-market-bot/internal/server/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%v port=%v user=%v password=%v dbname=marketbotdb sslmode=disable",
		config.Database.Host,
		config.Database.Port,
		config.Database.Username,
		config.Database.Password,
	)
	adp, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	adp.AutoMigrate(&models.User{})

	return adp, nil
}

// func New(ctx context.Context) (*pgxpool.Pool, error) {
// 	return pgxpool.Connect(ctx, dsn)
// }

func AddUser(adp *gorm.DB, usr *models.User) {
	log.Printf("adding new user: %v\n", *usr)
	res := adp.Create(usr)
	log.Printf("res = %v\n", res)
}
