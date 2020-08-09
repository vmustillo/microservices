package main

import (
	"fmt"
	"log"
	"net"

	"github.com/vmustillo/microservices/card/internal/server"
)

func main() {

	listener, err := net.Listen("tcp", ":4042")
	if err != nil {
		panic(err)
	}

	a := &server.CardServer{}
	a.ConnectDatabase()
	a.ConnectKafkaCards()

	go a.ConsumeCards()

	err = a.Db.Ping()
	if err != nil {
		log.Println("ERROR", err)
	} else {
		log.Println("SUCCESS")
	}

	cardService := server.NewCardServer(a)
	fmt.Println("Listening on port 4042")
	if e := cardService.Serve(listener); e != nil {
		panic(e)
	}
}
