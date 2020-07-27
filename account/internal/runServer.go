package internal

import (
  "flag"
  "net/http"
   
  "golang.org/x/net/context"
  "github.com/grpc-ecosystem/grpc-gateway/runtime"
  "google.golang.org/grpc"
   	
  gw "github.com/vmustillo/microservices/accountservice/gen"
)
   
var (
  echoEndpoint = flag.String("echo_endpoint", "localhost:9090", "endpoint of YourService")
)
   
func Run() error {
  ctx := context.Background()
  ctx, cancel := context.WithCancel(ctx)
  defer cancel()
   
  mux := runtime.NewServeMux()
  opts := []grpc.DialOption{grpc.WithInsecure()}
  err := gw.RegisterYourServiceHandlerFromEndpoint(ctx, mux, *echoEndpoint, opts)
  if err != nil {
    return err
  }
   
  return http.ListenAndServe(":8080", mux)
}