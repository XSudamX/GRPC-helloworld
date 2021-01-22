package chat

//This is the file that contains the definitions of the RPC calls.

import (
	"log"

	"golang.org/x/net/context"
)

// Server ... struct that is exported in the server.go file
type Server struct {
}

/*
SayHello ... This is the definition of the RPC Call SayHello
1. THIS IS A METHOD
2. Anything with the type "Pointer to a Server" can automatically access this method(s)
3. This is how the RPC Works
4. The Client can now call this method when/if he needs it

What this particular method does;
The Say Hello method, which is a part of the "Pointer to a server type" takes in a
context from the "context.Context" which is calling the context from the Context Package
and assigning it to the variable ctx

and it also takes in a pointer to a message

and outputs a pointer to a message and an error

in the function it logs to console that it has been called and that the message body from the client
has been recieved

and returns a message with a body with error as nil.

*/
func (s *Server) SayHello(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &Message{Body: "World here! Hello! (I am the server btw)."}, nil
}
