package main

import (
	"easytable/api"
	"easytable/app"
	"easytable/persistence"
	"log"
	"net/http"
)

func main() {
	db, err := persistence.Connect()

	if err != nil {
		log.Panic("Failed to connect to the database", err.Error())
	}

	app := app.NewApp(app.AppInput{ReservationsRepo: persistence.NewReservationsRepo(db)})

	api := api.NewAPI(api.APIInput{App: app})

	api.Init()

	http.ListenAndServe(":3000", api.Router())
}
