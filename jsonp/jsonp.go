package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type helloWorldResponse struct {
	Message string `json:"message"`
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	response := helloWorldResponse{Message: "Hello World"}
	data, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	callback := r.URL.Query().Get("callback")
	if callback != "" {
		r.Header.Add("Content-Type", "application/javascript")
		fmt.Fprintf(w, "%s(%s)", callback, string(data))
	} else {
		fmt.Fprint(w, string(data))
	}
}

func main() {
	port := 8080
	http.HandleFunc("/hello", helloWorldHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
