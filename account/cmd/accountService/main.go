package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	gw "github.com/vmustillo/microservices/account/gen"
	pkg "github.com/vmustillo/microservices/account/pkg"
)

var (
	echoEndpoint = flag.String("echo_endpoint", "localhost:9090", "endpoint of account")
)

type AccountServer struct {
	Db *sql.DB
}

func (*AccountServer) GetAccount(ctx context.Context, req *gw.GetAccountRequest) (*gw.Account, error) {
	return &gw.Account{
		Id:      "guid1",
		OwnerID: "ownerID1",
		Balance: 24560.75,
		Type:    "Checking",
	}, nil
}

func newAccountServer(actSrv *AccountServer) *grpc.Server {

	srv := grpc.NewServer()
	gw.RegisterAccountServiceServer(srv, actSrv)
	reflection.Register(srv)

	return srv
}

func (s *AccountServer) connectDatabase() error {
	db, err := pkg.GetConnection("postgres", "postgres", "127.0.0.1", 5433, "account")
	if err != nil {
		log.Println("Error connecting to database", err)

		return err
	}

	s.Db = db

	return nil
}

func main() {

	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	a := &AccountServer{}
	a.connectDatabase()
	err = a.Db.Ping()
	if err != nil {
		log.Println("ERROR", err)
	} else {
		log.Println("SUCCESS")
	}

	accountService := newAccountServer(a)
	fmt.Println("Listening on port 4040")
	if e := accountService.Serve(listener); e != nil {
		panic(e)
	}
}
