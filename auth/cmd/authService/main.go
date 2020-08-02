package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	gw "github.com/vmustillo/microservices/auth/gen"
	pkg "github.com/vmustillo/microservices/auth/pkg"
)

type AuthServer struct {
	Db *sql.DB
}

func (*AuthServer) GetUsers(ctx context.Context, req *gw.GetUsersRequest) (*gw.Users, error) {
	return &gw.Users{
		Users: []*gw.User{
			&gw.User{
				Id:        "guid1",
				AccountID: "accountID1",
				FirstName: "Vin",
				LastName:  "M",
				FullName:  "Vin M",
			},
		},
	}, nil
}

func (*AuthServer) PostUser(ctx context.Context, req *gw.PostUserRequest) (*gw.PostUserResponse, error) {
	return &gw.PostUserResponse{
		User: &gw.User{},
	}, nil
}

func newAuthServer(authSrv *AuthServer) *grpc.Server {

	srv := grpc.NewServer()
	gw.RegisterAuthServiceServer(srv, authSrv)
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
