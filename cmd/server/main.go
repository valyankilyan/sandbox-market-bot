package main

import (
	"fmt"
	"log"
	"os"

	"gitlab.ozon.dev/valyankilyan/homework-2-market-bot/internal/server/db"
)

func main() {
	apd, err := db.New()
	if err != nil {
		log.Println(err)
		os.Exit(0)
	}

	fmt.Println(apd)
}
