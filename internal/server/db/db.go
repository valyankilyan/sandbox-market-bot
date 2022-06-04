package db

import (
	"log"

	"gitlab.ozon.dev/valyankilyan/homework-2-market-bot/internal/server/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const dsn = "host=127.0.0.1 port=5432 user=postgres password=postgres dbname=marketbotdb sslmode=disable"

func New() (*gorm.DB, error) {
	adp, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	adp.AutoMigrate(&models.User{})

	return adp, nil
}

func AddUser(adp *gorm.DB, usr *models.User) {
	log.Printf("adding new user: %v\n", *usr)
	res := adp.Create(usr)
	log.Printf("res = %v\n", res)
}
