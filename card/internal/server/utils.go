package server

import (
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/vmustillo/microservices/card/pkg"
)

func (s *CardServer) ConnectDatabase() error {
	db, err := pkg.GetConnection("postgres", "postgres", "127.0.0.1", 5435, "cards")
	if err != nil {
		log.Println("Error connecting to database", err)

		return err
	}

	s.Db = db

	return nil
}

func (s *CardServer) ConnectKafkaCards() error {
	topic := "cards"

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   topic,
		Async:   true,
	})

	s.CardProducer = writer

	conn := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{"localhost:9092"},
		Topic:       topic,
		Partition:   0,
		MinBytes:    10e3, // 10KB
		MaxBytes:    10e6, // 10MB
		StartOffset: 6,
	})

	s.CardConsumer = conn

	return nil
}
