package main

import (
	"fmt"
	"log"
	"net"

	"github.com/roopu1997/greeter-grpc/greeter"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := greeter.Server{}

	grpcServer := grpc.NewServer()

	greeter.RegisterGreeterServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
