package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	account "github.com/vmustillo/microservices/account/gen"
	auth "github.com/vmustillo/microservices/auth/gen"
)

var (
	echoEndpoint = flag.String("echo_endpoint", "localhost:4040", "endpoint of account")
)

func Run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := account.RegisterAccountServiceHandlerFromEndpoint(ctx, mux, *echoEndpoint, opts)
	if err != nil {
		return err
	}

	authEndpoint := "localhost:4041"
	err = auth.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, *&authEndpoint, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":8080", mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	fmt.Println("Listening on port 8080...")
	if err := Run(); err != nil {
		glog.Fatal(err)
	}
}
