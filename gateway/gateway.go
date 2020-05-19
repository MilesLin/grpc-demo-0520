package main

import (
	"context"
	"fmt"
	"grpcdemo/pb"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterGreeterHandlerFromEndpoint(ctx, mux, "localhost:50051", opts)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("gRPC-gateway is running on port 8080.")
	http.ListenAndServe(":8080", mux)
}
