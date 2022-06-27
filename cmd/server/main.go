package main

import (
	"log"
	"net"
	"os"
	"time"

	"github.com/valyankilyan/sandbox-market-bot/config"
	"github.com/valyankilyan/sandbox-market-bot/internal/server/db"
	"github.com/valyankilyan/sandbox-market-bot/internal/server/mw"
	"github.com/valyankilyan/sandbox-market-bot/internal/server/srv"
	pb "github.com/valyankilyan/sandbox-market-bot/pkg/server_api"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Starting USER SERVER")
	b, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	err = config.ParseConfig(b)
	if err != nil {
		log.Fatal(err)
	}
	if db_host := os.Getenv("DB_HOST"); db_host != "" {
		config.Database.Host = db_host
	}
	if server_host := os.Getenv("SERVER_HOST"); server_host != "" {
		config.Rpc.Host = server_host
	}
	config.Rpc.Host = config.Rpc.Host + ":" + config.Rpc.Port

	adp, err := db.New()
	if err != nil {
		log.Println(err)
		os.Exit(0)
	}
	newServer := srv.New(adp)
	lis, err := net.Listen("tcp", config.Rpc.Host)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(mw.LogInterceptor),
	}

	log.Println("User Service listens at", config.Rpc.Host)
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterUserServiceServer(grpcServer, newServer)
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
	for {
		time.Sleep(time.Second)
	}
}
