package main

import (
	"fmt"
	"log"
	"net"

	server "github.com/vmustillo/microservices/auth/internal/server"
)

func main() {

	listener, err := net.Listen("tcp", ":4041")
	if err != nil {
		panic(err)
	}

	a := &server.AuthServer{}
	a.ConnectDatabase()
	a.ConnectKafkaAccounts()
	a.ConnectKafkaAuth()

	err = a.Db.Ping()

	if err != nil {
		log.Println("ERROR", err)
	} else {
		log.Println("SUCCESS")
	}

	go a.ConsumeAuth()

	authService := server.NewAuthServer(a)

	fmt.Println("Listening on port 4041")
	if e := authService.Serve(listener); e != nil {
		panic(e)
	}
}
