package main

import (
	"fmt"
	"log"
	"os"

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
	fmt.Println(config.Conf)

	bot := telegram.New(string(config.Conf.Telegram.Token))
	var user config.User
	for _, u := range config.Conf.Users {
		user = u
	}
	fmt.Println("USER:", user)
	bot.SendMessage(user.TgId, "Hello from Go")
	go bot.GetUpdates()
	// body, _ := bot.GetUpdates()
	// bot.GetMessages(body)
}
