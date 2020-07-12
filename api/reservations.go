package api

import (
	"encoding/json"
	"net/http"
)

func (api *API) initReservations() {
	api.router.Get("/reservations", api.listReservations)
}

func (api *API) listReservations(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

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
