package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"gitlab.ozon.dev/valyankilyan/homework-2-market-bot/internal/server/db"
	"gitlab.ozon.dev/valyankilyan/homework-2-market-bot/internal/server/dbhand"
	"gitlab.ozon.dev/valyankilyan/homework-2-market-bot/internal/server/srv"
	pb "gitlab.ozon.dev/valyankilyan/homework-2-market-bot/pkg/api"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	adp, err := db.New(ctx)
	if err != nil {
		log.Println(err)
		os.Exit(0)
	}
	newServer := srv.New(dbhand.New(adp))
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// var opts []grpc.ServerOption
	// opts = []grpc.ServerOption{
	// 	grpc.UnaryInterceptor(mw.LogInterceptor),
	// }

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	grpcServer := grpc.NewServer()
	pb.RegisterMarketBotServer(grpcServer, newServer)
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
	for {
		time.Sleep(time.Second)
	}
}
