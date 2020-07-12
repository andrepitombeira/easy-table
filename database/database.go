package database

import (
	"bufio"
	"database/sql"
	"easytable/app"
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func Connect(databaseURL string, shouldRunMigration bool) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", databaseURL)

	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}

	// TODO: move this logic to db migration
	if shouldRunMigration {
		createReservationsTable(db)
		populateDBWithMockData(db)
	}

	return db, nil
}

func createReservationsTable(db *sql.DB) {
	fmt.Printf("Creating DB tables...\n")

	query := `
		CREATE TABLE IF NOT EXISTS reservations (
			id INTEGER PRIMARY KEY, 
			start TEXT, 
			end TEXT, 
			guest TEXT,
			number_of_guests INTEGER
		)
	`

	statement, _ := db.Prepare(query)
	statement.Exec()

	fmt.Printf("Finished create DB tables.\n")
}

func populateDBWithMockData(db *sql.DB) {
	fmt.Printf("Inserting data...\n")
	query := `
		INSERT INTO reservations (start, end, guest, number_of_guests) VALUES (?, ?, ?, ?)
	`
	statement, _ := db.Prepare(query)

	reservationsFile, err := os.Open("./data/reservations.json")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer reservationsFile.Close()

	reservationsByteValue := bufio.NewReader(reservationsFile)

	output := app.ReservationsOutput{}

	if err := json.NewDecoder(reservationsByteValue).Decode(&output); err != nil {
		fmt.Println(err)
	}

	for _, reservation := range output.Reservations {
		fmt.Printf("Reservation: %v\n", reservation)
		statement.Exec(reservation.Start, reservation.End, reservation.Guest, reservation.NumberOfGuests)
	}

	fmt.Printf("Finished insert data.\n")
}
