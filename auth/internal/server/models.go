package server

import (
	"database/sql"

	"github.com/segmentio/kafka-go"
)

type AuthServer struct {
	Db               *sql.DB
	AccountsProducer *kafka.Writer
	AuthProducer     *kafka.Writer
	AuthConsumer     *kafka.Reader
}
