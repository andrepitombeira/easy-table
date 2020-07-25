package api_test

import (
	"easytable/api"
	"easytable/app"
	"easytable/mocks"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestListReservations(t *testing.T) {
	reservationsRepo := new(mocks.ReservationsRepo)
	appMock := app.NewApp(app.AppInput{ ReservationsRepo: reservationsRepo}) 
	api := api.NewAPI(api.APIInput{ App: appMock})
	
	reservations := []app.Reservation{
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
	reservationsRepo.On("ListReservations").Return(reservations, nil)

	req, err := http.NewRequest("GET", "/reservations", nil)

	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	api.Router().ServeHTTP(res, req)

	output := []app.Reservation{}

	if err = json.NewDecoder(res.Body).Decode(&output); err != nil {
		t.Fatal(err)
	}

	if status := res.Code; status != http.StatusOK {
		t.Errorf("Expected status to be: %v, got: %v", http.StatusOK, status)
	}

	if !reflect.DeepEqual(output, reservations) {
		t.Errorf("Expected reservations to be: %v, got: %v", reservations, output)
	}
}

func TestGetReservation(t *testing.T) {
	reservationsRepo := new(mocks.ReservationsRepo)
	appMock := app.NewApp(app.AppInput{ ReservationsRepo: reservationsRepo}) 
	api := api.NewAPI(api.APIInput{ App: appMock})
	
	reservation := app.Reservation{
		ID:             1,
		Start:          "2020-07-03 08:00:00",
		End:            "2020-07-03 08:30:00 UTC",
		NumberOfGuests: 11,
		Guest:          "Giffer Tincey",
	}

	reservationsRepo.On("GetReservationByID").Return(reservation, nil)

	req, err := http.NewRequest("GET", "/reservations/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	api.Router().ServeHTTP(res, req)

	output := app.Reservation{}

	if err = json.NewDecoder(res.Body).Decode(&output); err != nil {
		t.Fatal(err)
	}

	if status := res.Code; status != http.StatusOK {
		t.Errorf("Expected status to be: %v, got: %v", http.StatusOK, status)
	}

	if !reflect.DeepEqual(output, reservation) {
		t.Errorf("Expected reservations to be: %v, got: %v", reservation, output)
	}
}