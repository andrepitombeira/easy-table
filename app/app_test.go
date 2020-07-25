package app_test

import (
	"easytable/app"
	"easytable/mocks"
	"errors"
	"reflect"
	"testing"
)

func TestListReservations(t *testing.T) {
	reservations := []app.Reservation{
		{
			ID:             1,
			Start:          "2020-07-03 08:00:00",
			End:            "2020-07-03 08:30:00 UTC",
			Guest:          "John Doe",
			NumberOfGuests: 2,
		},
	}
	reservationsRepo := new(mocks.ReservationsRepo)
	reservationsRepo.On("ListReservations").Return(reservations, nil)

	app := app.NewApp(app.AppInput{ReservationsRepo: reservationsRepo})

	reservationsOutput, err := app.ListReservations()

	if !reflect.DeepEqual(reservationsOutput, reservations) {
		t.Errorf("Expected: %v, got: %v\n", reservations, reservationsOutput)
	}

	if err != nil {
		t.Errorf("Expected error to be: %v, got: %v", nil, err)
	}
}

func TestListReservationsWithError(t *testing.T) {
	reservations := []app.Reservation{}
	dbError := errors.New("The DB could not fetch the reservations!")
	reservationsRepo := new(mocks.ReservationsRepo)
	reservationsRepo.On("ListReservations").Return(reservations, dbError)

	app := app.NewApp(app.AppInput{ReservationsRepo: reservationsRepo})

	reservationsOutput, err := app.ListReservations()

	if !reflect.DeepEqual(reservationsOutput, reservations) {
		t.Errorf("Expected Reservations to be: %v, got: %v\n", reservations, reservationsOutput)
	}

	if err != dbError {
		t.Errorf("Expected Error to be: %v, got: %v\n", dbError, err)
	}
}

func TestGetReservationByID(t *testing.T) {
	reservation := app.Reservation{
			ID:             1,
			Start:          "2020-07-03 08:00:00",
			End:            "2020-07-03 08:30:00 UTC",
			Guest:          "John Doe",
			NumberOfGuests: 2,
	}
	reservationsRepo := new(mocks.ReservationsRepo)
	reservationsRepo.On("GetReservationByID").Return(reservation, nil)

	app := app.NewApp(app.AppInput{ReservationsRepo: reservationsRepo})

	reservationOutput, err := app.GetReservationByID(1)

	if !reflect.DeepEqual(reservationOutput, reservation) {
		t.Errorf("Expected reservation to be: %v, got: %v\n", reservation, reservationOutput)
	}

	if err != nil {
		t.Errorf("Expected error to be: %v, got: %v", nil, err)
	}
}