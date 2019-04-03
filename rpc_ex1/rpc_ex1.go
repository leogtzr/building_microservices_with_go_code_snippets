package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

const port = 1234

// HelloWorldHandler ...
type HelloWorldHandler struct{}

// HelloWorldRequest ...
type HelloWorldRequest struct {
	Name string
}

// HelloWorldResponse ...
type HelloWorldResponse struct {
	Message string
}

func (h *HelloWorldHandler) HelloWorld(args *HelloWorldRequest, reply *HelloWorldResponse) error {
	reply.Message = "Hello " + args.Name
	return nil
}

func main() {

}

func startServer() {
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
