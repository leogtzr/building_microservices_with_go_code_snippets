package main

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