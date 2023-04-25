package requests

type TicketModelRequest struct {
	ID         string `json:"id"`
	GuessName  string `json:"guess_name"`
	GuessId    string `json:"guess_id"`
	IdPersonal string `json:"id_personal" binding:"required"`
	From       string `json:"from" binding:"required"`
	To         string `json:"to" binding:"required"`
	Date       string `json:"date" binding:"required"`
	PlaneID    string `json:"plane_id" binding:"required"`
	AirportID  string `json:"airport_id" binding:"required"`
	SeatName   string `json:"seat_name" binding:"required"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	SeatID     int64  `json:"seat_id" binding:"required"`
	Status     int64  `json:"status"`
}

type FindTicketRequest struct {
	Code string `json:"code" binding:"required"`
}
