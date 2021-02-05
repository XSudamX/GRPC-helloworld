package main

//this is basically the register.go file in fullscale projects
// Used go mod to manage dependencies and packages  

import (
	"fmt"
	"log"
	"net"

	"grpc-helloworld/chat"

	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Basic GRPC Server")

	// Start Listening on Port 9000
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		fmt.Println("Now listening on port 9000....")
	}

	// Start a new grpc server called grpcServer
	grpcServer := grpc.NewServer()

	// Calls the server struct from chat.go to access its methods later on
	s := chat.Server{}

	// Calls the RegisterChatServiceServer method from chat.pb.go and passes in our new grpc server and the file with our RPC functions
	chat.RegisterChatServiceServer(grpcServer, &s)

	// GRPC Serve Error Handling passes in our listening port
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
