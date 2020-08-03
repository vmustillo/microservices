package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	gen "github.com/vmustillo/microservices/auth/gen"
	pkg "github.com/vmustillo/microservices/auth/pkg"
)

type AuthServer struct {
	Db *sql.DB
}

func (s *AuthServer) GetUsers(ctx context.Context, req *gen.GetUsersRequest) (*gen.Users, error) {
	return &gen.Users{
		Users: []*gen.User{
			{
				Id:        "guid1",
				AccountID: "accountID1",
				FirstName: "Vin",
				LastName:  "M",
			},
		},
	}, nil
}

func (s *AuthServer) CreateUser(ctx context.Context, req *gen.CreateUserRequest) (*gen.CreateUserResponse, error) {

	id, err := pkg.CreateUser(s.Db, req.GetFirstName(), req.GetLastName(), req.GetUsername(), req.GetPassword())
	if err != nil {
		log.Println("Error creating user", err)

		return &gen.CreateUserResponse{
			User: &gen.User{},
		}, err
	}

	return &gen.CreateUserResponse{
		User: &gen.User{
			Id:        id,
			FirstName: req.GetFirstName(),
			LastName:  req.GetLastName(),
		},
	}, nil
}

func (s *AuthServer) Login(ctx context.Context, req *gen.LoginRequest) (*gen.LoginResponse, error) {

	token, err := pkg.Login(s.Db, req.GetUsername(), req.GetPassword())
	if err != nil {
		log.Println("Error creating token", err)
	}

	return &gen.LoginResponse{
		Token: token,
	}, nil
}

func newAuthServer(authSrv *AuthServer) *grpc.Server {

	srv := grpc.NewServer()
	gen.RegisterAuthServiceServer(srv, authSrv)
	reflection.Register(srv)

	return srv
}

func (s *AuthServer) connectDatabase() error {
	db, err := pkg.GetConnection("postgres", "postgres", "127.0.0.1", 5434, "auth")
	if err != nil {
		log.Println("Error connecting to database", err)

		return err
	}

	s.Db = db

	return nil
}

func main() {

	listener, err := net.Listen("tcp", ":4041")
	if err != nil {
		panic(err)
	}

	a := &AuthServer{}
	a.connectDatabase()
	err = a.Db.Ping()
	if err != nil {
		log.Println("ERROR", err)
	} else {
		log.Println("SUCCESS")
	}

	authService := newAuthServer(a)
	fmt.Println("Listening on port 4040")
	if e := authService.Serve(listener); e != nil {
		panic(e)
	}
}
