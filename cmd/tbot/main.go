package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/valyankilyan/sandbox-market-bot/config"
	"github.com/valyankilyan/sandbox-market-bot/internal/telegram"
	mi "github.com/valyankilyan/sandbox-market-bot/internal/telegram/myinvest_client"
	sc "github.com/valyankilyan/sandbox-market-bot/internal/telegram/server_client"
	mipb "github.com/valyankilyan/sandbox-market-bot/pkg/myinvestapi"
	srvpb "github.com/valyankilyan/sandbox-market-bot/pkg/server_api"
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

	srvconn := serverConnection()
	defer srvconn.Close()
	srvclient := srvpb.NewUserServiceClient(srvconn)

	invconn := myinvestConnection()
	defer invconn.Close()
	invclient := mipb.NewMyInvestServiceClient(invconn)
	ctx := context.Background()
	ctx = metadata.AppendToOutgoingContext(ctx,
		"sender", "TelegramHandler",
		"when", time.Now().Format(time.RFC3339),
	)

	server_client := sc.New(ctx, srvclient)
	invest_client := mi.New(ctx, invclient)

	bot := telegram.New(string(config.Telegram.Token), server_client, invest_client)
	msgch := make(chan telegram.Message, 20)
	go bot.GetUpdates(msgch)
	go bot.HandleMessages(msgch)

	for {
		time.Sleep(time.Hour)
	}
}

func serverConnection() *grpc.ClientConn {
	if server_host := os.Getenv("SERVER_HOST"); server_host != "" {
		config.Rpc.Host = server_host
	}
	config.Rpc.Host = config.Rpc.Host + ":" + config.Rpc.Port

	conn, err := grpc.Dial(config.Rpc.Host,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("can't make grpc conection with server_host", err)
	}

	return conn
}

func myinvestConnection() *grpc.ClientConn {
	if myinvest_host := os.Getenv("MYINVEST_HOST"); myinvest_host != "" {
		config.Myinvest.Host = myinvest_host
	}
	config.Myinvest.Host = config.Myinvest.Host + ":" + config.Myinvest.Port

	conn, err := grpc.Dial(config.Myinvest.Host,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("can't make grpc connection with myinvest_host")
	}

	return conn
}
