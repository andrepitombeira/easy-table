package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Reservation struct {
	ID             int    `json:"id"`
	GuestID        int    `json:"guestId"`
	RestaurantID   int    `json:"restaurantId"`
	NumberOfGuests int    `json:"numberOfGuests"`
	Start          string `json:"start"`
	End            string `json:"end"`
}

type Reservations struct {
	Reservations []Reservation `json:"reservations"`
}

func main() {
	http.HandleFunc("/reservations", reservations)
	http.ListenAndServe(":8080", nil)
}

func reservations(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	reservationsFile, err := os.Open("reservations.json")

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println("Successfully opened reservations.json")

	defer reservationsFile.Close()

	reservationsByteValue, _ := ioutil.ReadAll(reservationsFile)

	w.WriteHeader(http.StatusOK)
	w.Write(reservationsByteValue)
}
