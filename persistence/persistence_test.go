package persistence

import (
	"database/sql"
	"os"
)

func SetupTestDB() (*sql.DB, error) {
	databaseURL := os.Getenv("TEST_DATABASE_URL")

	db, err := Connect(databaseURL, false)

	if err != nil {
		panic("Failed to connect to database")
	}

	createReservationsTable(db)
	resetDB(db)

	return db, nil
}

func resetDB(db *sql.DB) {
	db.Exec("DELETE FROM reservations")
}
