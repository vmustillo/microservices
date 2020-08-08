package server

import (
	"context"
	"fmt"

	"github.com/vmustillo/microservices/account/internal"
)

func (s *AccountServer) ConsumeAccounts() {

	defer s.AccountsConsumer.Close()

	for {
		m, err := s.AccountsConsumer.ReadMessage(context.Background())
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

		switch operation {
		case "create":
			internal.CreateAccount(s.Db, string(m.Key))
		}
	}
}
