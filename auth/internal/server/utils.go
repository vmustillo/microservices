package server

import (
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/vmustillo/microservices/auth/pkg"
)

func (s *AuthServer) ConnectDatabase() error {
	db, err := pkg.GetConnection("postgres", "postgres", "127.0.0.1", 5434, "auth")
	if err != nil {
		log.Println("Error connecting to database", err)

		return err
	}

	s.Db = db

	return nil
}

func (s *AuthServer) ConnectKafkaAccounts() error {
	topic := "accounts"

	conn := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   topic,
		Async:   true,
	})

	s.AccountsProducer = conn

	return nil

}

func (s *AuthServer) ConnectKafkaAuth() error {
	topic := "auth"

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   topic,
		Async:   true,
	})

	s.AuthProducer = writer

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     topic,
		Partition: 0,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})

	s.AuthConsumer = reader

	return nil

}
