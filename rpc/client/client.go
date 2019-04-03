package client

import (
	"fmt"
	"log"
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

func CreateClient() *rpc.Client {
	client, err := rpc.Dial("tcp", fmt.Sprintf("localhost:%v", port))
	if err != nil {
		log.Fatal("dialing:", err)
	}

	return client
}

func PerformRequest(client *rpc.Client) HelloWorldResponse {
	args := &HelloWorldRequest{Name: "World"}
	var reply HelloWorldResponse

	err := client.Call("HelloWorldHandler.HelloWorld", args, &reply)
	if err != nil {
		log.Fatal("error:", err)
	}

	return reply
}
