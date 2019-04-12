package search

import (
	"encoding/json"
	"net/http"
)

// User ...
type User struct {
	name string `json:"user"`
}

func someUsers() []User {
	return []User{}
}

func getUsers(w http.ResponseWriter, r *http.Request) {

	users, err := json.Marshal(someUsers())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(users)
}

func Search() {

}
