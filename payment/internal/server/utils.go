package server

import (
	"log"

	"github.com/vmustillo/microservices/payment/pkg"
)

func (s *PaymentServer) ConnectDatabase() error {
	db, err := pkg.GetConnection("postgres", "postgres", "127.0.0.1", 5436, "payments")
	if err != nil {
		log.Println("Error connecting to database", err)

		return err
	}

	s.Db = db

	return nil
}
