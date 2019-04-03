package main

import (
	"fmt"

	"github.com/leogtzr/chapter1/rpc/client"
	"github.com/leogtzr/chapter1/rpc/server"
)

func main() {
	go server.StartServer()

	c := client.CreateClient()
	defer c.Close()

	reply := client.PerformRequest(c)
	fmt.Println(reply.Message)
}
