package main

import (
	"context"
	"encoding/json"
	"net/http"
)

type helloWorldRequest struct {
	Name string `json:"name"`
}

type helloWorldHandler struct{}

type validationContextKey string

func newHelloWorldHandler() http.Handler {
	return helloWorldHandler{}
}

func (h helloWorldHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// response := helloWorldResponse{Message: "Hello " + "World"}
	// encoder := json.NewEncoder(rw)
	// encoder.Encode(response)

	name := r.Context().Value(validationContextKey("name")).(string)
	response := helloWorldResponse{Message: "Hello " + name}

	encoder := json.NewEncoder(rw)
	encoder.Encode(response)
}

type helloWorldResponse struct {
	Message string `json:"message"`
}

type validationHandler struct {
	next http.Handler
}

func newValidationHandler(next http.Handler) http.Handler {
	return validationHandler{next: next}
}

func (h validationHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var request helloWorldRequest
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)
	if err != nil {
		http.Error(rw, "Bad request", http.StatusBadRequest)
		return
	}

	c := context.WithValue(r.Context(), validationContextKey("name"), request.Name)
	r = r.WithContext(c)

	h.next.ServeHTTP(rw, r)
}

// func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
// 	var request helloWorldRequest
// 	decoder := json.NewDecoder(r.Body)
// 	err := decoder.Decode(&request)
// 	if err != nil {
// 		http.Error(w, "Bad request", http.StatusBadRequest)
// 		return
// 	}

// 	response := helloWorldResponse{Message: "Hello " + request.Name}
// 	encoder := json.NewEncoder(w)
// 	encoder.Encode(response)

// }

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("images"))
	mux.Handle("/images/", http.StripPrefix("/images/", fs))
	// mux.HandleFunc("/hello", helloWorldHandler)

	// Temporary redirect
	mux.Handle("/r", http.RedirectHandler("http://example.org", 307))

	http.ListenAndServe(":8080", mux)
}
