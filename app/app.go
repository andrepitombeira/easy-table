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

type ReservationsRepo interface {
	ListReservations() ([]Reservation, error)
}

type App struct {
	reservationsRepo ReservationsRepo
}

type AppInput struct {
	ReservationsRepo ReservationsRepo
}

func NewApp(input AppInput) *App {
	return &App{
		reservationsRepo: input.ReservationsRepo,
	}
}

func (app *App) ListReservations() ([]Reservation, error) {
	return app.reservationsRepo.ListReservations()
}
