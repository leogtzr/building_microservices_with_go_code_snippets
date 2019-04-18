package main

import (
	"fmt"
	"time"

	"github.com/eapache/go-resiliency/deadline"
)

func slowFunction() {
	time.Sleep(time.Millisecond * 3)
}

func makeTimeoutRequest() {
	dl := deadline.New(1 * time.Second)
	err := dl.Run(func(stopper <-chan struct{}) error {
		slowFunction()
		return nil
	})
	switch err {
	case deadline.ErrTimedOut:
		fmt.Println("Timeout")
	default:
		fmt.Println(err)
	}
}

func main() {
	makeTimeoutRequest()
}
