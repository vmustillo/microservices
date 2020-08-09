package server

import (
	"bytes"
	"context"
	"encoding/gob"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/google/uuid"
	kafka "github.com/segmentio/kafka-go"
	"github.com/vmustillo/microservices/card/gen"
)

func (*CardServer) GetCard(ctx context.Context, req *gen.GetCardRequest) (*gen.Card, error) {
	return &gen.Card{
		Id:        "guid1",
		Number:    "0123-4567-8910-1112",
		UserId:    "12345",
		AccountId: "1234",
	}, nil
}

func (s *CardServer) CreateCard(ctx context.Context, req *gen.CreateCardRequest) (*gen.Card, error) {
	guid := uuid.New()
	id := guid.String()

	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	enc.Encode(req)

	s.CardProducer.WriteMessages(context.Background(),
		kafka.Message{
			Headers: []kafka.Header{{Key: "operation", Value: []byte("create")}},
			Key:     []byte(id),
			Value:   b.Bytes(),
		},
	)

	return &gen.Card{}, nil
}

func NewCardServer(cardSrv *CardServer) *grpc.Server {

	srv := grpc.NewServer()
	gen.RegisterCardServiceServer(srv, cardSrv)
	reflection.Register(srv)

	return srv
}
