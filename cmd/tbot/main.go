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

	c, err := config.ParseConfig(b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c)

	bot := telegram.New(string(c.TgToken))
	var user config.User
	for _, u := range c.Users {
		user = u
	}
	fmt.Println("USER:", user)
	bot.SendMessage(user.TgId, "Hello from Go")
	bot.GetUpdates()
}
