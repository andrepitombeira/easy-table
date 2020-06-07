package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Reservation struct {
	ID             int
	Start          string
	End            string
	GuestID        int
	NumberOfGuests int
	RestaurantID   int
}

type Guest struct {
	ID        int
	FirstName int
	LastName  int
	Email     int
	Phone     string
}

type Restaurant struct {
	ID      int
	Name    int
	Phone   string
	Address RestaurantAddress
}

type RestaurantAddress struct {
	RestaurantID int
	Street       string
	ZipCode      string
	City         string
	Country      string
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
