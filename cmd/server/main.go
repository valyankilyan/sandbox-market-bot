package main

import (
	"log"
	"net"
	"os"
	"time"

	"gitlab.ozon.dev/valyankilyan/homework-2-market-bot/config"
	"gitlab.ozon.dev/valyankilyan/homework-2-market-bot/internal/server/db"
	"gitlab.ozon.dev/valyankilyan/homework-2-market-bot/internal/server/mw"
	"gitlab.ozon.dev/valyankilyan/homework-2-market-bot/internal/server/srv"
	pb "gitlab.ozon.dev/valyankilyan/homework-2-market-bot/pkg/api"
	"google.golang.org/grpc"
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
	// ctx := context.Background()
	adp, err := db.New()
	if err != nil {
		log.Println(err)
		os.Exit(0)
	}
	newServer := srv.New(adp)
	lis, err := net.Listen("tcp", config.Conf.Rpc.Host)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(mw.LogInterceptor),
	}

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterMarketBotServer(grpcServer, newServer)
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
	for {
		time.Sleep(time.Second)
	}
}
