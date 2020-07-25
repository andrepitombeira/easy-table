package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func (api *API) initReservations() {
	api.router.Get("/reservations", api.listReservations)
	api.router.Get("/reservations/{id}", api.getReservation)
}

func (api *API) listReservations(w http.ResponseWriter, r *http.Request) {
	reservations, err := api.app.ListReservations()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error"))
	}

	output, err := json.Marshal(reservations)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	w.Write(output)
}

func (api *API) getReservation(w http.ResponseWriter, r *http.Request) {
	id, err  := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error"))
	}

	reservation, err := api.app.GetReservationByID(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error"))
	}

	output, err := json.Marshal(reservation)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	w.Write(output)
}
