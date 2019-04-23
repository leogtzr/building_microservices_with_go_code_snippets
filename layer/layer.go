package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type SomeHandler struct{}

func (h *SomeHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	if r.Header.Get("X-Request-ID") == "" {
		r.Header.Set("X-Request-ID", uuid.New().String())
		rw.Header().Set("X-Request-ID", uuid.New().String())
	}

	fmt.Fprintf(rw, "Hello: %d", 1)
}

func main() {

	http.Handle("/hello", &SomeHandler{})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
