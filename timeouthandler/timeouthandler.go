package main

import (
	"net/http"
	"time"
)

func main() {
	h := http.FileServer(http.Dir("/tmp"))
	dt := 2 * time.Second
	http.ListenAndServe(":8080", http.TimeoutHandler(h, dt, "file scanning timedout!"))
}
