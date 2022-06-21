package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/valyankilyan/sandbox-market-bot/config"
	"github.com/valyankilyan/sandbox-market-bot/internal/telegram"
	"github.com/valyankilyan/sandbox-market-bot/internal/tinkoff"
	"github.com/valyankilyan/sandbox-market-bot/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
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

	if server_host := os.Getenv("SERVER_HOST"); server_host != "" {
		config.Conf.Rpc.Host = server_host
	}
	config.Conf.Rpc.Host = config.Conf.Rpc.Host + ":" + config.Conf.Rpc.Port

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
	tinkoff.DefTinkInit()
	tinkoff.Currencies.InitCurrencies()
	bot.GetUpdates()
}
