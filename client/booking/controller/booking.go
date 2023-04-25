package controller

import (
	"grpc/client/booking/requests"
	"grpc/client/booking/responses"
	"grpc/helpers"
	"grpc/pb"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// create ticket
func (h *BookingHandler) CreateTicket(c *gin.Context) {
	body := new(requests.TicketModelRequest)
	if err := c.ShouldBind(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ErrorResponse{
			ErrorCode:    helpers.Errors["MISSING_PARAMS"],
			ErrorMessage: "MISSING_PARAMS",
			ErrorSystem:  err.Error(),
		})
		return
	}

	newTicket, err := h.Client.CreateTicket(c.Request.Context(), &pb.TicketModel{
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
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ErrorResponse{
			ErrorCode:    helpers.ErrorsBooking["CREATE_TICKET_ERROR"],
			ErrorMessage: "CREATE_TICKET_ERROR",
			ErrorSystem:  err.Error(),
		})
		return
	}

	result := responses.TicketResponse{
		ID:         newTicket.Id,
		GuessName:  newTicket.GuessName,
		GuessId:    newTicket.GuessId,
		IdPersonal: newTicket.IdPersonal,
		From:       newTicket.From,
		To:         newTicket.To,
		Date:       newTicket.Date,
		PlaneID:    newTicket.PlaneId,
		AirportID:  newTicket.AirportId,
		SeatName:   newTicket.SeatName,
		SeatID:     newTicket.SeatId,
		CreatedAt:  newTicket.CreatedAt,
		UpdatedAt:  newTicket.UpdatedAt,
		Status:     newTicket.Status,
	}
	c.JSON(http.StatusOK, helpers.ResponseJSON{
		Status:  true,
		Message: "SUCCESS",
		Error:   "",
		Payload: result,
	})
	return
}

// find ticket
func (h *BookingHandler) FindTicket(c *gin.Context) {
	body := new(requests.FindTicketRequest)
	if err := c.ShouldBind(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ErrorResponse{
			ErrorCode:    helpers.Errors["MISSING_PARAMS"],
			ErrorMessage: "Missing Params",
			ErrorSystem:  err.Error(),
		})
		return
	}

	ticketInformation, err := h.Client.FindTicket(c.Request.Context(), &pb.FindTicketRequest{
		Code: body.Code,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ErrorResponse{
			ErrorCode:    helpers.ErrorsBooking["FIND_TICKET_ERROR"],
			ErrorMessage: "FIND_TICKET_ERROR",
			ErrorSystem:  err.Error(),
		})
		return
	}

	ticket := responses.TicketResponse{
		ID:         ticketInformation.TicketDetail.Id,
		GuessName:  ticketInformation.TicketDetail.GuessName,
		GuessId:    ticketInformation.TicketDetail.GuessId,
		IdPersonal: ticketInformation.TicketDetail.IdPersonal,
		From:       ticketInformation.TicketDetail.From,
		To:         ticketInformation.TicketDetail.To,
		Date:       ticketInformation.TicketDetail.Date,
		CreatedAt:  ticketInformation.TicketDetail.CreatedAt,
		UpdatedAt:  ticketInformation.TicketDetail.UpdatedAt,
		SeatName:   ticketInformation.TicketDetail.SeatName,
		SeatID:     ticketInformation.TicketDetail.SeatId,
		PlaneID:    ticketInformation.TicketDetail.PlaneId,
		Status:     ticketInformation.TicketDetail.Status,
	}
	customer := responses.CustomerResponse{
		ID:         ticketInformation.CustomerDetail.Id,
		Name:       ticketInformation.CustomerDetail.Name,
		IdPersonal: ticketInformation.CustomerDetail.IdPersonal,
		Age:        ticketInformation.CustomerDetail.Age,
		Address:    ticketInformation.CustomerDetail.Address,
		CreatedAt:  ticketInformation.CustomerDetail.CreatedAt,
		UpdatedAt:  ticketInformation.CustomerDetail.UpdatedAt,
		Tags:       ticketInformation.CustomerDetail.Tags,
	}

	result := responses.TicketInformationResponse{
		Customer: customer,
		Ticket:   ticket,
	}
	c.JSON(http.StatusOK, helpers.ResponseJSON{
		Status:  true,
		Message: "SUCCESS",
		Error:   "",
		Payload: result,
	})
	return
}
