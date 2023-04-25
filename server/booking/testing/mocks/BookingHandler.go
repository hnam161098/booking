// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	models "grpc/server/booking/models"
)

// BookingHandler is an autogenerated mock type for the BookingHandler type
type BookingHandler struct {
	mock.Mock
}

// CreateTicket provides a mock function with given fields: ctx, model
func (_m *BookingHandler) CreateTicket(ctx context.Context, model *models.Ticket) (*models.Ticket, error) {
	ret := _m.Called(ctx, model)

	var r0 *models.Ticket
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Ticket) (*models.Ticket, error)); ok {
		return rf(ctx, model)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *models.Ticket) *models.Ticket); ok {
		r0 = rf(ctx, model)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Ticket)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *models.Ticket) error); ok {
		r1 = rf(ctx, model)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindTicket provides a mock function with given fields: ctx, code
func (_m *BookingHandler) FindTicket(ctx context.Context, code string) (*models.Ticket, error) {
	ret := _m.Called(ctx, code)

	var r0 *models.Ticket
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*models.Ticket, error)); ok {
		return rf(ctx, code)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.Ticket); ok {
		r0 = rf(ctx, code)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Ticket)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, code)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewBookingHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewBookingHandler creates a new instance of BookingHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBookingHandler(t mockConstructorTestingTNewBookingHandler) *BookingHandler {
	mock := &BookingHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
