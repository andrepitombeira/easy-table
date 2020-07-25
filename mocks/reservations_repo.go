package mocks

import (
	"easytable/app"

	"github.com/stretchr/testify/mock"
)

type ReservationsRepo struct {
	mock.Mock
}

func (repo *ReservationsRepo) ListReservations() ([]app.Reservation, error) {
	args := repo.Called()
	return args.Get(0).([]app.Reservation), args.Error(1)
}

func (repo *ReservationsRepo) GetReservationByID(id int) (app.Reservation, error) {
	args := repo.Called()
	return args.Get(0).(app.Reservation), args.Error(1)
}
