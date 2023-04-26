package controller

import (
	"grpc/client/booking/requests"
	"grpc/client/booking/responses"
	"grpc/helpers"
	"grpc/pb"
	"net/http"

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
	request := MappingCreateticketRequest(body)
	newTicket, err := h.Client.CreateTicket(c.Request.Context(), request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ErrorResponse{
			ErrorCode:    helpers.ErrorsBooking["CREATE_TICKET_ERROR"],
			ErrorMessage: "CREATE_TICKET_ERROR",
			ErrorSystem:  err.Error(),
		})
		return
	}

	result := MappingCreateTicketResponse(newTicket)
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

	ticketInfo, err := h.Client.FindTicket(c.Request.Context(), &pb.FindTicketRequest{Code: body.Code})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ErrorResponse{
			ErrorCode:    helpers.ErrorsBooking["FIND_TICKET_ERROR"],
			ErrorMessage: "FIND_TICKET_ERROR",
			ErrorSystem:  err.Error(),
		})
		return
	}
	result := responses.TicketInformationResponse{
		Customer: MappingCustomerResponse(ticketInfo),
		Ticket:   MappingTicketResponse(ticketInfo),
	}
	c.JSON(http.StatusOK, helpers.ResponseJSON{
		Status:  true,
		Message: "SUCCESS",
		Error:   "",
		Payload: result,
	})
	return
}
