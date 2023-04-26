package controller

import (
	"grpc/pb"
	"grpc/server/booking/models"
)

func MappingCreateTicketToPbModel(ticket *models.Ticket) *pb.TicketModel {
	result := &pb.TicketModel{
		Id:         ticket.ID,
		GuessId:    ticket.GuessId,
		GuessName:  ticket.GuessName,
		IdPersonal: ticket.IdPersonal,
		From:       ticket.From,
		To:         ticket.To,
		Date:       ticket.Date,
		PlaneId:    ticket.PlaneID,
		SeatName:   ticket.SeatName,
		SeatId:     ticket.SeatID,
		AirportId:  ticket.AirportID,
		Status:     ticket.Status,
		CreatedAt:  ticket.CreatedAt,
		UpdatedAt:  ticket.UpdatedAt,
	}
	return result
}

func MappingFindTicketToPbModel(ticket *models.Ticket) *pb.TicketModel {
	result := &pb.TicketModel{
		Id:         ticket.ID,
		GuessId:    ticket.GuessId,
		GuessName:  ticket.GuessName,
		IdPersonal: ticket.IdPersonal,
		From:       ticket.From,
		To:         ticket.To,
		Date:       ticket.Date,
		PlaneId:    ticket.PlaneID,
		SeatName:   ticket.SeatName,
		SeatId:     ticket.SeatID,
		AirportId:  ticket.AirportID,
		Status:     ticket.Status,
		CreatedAt:  ticket.CreatedAt,
		UpdatedAt:  ticket.UpdatedAt,
	}
	return result
}
