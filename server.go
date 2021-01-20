package main

import (
	"fmt"
	"log"
	"net"

	"grpc-helloworld/chat"

	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Basic GRPC Server")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		fmt.Println("Now listening on port 9000")
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
