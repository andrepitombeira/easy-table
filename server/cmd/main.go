package main

import (
	"easytable/api"
	"easytable/persistence"
	"net/http"
)

func main() {
	persistence.Connect()

	api := api.NewAPI()

	api.Init()

	http.ListenAndServe(":3000", api.Router())
}
