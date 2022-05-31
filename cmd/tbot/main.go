package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gitlab.ozon.dev/valyankilyan/homework-2-market-bot/config"
	"gitlab.ozon.dev/valyankilyan/homework-2-market-bot/internal/app/telegram"
)

func main() {
	b, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	err = config.ParseConfig(b)
	if err != nil {
		log.Fatal(err)
	}

	bot := telegram.New(string(config.Conf.Telegram.Token))
	go bot.GetUpdates()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// body, _ := bot.GetUpdates()
	// bot.GetMessages(body)
	time.Sleep(100 * time.Second)
}
