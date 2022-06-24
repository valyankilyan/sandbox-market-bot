package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/valyankilyan/sandbox-market-bot/config"
	"github.com/valyankilyan/sandbox-market-bot/internal/telegram"
	sc "github.com/valyankilyan/sandbox-market-bot/internal/telegram/server_client"
	srv "github.com/valyankilyan/sandbox-market-bot/pkg/server_api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {
	b, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		log.Fatal("can't open config file", err)
	}

	err = config.ParseConfig(b)
	if err != nil {
		log.Fatal("can't parse config file", err)
	}

	conn := connection()
	defer conn.Close()

	client := srv.NewUserServiceClient(conn)
	ctx := context.Background()
	ctx = metadata.AppendToOutgoingContext(ctx,
		"sender", "TelegramHandler",
		"when", time.Now().Format(time.RFC3339),
	)

	server_client := sc.New(client, ctx)

	bot := telegram.New(string(config.Telegram.Token), server_client)
	// tinkoff.DefTinkInit()
	// tinkoff.Currencies.InitCurrencies()
	bot.GetUpdates()
}

func connection() *grpc.ClientConn {
	if server_host := os.Getenv("SERVER_HOST"); server_host != "" {
		config.Rpc.Host = server_host
	}
	config.Rpc.Host = config.Rpc.Host + ":" + config.Rpc.Port

	conn, err := grpc.Dial(config.Rpc.Host,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("can't make grpc conection", err)
	}

	return conn
}
