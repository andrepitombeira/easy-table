package persistence_test

import (
	"database/sql"
	"easytable/app"
	"easytable/persistence"
	"reflect"
	"testing"
)

func TestListReservations(t *testing.T) {
	db, err := persistence.SetupTestDB()

	if err != nil {
		t.Fatal("Cannot connect to DB!")
	}

	reservationsMockData := []app.Reservation{
		{
			ID:             1,
			Start:          "2020-07-03 08:00:00",
			End:            "2020-07-03 08:30:00 UTC",
			NumberOfGuests: 11,
			Guest:          "Giffer Tincey",
		},
		{
			ID:             2,
			Start:          "2020-07-03 04:00:00",
			End:            "2020-07-03 04:30:00 UTC",
			NumberOfGuests: 8,
			Guest:          "Garner Mc Cahey",
		},
		{
			ID:             3,
			Start:          "2020-07-03 05:00:00",
			End:            "2020-07-03 05:30:00 UTC",
			NumberOfGuests: 9,
			Guest:          "Dawna MacNelly",
		},
	}

	repo := persistence.NewReservationsRepo(db)

	err = insertReservationsIntoDB(db, reservationsMockData)

	reservations, err := repo.ListReservations()

	if err != nil {
		t.Errorf("Expected error to be %v, got %v", nil, err)
	}

	if !reflect.DeepEqual(reservations, reservationsMockData) {
		t.Errorf("Expected reservations to be %v, got %v", reservationsMockData, reservations)
	}
}

func insertReservationsIntoDB(db *sql.DB, reservations []app.Reservation) error {
	query := `
		INSERT into reservations (start, end, number_of_guests, guest)
		VALUES ($1, $2, $3, $4)
	`

	for _, reservation := range reservations {
		_, err := db.Exec(query,
			reservation.Start,
			reservation.End,
			reservation.NumberOfGuests,
			reservation.Guest,
		)

		if err != nil {
			return err
		}
	}

	return nil
}
