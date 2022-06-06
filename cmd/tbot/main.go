package main

import (
	"context"
	"log"
	"os"
	"time"

	"gitlab.ozon.dev/valyankilyan/homework-2-market-bot/config"
	"gitlab.ozon.dev/valyankilyan/homework-2-market-bot/internal/telegram"
	"gitlab.ozon.dev/valyankilyan/homework-2-market-bot/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"gitlab.ozon.dev/valyankilyan/homework-2-market-bot/internal/app/tinkoff"
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

	conn, err := grpc.Dial(config.Conf.Rpc.Host,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := api.NewMarketBotClient(conn)
	ctx := context.Background()
	ctx = metadata.AppendToOutgoingContext(ctx,
		"sender", "TelegramHandler",
		"when", time.Now().Format(time.RFC3339),
	)

	bot := telegram.New(string(config.Conf.Telegram.Token), client, ctx)
	bot.GetUpdates()
	fmt.Println("token in main =", config.Conf.Users[0].TinkoffToken)
	// body, _ := bot.GetUpdates()
	// bot.GetMessages(body)
	// tinkoff.Not_main(config.Conf.Users[0].TinkoffToken)
	// t := tinkoff.New(config.Conf.Users[0].TinkoffToken)
	// t.AddAccount()
	t := tinkoff.New(config.Conf.Users[0].TinkoffToken)
	t.GetAccounts()
	time.Sleep(100 * time.Second)
}
