package main

import (
	"fmt"
	"log"
	"os"

	"gitlab.ozon.dev/valyankilyan/homework-2-market-bot/internal/server/db"
	"gitlab.ozon.dev/valyankilyan/homework-2-market-bot/internal/server/models"
)

func main() {
	adp, err := db.New()
	if err != nil {
		log.Println(err)
		os.Exit(0)
	}

	// usr := models.User{
	// 	Name:         "test",
	// 	TgUserName:   "username",
	// 	TgId:         12345,
	// 	TinkoffToken: "jsld;kgj;sldkfgj",
	// }

	// db.AddUser(adp, &usr)

	var usr1 models.User
	adp.First(&usr1)
	fmt.Printf("usr founded: %v\n", usr1)

	fmt.Println(adp)
}
