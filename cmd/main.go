package main

import (
	"easytable/api"
	"easytable/app"
	"easytable/database"
	"log"
	"net/http"
)

func main() {
	db, err := database.Connect()

	if err != nil {
		log.Panic("Failed to connect to the database", err.Error())
	}

	app := app.NewApp(app.AppInput{ReservationsRepo: database.NewReservationsRepo(db)})

	api := api.NewAPI(api.APIInput{App: app})

	api.Init()

	http.ListenAndServe(":3000", api.Router())
}
