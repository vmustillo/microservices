package server

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"log"

	"github.com/vmustillo/microservices/card/gen"
	"github.com/vmustillo/microservices/card/internal"
)

func (s *CardServer) ConsumeCards() {

	defer s.CardConsumer.Close()

	for {
		m, err := s.CardConsumer.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("Error reading message from Kafka:", err)
		}

		var operation string

		for _, v := range m.Headers {
			if v.Key == "operation" {
				operation = string(v.Value)
				break
			}
		}

		fmt.Println(operation, string(m.Key))
		var req *gen.CreateCardRequest
		r := bytes.NewReader(m.Value)
		dec := gob.NewDecoder(r)
		dec.Decode(&req)
		log.Println("Account ID:", req.GetAccountId())

		switch operation {
		case "create":
			log.Println("Creating Card and associating to account:", string(m.Key))
			internal.CreateCard(s.Db, string(m.Key), req.GetAccountId())
		}
	}
}
