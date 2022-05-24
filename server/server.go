package server

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"playground/api"
	"sync"

	"google.golang.org/grpc"
)

var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")
	jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of features")
	port       = flag.Int("port", 50051, "The server port")
)

var sometype api.SomeType

type relativeInformationServer struct {
	persons *[]api.Person

	api.UnimplementedRelativeInformationServer

	mu sync.Mutex // protects routeNotes
	// GetRelativesData(context.Context, *SomeRequest) (*SomeResponse, error)
	// mustEmbedUnimplementedRelativeInformationServer()
}

func (s *relativeInformationServer) load(filePath string) {
	var data []byte
	if filePath != "" {
		var err error
		data, err = ioutil.ReadFile(filePath)
		if err != nil {
			log.Fatalf("Failed to load default features: %v", err)
		}
	} else {
		data = exampleData
	}
	if err := json.Unmarshal(data, &s.persons); err != nil {
		log.Fatalf("Failed to load default features: %v", err)
	}
}

func (r *relativeInformationServer) GetRelativesData(ctx context.Context, sm *api.SomeRequest) (*api.SomeResponse, error) {
	fmt.Println(sm)
	sr := api.SomeResponse{Id: int64(sm.Id), Unix: 12345, Text: []string{"hi", "go", "fuck", "urself"}}
	return &sr, nil
}

func newServer() *relativeInformationServer {
	s := &relativeInformationServer{}
	s.load(*jsonDBFile)
	return s
}

func StartServer() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	api.RegisterRelativeInformationServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}

var exampleData = []byte(`[{
    "Person": {
        "UserId":   1
		"Nickname": "Nick"
		"Name":     "Nickolai"
		"Resource": "Tinkoff"
		"Message":  "HELLO I'm Nick"
    },
	"Person": {
        "UserId":   2
		"Nickname": "Nick2"
		"Name":     "Nickolai2"
		"Resource": "Tinkoff"
		"Message":  "HELLO I'm Nick2"
}]`)
