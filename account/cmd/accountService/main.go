package main

import (
	"fmt"
	"log"
	"net"

	"github.com/vmustillo/microservices/account/internal/server"
)

func main() {

	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	a := &server.AccountServer{}
	a.ConnectDatabase()
	a.ConnectKafkaAccounts()

	go a.ConsumeAccounts()

	err = a.Db.Ping()
	if err != nil {
		log.Println("ERROR", err)
	} else {
		log.Println("SUCCESS")
	}

	accountService := server.NewAccountServer(a)
	fmt.Println("Listening on port 4040")
	if e := accountService.Serve(listener); e != nil {
		panic(e)
	}
}
