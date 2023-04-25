package testing

import (
	"context"
	"errors"
	"grpc/server/booking/models"
	"grpc/server/booking/testing/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTicket(t *testing.T) {
	bookingMock := mocks.BookingHandler{}
	t.Run("success", func(t *testing.T) {
		agrs := models.Ticket{
			GuessName: "Nam",
			GuessId:   "123",
			From:      "Ha noi",
			To:        "Da Nang",
			Date:      "15/03/2023",
			PlaneID:   "01",
			AirportID: "HAN-01",
			SeatName:  "A-1",
			SeatID:    011,
			Status:    1,
		}
		bookingMock.On("CreateTicket", context.Background(), &agrs).Return(&agrs, nil)
		got, err := bookingMock.CreateTicket(context.Background(), &agrs)
		assert.NoError(t, err)
		assert.NotNil(t, got)
		bookingMock.AssertExpectations(t)
	})

	t.Run("empty_agrs", func(t *testing.T) {
		agrs := models.Ticket{
			GuessName: "",
			GuessId:   "",
			From:      "",
			To:        "",
			Date:      "",
			PlaneID:   "01",
			AirportID: "HAN-01",
			SeatName:  "A-1",
			SeatID:    011,
			Status:    1,
		}
		bookingMock.On("CreateTicket", context.Background(), &agrs).Return(nil, errors.New("agrs is empty"))
		got, err := bookingMock.CreateTicket(context.Background(), &agrs)
		assert.Error(t, err)
		assert.Nil(t, got)
		bookingMock.AssertExpectations(t)
	})

	t.Run("failed", func(t *testing.T) {
		agrs := models.Ticket{
			GuessName: "Nam",
			GuessId:   "111",
			From:      "Ha noi",
			To:        "Da Nang",
			Date:      "15/03/2023",
			PlaneID:   "01",
			AirportID: "HAN-01",
			SeatName:  "A-1",
			SeatID:    011,
			Status:    1,
		}
		bookingMock.On("CreateTicket", context.Background(), &agrs).Return(nil, errors.New("Cannot create ticket"))
		got, err := bookingMock.CreateTicket(context.Background(), &agrs)
		assert.Error(t, err)
		assert.Nil(t, got)
		bookingMock.AssertExpectations(t)
	})
}
