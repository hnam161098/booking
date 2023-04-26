package controller

import (
	"grpc/pb"
	"grpc/server/customer/models"
	"time"
)

func MappingCreateCustomerToPbModel(customer *models.Customer) *pb.CustomerModel {
	result := &pb.CustomerModel{
		Id:         customer.ID,
		Name:       customer.Name,
		IdPersonal: customer.IdPersonal,
		Address:    customer.Address,
		Age:        customer.Age,
		Tags:       customer.Tags,
		CreatedAt:  customer.CreatedAt,
		UpdatedAt:  customer.UpdatedAt,
	}
	return result
}

func MappingAddTagsCustomerToPbModel(customer *models.Customer) *pb.CustomerModel {
	result := &pb.CustomerModel{
		Id:         customer.ID,
		Name:       customer.Name,
		IdPersonal: customer.IdPersonal,
		Age:        customer.Age,
		Address:    customer.Address,
		Tags:       customer.Tags,
		CreatedAt:  customer.CreatedAt,
		UpdatedAt:  time.Now().Format(time.RFC3339),
	}
	return result
}

func MappingUpdateCustomerToPbModel(customer *models.Customer) *pb.CustomerModel {
	result := &pb.CustomerModel{
		Id:         customer.ID,
		Name:       customer.Name,
		IdPersonal: customer.IdPersonal,
		Age:        customer.Age,
		Address:    customer.Address,
		Tags:       customer.Tags,
		CreatedAt:  customer.CreatedAt,
		UpdatedAt:  time.Now().Format(time.RFC3339),
	}
	return result
}

func MappingDeleteTagsCustomerToPbModel(customer *models.Customer) *pb.CustomerModel {
	result := &pb.CustomerModel{
		Id:         customer.ID,
		Name:       customer.Name,
		IdPersonal: customer.IdPersonal,
		Age:        customer.Age,
		Address:    customer.Address,
		Tags:       customer.Tags,
		CreatedAt:  customer.CreatedAt,
		UpdatedAt:  time.Now().Format(time.RFC3339),
	}
	return result
}
