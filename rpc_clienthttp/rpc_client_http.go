package main

import (
	"fmt"

	"github.com/leogtzr/building_microservices_with_go_code_snippets/rpc/client"
	"github.com/leogtzr/building_microservices_with_go_code_snippets/rpc/serverhttp"
)

func main() {
	go server.StartServer()

	c := client.CreateClient()
	defer c.Close()

	reply := client.PerformRequest(c)
	fmt.Println(reply.Message)
}
