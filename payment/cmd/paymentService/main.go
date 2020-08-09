package main

import (
	"fmt"
	"log"
	"net"

	"github.com/vmustillo/microservices/payment/internal/server"
)

func main() {

	listener, err := net.Listen("tcp", ":4043")
	if err != nil {
		panic(err)
	}

	a := &server.PaymentServer{}
	a.ConnectDatabase()

	err = a.Db.Ping()
	if err != nil {
		log.Println("ERROR", err)
	} else {
		log.Println("SUCCESS")
	}

	paymentService := server.NewPaymentServer(a)
	fmt.Println("Listening on port 4043")
	if e := paymentService.Serve(listener); e != nil {
		panic(e)
	}
}
