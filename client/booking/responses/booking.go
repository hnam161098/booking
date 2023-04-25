package responses

type TicketResponse struct {
	ID         string `json:"id"`
	GuessName  string `json:"guess_name"`
	GuessId    string `json:"guess_id"`
	IdPersonal string `json:"id_personal"`
	From       string `json:"from"`
	To         string `json:"to"`
	Date       string `json:"date"`
	PlaneID    string `json:"plane_id"`
	AirportID  string `json:"airport_id"`
	SeatName   string `json:"seat_name"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	SeatID     int64  `json:"seat_id"`
	Status     int64  `json:"status"`
}

type CustomerResponse struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	IdPersonal string   `json:"id_personal"`
	Age        string   `json:"age"`
	Address    string   `json:"address"`
	Tags       []string `json:"tags"`
	CreatedAt  string   `json:"created_at"`
	UpdatedAt  string   `json:"updated_at"`
}

type TicketInformationResponse struct {
	Ticket   TicketResponse   `json:"ticket_information"`
	Customer CustomerResponse `json:"customer_information"`
}
