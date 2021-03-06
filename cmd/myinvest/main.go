package main

import (
	"log"
	"net"
	"os"
	"time"

	"github.com/valyankilyan/sandbox-market-bot/config"
	"github.com/valyankilyan/sandbox-market-bot/internal/myinvest"
	"github.com/valyankilyan/sandbox-market-bot/internal/myinvest/invest_client"
	"github.com/valyankilyan/sandbox-market-bot/internal/server/mw"
	"github.com/valyankilyan/sandbox-market-bot/pkg/myinvestapi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	log.Println("starting MYINVEST")
	//parse config
	b, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		log.Fatal("can't open config file", err)
	}

	err = config.ParseConfig(b)
	if err != nil {
		log.Fatal("can't parse config file", err)
	}

	// change myinvest_host if it is in env
	if myinvest_host := os.Getenv("MYINVEST_HOST"); myinvest_host != "" {
		config.Myinvest.Host = myinvest_host
	}
	config.Myinvest.Host = config.Myinvest.Host + ":" + config.Myinvest.Port
	log.Println("MyinvestServer listents on", config.Myinvest.Host)

	// init CurrencyProcessor
	currency_processor :=
		invest_client.NewInvestCurrencyProcessor(config.Tinkoff.DefaultToken)

	// init InvestClientProcessor
	invest_client_processor := invest_client.NewInvestClient()

	myinvestServer := myinvest.New(currency_processor, invest_client_processor)
	listener, err := net.Listen("tcp", config.Myinvest.Host)
	if err != nil {
		panic(err)
	}

	// make grpc server
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(mw.LogInterceptor),
	}

	grpcServer := grpc.NewServer(opts...)
	myinvestapi.RegisterMyInvestServiceServer(grpcServer, myinvestServer)
	err = grpcServer.Serve(listener)
	if err != nil {
		panic(err)
	}
	for {
		time.Sleep(time.Second)
	}
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
