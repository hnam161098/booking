// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: protoc/booking.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BookingClient is the client API for Booking service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookingClient interface {
	CreateTicket(ctx context.Context, in *TicketModel, opts ...grpc.CallOption) (*TicketModel, error)
	FindTicket(ctx context.Context, in *FindTicketRequest, opts ...grpc.CallOption) (*TicketInformation, error)
}

type bookingClient struct {
	cc grpc.ClientConnInterface
}

func NewBookingClient(cc grpc.ClientConnInterface) BookingClient {
	return &bookingClient{cc}
}

func (c *bookingClient) CreateTicket(ctx context.Context, in *TicketModel, opts ...grpc.CallOption) (*TicketModel, error) {
	out := new(TicketModel)
	err := c.cc.Invoke(ctx, "/asm.Booking/CreateTicket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingClient) FindTicket(ctx context.Context, in *FindTicketRequest, opts ...grpc.CallOption) (*TicketInformation, error) {
	out := new(TicketInformation)
	err := c.cc.Invoke(ctx, "/asm.Booking/FindTicket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookingServer is the server API for Booking service.
// All implementations must embed UnimplementedBookingServer
// for forward compatibility
type BookingServer interface {
	CreateTicket(context.Context, *TicketModel) (*TicketModel, error)
	FindTicket(context.Context, *FindTicketRequest) (*TicketInformation, error)
	mustEmbedUnimplementedBookingServer()
}

// UnimplementedBookingServer must be embedded to have forward compatible implementations.
type UnimplementedBookingServer struct {
}

func (UnimplementedBookingServer) CreateTicket(context.Context, *TicketModel) (*TicketModel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTicket not implemented")
}
func (UnimplementedBookingServer) FindTicket(context.Context, *FindTicketRequest) (*TicketInformation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindTicket not implemented")
}
func (UnimplementedBookingServer) mustEmbedUnimplementedBookingServer() {}

// UnsafeBookingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookingServer will
// result in compilation errors.
type UnsafeBookingServer interface {
	mustEmbedUnimplementedBookingServer()
}

func RegisterBookingServer(s grpc.ServiceRegistrar, srv BookingServer) {
	s.RegisterService(&Booking_ServiceDesc, srv)
}

func _Booking_CreateTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TicketModel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServer).CreateTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/asm.Booking/CreateTicket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServer).CreateTicket(ctx, req.(*TicketModel))
	}
	return interceptor(ctx, in, info, handler)
}

func _Booking_FindTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServer).FindTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/asm.Booking/FindTicket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServer).FindTicket(ctx, req.(*FindTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Booking_ServiceDesc is the grpc.ServiceDesc for Booking service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Booking_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "asm.Booking",
	HandlerType: (*BookingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTicket",
			Handler:    _Booking_CreateTicket_Handler,
		},
		{
			MethodName: "FindTicket",
			Handler:    _Booking_FindTicket_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protoc/booking.proto",
}
