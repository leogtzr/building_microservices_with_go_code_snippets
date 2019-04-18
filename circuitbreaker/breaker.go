package main

import (
	"fmt"
	"time"

	"github.com/eapache/go-resiliency/breaker"
)

func main() {
	/*
	   - The first errorThreshold, is the number of times a request can fail before the circuit opens
	   - The successThreshold, is the number of times that we need a successful request in the half-open state before we move back to open
	   - The timeout, is the time that the circuit will stay in the open state before changing to half-open
	*/
	b := breaker.New(3, 1, 5*time.Second)
	for {
		result := b.Run(func() error {
			// Call some service
			// fmt.Println("Calling a service ... ", time.Now())
			//time.Sleep(1 * time.Second)
			time.Sleep(500 * time.Millisecond)
			return fmt.Errorf("Timeout: %q", time.Now())
		})

		switch result {
		case nil:
			// success!
			fmt.Println("Success ... ", time.Now())
		case breaker.ErrBreakerOpen:
			// our function wasn't run because the breaker was open
			fmt.Println("Breaker open", time.Now())
		default:
			// fmt.Println("Result: ")
			fmt.Println("This:", result, time.Now())
		}

		time.Sleep(500 * time.Millisecond)
	}

	/*
	If you run this code you should see the following output. 
	After three failed requests the breaker enters the open state, 
	then after our five-second interval, we enter the half-open state, and we are allowed to make another request. 
	Unfortunately, this fails, and we again enter the fully open state, and we no longer even attempt to make the call.
	*/

}
