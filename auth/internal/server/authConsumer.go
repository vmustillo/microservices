package server

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"

	"github.com/vmustillo/microservices/auth/gen"
	internal "github.com/vmustillo/microservices/auth/internal"
)

func (s *AuthServer) ConsumeAuth() {
	defer s.AuthConsumer.Close()

	for {
		m, err := s.AuthConsumer.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("Error reading message from Kafka:", err)
		}

		var operation string

		for _, v := range m.Headers {
			fmt.Println("Headers: ", v)
			if v.Key == "operation" {
				operation = string(v.Value)
				break
			}
		}

		var req *gen.CreateUserRequest
		r := bytes.NewReader(m.Value)
		dec := gob.NewDecoder(r)
		dec.Decode(&req)

		fmt.Println("Request from kafka:", req)

		switch operation {
		case "create":
			fmt.Println("Entering create function")
			internal.CreateUser(s.Db, req.GetFirstName(), req.GetLastName(), req.GetUsername(), req.GetPassword())
		}
	}
}
