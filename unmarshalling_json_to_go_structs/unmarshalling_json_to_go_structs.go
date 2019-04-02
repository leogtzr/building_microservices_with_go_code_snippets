package main

import (
	"encoding/json"
	"net/http"
)

type helloWorldRequest struct {
	Name string `json:"name"`
}

type helloWorldResponse struct {
	Message string `json:"message"`
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	var request helloWorldRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	response := helloWorldResponse{Message: "Hello " + request.Name}
	encoder := json.NewEncoder(w)
	encoder.Encode(response)

}

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("images"))
	mux.Handle("/images/", http.StripPrefix("/images/", fs))
	mux.HandleFunc("/hello", helloWorldHandler)

	// Temporary redirect
	mux.Handle("/r", http.RedirectHandler("http://example.org", 307))

	http.ListenAndServe(":8080", mux)
}
