package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"grpc-helloworld/chat"
)

func main() {

	// Tries to connect with port 9000 with insecure method (HTTP and not HTTPS) and handles that connection appropriately
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	// Calls the NewChatServiceClient method from the chat.pb.go file to initiate a client instance and passes in our connection
	c := chat.NewChatServiceClient(conn)

	//Calls the RPC Method SayHello and sends an empty background and a message with a body and handles the error, also prints the server response to the log
	response, err := c.SayHello(context.Background(), &chat.Message{Body: "Hello World, Is anyone there? (I am the client btw)"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)

}
