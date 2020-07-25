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

func (repo *ReservationsRepo) GetReservationByID(id int) (app.Reservation, error) {
	var reservation app.Reservation
	row := repo.db.QueryRow("SELECT * from reservations where id = $1", id)

	switch err := row.Scan(&reservation.ID, &reservation.Start, &reservation.End, &reservation.Guest, &reservation.NumberOfGuests); err {
		case nil:
			return reservation, nil
		case sql.ErrNoRows:
			return app.Reservation{}, nil
		default:
			panic(err)
	}
}