package main

import (
	"net/http"

	"./search"
	"github.com/gorilla/mux"
)

func setUserRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/user", search.GetUsers).Methods("GET")
	return r
}

func main() {
	http.ListenAndServe(":8080", setUserRoutes())
}
