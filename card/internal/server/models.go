package server

import (
	"database/sql"

	kafka "github.com/segmentio/kafka-go"
)

type CardServer struct {
	Db           *sql.DB
	CardConsumer *kafka.Reader
	CardProducer *kafka.Writer
}
