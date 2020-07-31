package main

import (
	"context"
	"flag"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	gw "github.com/vmustillo/microservices/account/gen"
)

var (
	echoEndpoint = flag.String("echo_endpoint", "localhost:9090", "endpoint of account")
)

type AccountServer struct{}

func (*AccountServer) GetAccount(ctx context.Context, req *gw.GetAccountRequest) (*gw.Account, error) {
	return &gw.Account{
		Id:      "guid1",
		OwnerID: "ownerID1",
		Balance: 24560.75,
		Type:    "Checking",
	}, nil
}

func newAccountServer() *grpc.Server {

	srv := grpc.NewServer()
	gw.RegisterAccountServiceServer(srv, &AccountServer{})
	reflection.Register(srv)

	return srv
}

func main() {

	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	accountService := newAccountServer()
	fmt.Println("Listening on port 4040")
	if e := accountService.Serve(listener); e != nil {
		panic(e)
	}
}
