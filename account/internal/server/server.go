package server

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/vmustillo/microservices/account/gen"
)

func (s *AccountServer) GetAccount(ctx context.Context, req *gen.GetAccountRequest) (*gen.Account, error) {
	var acc gen.Account

	log.Println("Getting Account with ID:", req.GetId())
	checkStmt := `select id, user_id, balance, credit_limit from accounts where id = $1`
	err := s.Db.QueryRow(checkStmt, req.GetId()).Scan(&acc.Id, &acc.OwnerID, &acc.Balance, &acc.CreditLimit)
	if err != nil {
		return &gen.Account{}, err
	}

	return &acc, nil
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
