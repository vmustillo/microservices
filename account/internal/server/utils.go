package server

import (
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/vmustillo/microservices/account/pkg"
)

func (s *AccountServer) ConnectDatabase() error {
	db, err := pkg.GetConnection("postgres", "postgres", "127.0.0.1", 5433, "account")
	if err != nil {
		log.Println("Error connecting to database", err)

		return err
	}

	s.Db = db

	return nil
}

func (s *AccountServer) ConnectKafkaAccounts() error {
	topic := "accounts"
	// partition := 0

	conn := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{"localhost:9092"},
		Topic:       topic,
		Partition:   0,
		MinBytes:    10e3, // 10KB
		MaxBytes:    10e6, // 10MB
		StartOffset: 6,
	})
	// conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	// if err != nil {
	// 	log.Println("Cannot connect to Kakfa")

	// 	return err
	// }

	// conn.SetReadDeadline(time.Now().Add(10 * time.Second))

	s.AccountsConsumer = conn

	return nil
}
