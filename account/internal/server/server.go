package server

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/vmustillo/microservices/account/gen"
)

func (*AccountServer) GetAccount(ctx context.Context, req *gen.GetAccountRequest) (*gen.Account, error) {
	return &gen.Account{
		Id:      "guid1",
		OwnerID: "ownerID1",
		Balance: 24560.75,
		Type:    "Checking",
	}, nil
}

func (*AccountServer) CreateAccount(ctx context.Context, req *gen.CreateAccountRequest) (*gen.CreateAccountResponse, error) {
	return &gen.CreateAccountResponse{}, nil
}

func NewAccountServer(actSrv *AccountServer) *grpc.Server {

	srv := grpc.NewServer()
	gen.RegisterAccountServiceServer(srv, actSrv)
	reflection.Register(srv)

	return srv
}
