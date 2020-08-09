package server

import (
	"database/sql"

	kafka "github.com/segmentio/kafka-go"
)

type PaymentServer struct {
	Db               *sql.DB
	AccountsConsumer *kafka.Reader
	AccountsProducer *kafka.Writer
}
