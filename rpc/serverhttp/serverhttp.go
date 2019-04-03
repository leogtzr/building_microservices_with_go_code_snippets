package server

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

// HelloWorldRequest ...
type HelloWorldRequest struct {
	Name string
}

// HelloWorldResponse ...
type HelloWorldResponse struct {
	Message string
}

// HelloWorldHandler ...
type HelloWorldHandler struct{}

const port = 1234

func main() {
	log.Printf("Server starting on port %v\n", port)
	StartServer()
}

// StartServer ...
func StartServer() {
	helloWorld := &HelloWorldHandler{}
	rpc.Register(helloWorld)
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to listen on given port: %s", err))
	}

	log.Printf("Server starting on port %v\n", port)
	http.Serve(l, nil)
}

func (h *HelloWorldHandler) HelloWorld(args *HelloWorldRequest, reply *HelloWorldResponse) error {
	reply.Message = "Hello " + args.Name
	return nil
}
