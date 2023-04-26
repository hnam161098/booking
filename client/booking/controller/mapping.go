package controller

import (
	"grpc/client/booking/requests"
	"grpc/client/booking/responses"
	"grpc/pb"
	"time"
)

func MappingTicketResponse(ticketInfo *pb.TicketInformation) responses.TicketResponse {
	result := responses.TicketResponse{
		ID:         ticketInfo.TicketDetail.Id,
		GuessName:  ticketInfo.TicketDetail.GuessName,
		GuessId:    ticketInfo.TicketDetail.GuessId,
		IdPersonal: ticketInfo.TicketDetail.IdPersonal,
		From:       ticketInfo.TicketDetail.From,
		To:         ticketInfo.TicketDetail.To,
		Date:       ticketInfo.TicketDetail.Date,
		CreatedAt:  ticketInfo.TicketDetail.CreatedAt,
		UpdatedAt:  ticketInfo.TicketDetail.UpdatedAt,
		SeatName:   ticketInfo.TicketDetail.SeatName,
		SeatID:     ticketInfo.TicketDetail.SeatId,
		AirportID:  ticketInfo.TicketDetail.AirportId,
		PlaneID:    ticketInfo.TicketDetail.PlaneId,
		Status:     ticketInfo.TicketDetail.Status,
	}
	return result
}

func MappingCustomerResponse(ticketInfo *pb.TicketInformation) responses.CustomerResponse {
	result := responses.CustomerResponse{
		ID:         ticketInfo.CustomerDetail.Id,
		Name:       ticketInfo.CustomerDetail.Name,
		IdPersonal: ticketInfo.CustomerDetail.IdPersonal,
		Age:        ticketInfo.CustomerDetail.Age,
		Address:    ticketInfo.CustomerDetail.Address,
		CreatedAt:  ticketInfo.CustomerDetail.CreatedAt,
		UpdatedAt:  ticketInfo.CustomerDetail.UpdatedAt,
		Tags:       ticketInfo.CustomerDetail.Tags,
	}
	return result
}

func MappingCreateticketRequest(body *requests.TicketModelRequest) *pb.TicketModel {
	result := &pb.TicketModel{
		GuessId:    body.GuessId,
		IdPersonal: body.IdPersonal,
		From:       body.From,
		To:         body.To,
		Date:       body.Date,
		SeatName:   body.SeatName,
		SeatId:     body.SeatID,
		PlaneId:    body.PlaneID,
		AirportId:  body.AirportID,
		Status:     body.Status,
		CreatedAt:  time.Now().Format(time.RFC3339),
		UpdatedAt:  time.Now().Format(time.RFC3339),
	}
	return result
}

func MappingCreateTicketResponse(ticket *pb.TicketModel) responses.TicketResponse {
	result := responses.TicketResponse{
		ID:         ticket.Id,
		GuessName:  ticket.GuessName,
		GuessId:    ticket.GuessId,
		IdPersonal: ticket.IdPersonal,
		From:       ticket.From,
		To:         ticket.To,
		Date:       ticket.Date,
		PlaneID:    ticket.PlaneId,
		AirportID:  ticket.AirportId,
		SeatName:   ticket.SeatName,
		SeatID:     ticket.SeatId,
		CreatedAt:  ticket.CreatedAt,
		UpdatedAt:  ticket.UpdatedAt,
		Status:     ticket.Status,
	}
	return result
}
