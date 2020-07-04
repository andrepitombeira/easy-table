package app

type Reservation struct {
	ID             int    `json:"id"`
	Start          string `json:"start"`
	End            string `json:"end"`
	Guest          string `json:"guest"`
	NumberOfGuests int    `json:"number_of_guests"`
}

type ReservationsOutput struct {
	Reservations []Reservation `json:"reservations"`
}
