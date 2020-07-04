package main

import (
	"easytable/persistence"
	"net/http"
)

func main() {
	persistence.Connect()
	http.HandleFunc("/reservations", reservations)
	http.ListenAndServe(":8080", nil)
}

func reservations(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Reservations..."))
}
