package persistence

import (
	"database/sql"
	"easytable/app"
)

type ReservationsRepo struct {
	db *sql.DB
}

func NewReservationsRepo(db *sql.DB) *ReservationsRepo {
	return &ReservationsRepo{
		db: db,
	}
}

func (repo *ReservationsRepo) ListReservations() ([]app.Reservation, error) {
	rows, err := repo.db.Query("SELECT * from reservations")

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	reservations := []app.Reservation{}

	for rows.Next() {
		reservation := app.Reservation{}

		err = rows.Scan(&reservation.ID, &reservation.Start, &reservation.End, &reservation.Guest, &reservation.NumberOfGuests)

		if err != nil {
			panic(err)
		}

		reservations = append(reservations, reservation)
	}

	return reservations, nil
}
