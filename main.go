package main

import (
	"encoding/json"
	"net/http"
)

var apiBase = "/"

func main() {

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		data := generateUsers(100)

		json.NewEncoder(w).Encode(data)
	})

	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
