package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type helloWorldResponse struct {
	Message string `json:"message"`
	// Do not output the following field:
	Author string `json:"-"`
	// Do not output the following field if the value is empty:
	Date string `json:",omitempty"`
	// convert output to a string and rename "id"
	ID int `json:"id,string"`
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	response := helloWorldResponse{Message: "HelloWorld", Author: "Leo", Date: "789sdf", ID: 23}
	encoder := json.NewEncoder(w)
	encoder.Encode(&response)
	// data, err := json.Marshal(response)
	// if err != nil {
	// 	panic("oops!")
	// }
	// fmt.Fprint(w, string(data))
}

func main() {
	port := 8080
	http.HandleFunc("/hello", helloWorldHandler)

	// Using nil since we are using a default router ...
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
