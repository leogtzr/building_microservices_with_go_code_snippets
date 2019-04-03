package server

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

// HelloWorldResponse ...
type HelloWorldRequest struct {
	Name string
}

// HelloWorldResponse ...
type HelloWorldResponse struct {
	Message string
}

const port = 1234

func main() {
	log.Printf("Server starting on port %v\n", port)
	StartServer()
}

func StartServer() {
	helloWorld := &HelloWorldHandler{}
	rpc.Register(helloWorld)

	l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to listen on given port: %s", err))
	}
	defer l.Close()

	for {
		conn, _ := l.Accept()
		go rpc.ServeConn(conn)
	}
}

type HelloWorldHandler struct{}

func (h *HelloWorldHandler) HelloWorld(args *HelloWorldRequest, reply *HelloWorldResponse) error {
	reply.Message = "Hello " + args.Name
	return nil
}
