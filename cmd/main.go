package main

import (
	"easytable/api"
	"easytable/app"
	"easytable/persistence"
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	databaseURL := os.Getenv("DATABASE_URL")
	shouldRunDBMigration := getMigrationFlagFromCommandLine()
	db, err := persistence.Connect(databaseURL, shouldRunDBMigration)

	if err != nil {
		log.Panic("Failed to connect to the database", err.Error())
	}

	app := app.NewApp(app.AppInput{ReservationsRepo: persistence.NewReservationsRepo(db)})

	api := api.NewAPI(api.APIInput{App: app})

	api.Init()

	http.ListenAndServe(":3000", api.Router())
}

func getMigrationFlagFromCommandLine() bool {
	migration := flag.Bool("migration", false, "indicates if we should run a db migration on startup")
	flag.Parse()
	return *migration
}
