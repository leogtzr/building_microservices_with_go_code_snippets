package main

import (
	"log"
	"net/http"

	"./data"
	"./handlers"
)

func main() {

	store := data.MemoryStore{}
	handler := handlers.Search{DataStore: store}
	err := http.ListenAndServe(":8080", &handler)
	if err != nil {
		log.Fatal(err)
	}
}
