package models

type Ticket struct {
	ID         string `bson:"_id, omitempty" json:"_id"`
	GuessName  string `bson:"guess_name, omitempty" json:"guess_name"`
	GuessId    string `bson:"guess_id, omitempty" json:"guess_id"`
	IdPersonal string `bson:"id_personal" json:"id_personal"`
	From       string `bson:"from, omitempty" json:"from"`
	To         string `bson:"to, omitempty" json:"to"`
	Date       string `bson:"date, omitempty" json:"date"`
	PlaneID    string `bson:"plane_id" json:"plane_id"`
	AirportID  string `bson:"airport_id" json:"airport_id"`
	SeatName   string `bson:"seat_name" json:"seat_name"`
	CreatedAt  string `bson:"created_at" json:"created_at"`
	UpdatedAt  string `bson:"updated_at" json:"updated_at"`
	SeatID     int64  `bson:"seat_id" json:"seat_id"`
	Status     int64  `bson:"status" json:"status"`
}

type TicketESModel struct {
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
