// curl -I -X OPTIONS 'http://localhost:8080/hello'
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

	if r.Method == "OPTIONS" {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "GET")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	response := helloWorldResponse{Message: "Hello World"}
	data, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	fmt.Fprint(w, string(data))
}

func main() {
	port := 8080
	http.HandleFunc("/hello", helloWorldHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
