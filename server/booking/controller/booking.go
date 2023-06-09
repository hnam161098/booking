package controller

import (
	"context"
	"fmt"
	"grpc/helpers"
	"grpc/pb"
	"grpc/server/booking/models"
	"log"
	"time"
)

func (c *BookingHandlerStruct) CreateTicket(ctx context.Context, in *pb.TicketModel) (*pb.TicketModel, error) {
	// find customer
	customer, errC := c.CustomerClient.GetCustomer(ctx, &pb.GetCustomerRequest{IdPersonal: in.IdPersonal})
	if errC != nil {
		fmt.Println("errC: ", errC)
		return nil, errC
	}
	if customer.Id == "" {
		return nil, nil
	}
	in.Status = 1
	req := models.Ticket{
		ID:         in.Id,
		GuessId:    customer.Id,
		GuessName:  customer.Name,
		IdPersonal: in.IdPersonal,
		From:       in.From,
		To:         in.To,
		Date:       in.Date,
		PlaneID:    in.PlaneId,
		SeatName:   in.SeatName,
		SeatID:     in.SeatId,
		AirportID:  in.AirportId,
		Status:     in.Status,
		CreatedAt:  time.Now().Format(time.RFC3339),
		UpdatedAt:  time.Now().Format(time.RFC3339),
	}
	// create ticket
	ticket, err := c.BookingHandler.CreateTicket(ctx, &req)
	if err != nil {
		return nil, err
	}
	result := MappingCreateTicketToPbModel(ticket)
	// sync to ES
	go func() {
		err1 := helpers.InsertDocumentES("ticket", result, ticket.ID)
		if err1 != nil {
			log.Println(err1)
		}
	}()
	return result, nil
}

func (c *BookingHandlerStruct) FindTicket(ctx context.Context, in *pb.FindTicketRequest) (*pb.TicketInformation, error) {
	// code = id ticket
	ticket, errT := c.BookingHandler.FindTicket(ctx, in.Code)
	if errT != nil {
		return nil, errT
	}
	customer, errC := c.CustomerClient.GetCustomer(ctx, &pb.GetCustomerRequest{IdPersonal: ticket.IdPersonal})
	if errC != nil {
		return nil, errC
	}
	ticket.GuessName = customer.Name
	ticketInformation := MappingFindTicketToPbModel(ticket)
	// update ticket ES
	go func() {
		queryES := map[string]interface{}{
			"guess_name":  customer.Name,
			"id_personal": customer.IdPersonal,
		}
		err1 := helpers.UpdateES("ticket", queryES, ticket.ID)
		if err1 != nil {
			log.Printf("Cannot update ES: %v", err1)
		}
	}()
	result := pb.TicketInformation{
		CustomerDetail: customer,
		TicketDetail:   ticketInformation,
	}

	return &result, nil
}
