package main

import (
	"database/sql"
	"log"
)

func Connect(databaseURL string) {
	db, err := sql.Open("postgress", databaseURL)

	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
}
