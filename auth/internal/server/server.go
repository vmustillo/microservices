package server

import (
	"bytes"
	"context"
	"encoding/gob"
	"log"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"github.com/vmustillo/microservices/auth/gen"
	"github.com/vmustillo/microservices/auth/internal"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func (s *AuthServer) GetUsers(ctx context.Context, req *gen.GetUsersRequest) (*gen.Users, error) {
	log.Println(req)
	log.Println(ctx)
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

	guid := uuid.New()
	id := guid.String()

	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	enc.Encode(req)

	s.AuthProducer.WriteMessages(context.Background(),
		kafka.Message{
			Headers: []kafka.Header{{Key: "operation", Value: []byte("create")}},
			Key:     []byte(id),
			Value:   b.Bytes(),
		},
	)

	s.AccountsProducer.WriteMessages(context.Background(),
		kafka.Message{
			Headers: []kafka.Header{{Key: "operation", Value: []byte("create")}},
			Key:     []byte(id),
		},
	)

	return &gen.CreateUserResponse{
		User: &gen.User{
			Id:        id,
			FirstName: req.GetFirstName(),
			LastName:  req.GetLastName(),
		},
	}, nil
}

func (s *AuthServer) Login(ctx context.Context, req *gen.LoginRequest) (*gen.LoginResponse, error) {

	token, err := internal.Login(s.Db, req.GetUsername(), req.GetPassword())
	if err != nil {
		log.Println("Error creating token", err)
	}

	return &gen.LoginResponse{
		Token: token,
	}, nil
}

func NewAuthServer(authSrv *AuthServer) *grpc.Server {

	srv := grpc.NewServer()
	gen.RegisterAuthServiceServer(srv, authSrv)
	reflection.Register(srv)

	return srv
}
