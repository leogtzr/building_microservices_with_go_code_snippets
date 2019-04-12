package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func setUserRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/user", getUsers).Methods("GET")
	return r
}

func main() {
	http.ListenAndServe(":8080", setUserRoutes())
}
