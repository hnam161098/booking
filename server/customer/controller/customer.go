package controller

import (
	"context"
	"errors"
	"grpc/helpers"
	"grpc/pb"
	"grpc/server/customer/models"
	"log"
	"time"
)

func MappingResultEStoCustomerPbModel(m map[string]interface{}) *pb.CustomerModel {
	defer func() {
		r := recover()
		if r != nil {
			log.Printf("message panic: %v", r)
			return
		}
	}()

	rsES := m["hits"].(map[string]interface{})["hits"].([]interface{})[0].(map[string]interface{})["_source"]
	listTags := []string{}
	for _, item := range rsES.(map[string]interface{})["tags"].([]interface{}) {
		listTags = append(listTags, item.(string))
	}

	result := &pb.CustomerModel{
		Id:         rsES.(map[string]interface{})["id"].(string),
		Name:       rsES.(map[string]interface{})["name"].(string),
		IdPersonal: rsES.(map[string]interface{})["id_personal"].(string),
		Age:        rsES.(map[string]interface{})["age"].(string),
		Address:    rsES.(map[string]interface{})["address"].(string),
		Tags:       listTags,
		CreatedAt:  rsES.(map[string]interface{})["created_at"].(string),
		UpdatedAt:  rsES.(map[string]interface{})["updated_at"].(string),
	}
	return result

}

func (c *CustomerHandlerStruct) CreateCustomer(ctx context.Context, in *pb.CustomerModel) (*pb.CustomerModel, error) {
	req := models.Customer{
		Name:       in.Name,
		IdPersonal: in.IdPersonal,
		Age:        in.Age,
		Address:    in.Address,
		Tags:       in.Tags,
		CreatedAt:  time.Now().Format(time.RFC3339),
		UpdatedAt:  time.Now().Format(time.RFC3339),
	}
	customer, err := c.CustomerHandler.CreateCustomer(ctx, &req)
	if err != nil {
		return nil, err
	}

	// sync to ES
	go func() {
		reqES := models.CustomerESModel{
			ID:         customer.ID,
			Name:       customer.Name,
			IdPersonal: customer.IdPersonal,
			Age:        customer.Age,
			Address:    customer.Address,
			Tags:       customer.Tags,
			CreatedAt:  customer.CreatedAt,
			UpdatedAt:  customer.UpdatedAt,
		}
		err1 := helpers.InsertDocumentES(models.IndexCustomer, reqES, customer.ID)
		if err1 != nil {
			log.Println(err1)
		}
	}()

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

	return result, nil
}

func (c *CustomerHandlerStruct) GetCustomer(ctx context.Context, in *pb.GetCustomerRequest) (*pb.CustomerModel, error) {

	// get from Elasticsearch
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": map[string]interface{}{
					"match": map[string]interface{}{
						"id_personal": in.IdPersonal,
					},
				},
			},
		},
	}
	rsES, err := helpers.QueryES(models.IndexCustomer, query)
	if err != nil {
		return nil, err
	}

	if rsES["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64) == 0 {
		return nil, errors.New("Customer not exist")
	}

	customer := MappingResultEStoCustomerPbModel(rsES)
	return customer, nil
}

func (c *CustomerHandlerStruct) UpdateCustomer(ctx context.Context, in *pb.CustomerModel) (*pb.CustomerModel, error) {

	req := models.Customer{
		ID:      in.Id,
		Name:    in.Name,
		Age:     in.Age,
		Address: in.Address,
		Tags:    in.Tags,
	}
	customer, err := c.CustomerHandler.UpdateCustomer(ctx, &req)
	if err != nil {
		return nil, err
	}
	// sync ES
	go func() {
		queryES := map[string]interface{}{
			"id":         customer.ID,
			"name":       customer.Name,
			"age":        customer.Age,
			"address":    customer.Address,
			"tags":       customer.Tags,
			"created_at": customer.CreatedAt,
			"updated_at": customer.UpdatedAt,
		}

		errUpdate := helpers.UpdateES(models.IndexCustomer, queryES, customer.ID)
		if errUpdate != nil {
			log.Println(errUpdate)
		}
	}()

	result := &pb.CustomerModel{
		Id:        customer.ID,
		Name:      customer.Name,
		Age:       customer.Age,
		Address:   customer.Address,
		Tags:      customer.Tags,
		CreatedAt: customer.CreatedAt,
		UpdatedAt: customer.UpdatedAt,
	}
	return result, nil

}

func (c *CustomerHandlerStruct) DeleteCustomer(ctx context.Context, in *pb.CustomerModel) (*pb.DeleteCustomerResponse, error) {
	req := models.Customer{
		ID: in.Id,
	}

	// delete in mongoDB
	count, err := c.CustomerHandler.DeleteCustomer(ctx, &req)
	if err != nil {
		return nil, err
	}

	// delete in ES
	go func() {
		queryES := map[string]interface{}{
			"query": map[string]interface{}{
				"bool": map[string]interface{}{
					"must": map[string]interface{}{
						"match": map[string]interface{}{
							"id": in.Id,
						},
					},
				},
			},
		}
		errES := helpers.DeleteES(models.IndexCustomer, queryES)
		if errES != nil {
			log.Println(errES)
		}
	}()

	return &pb.DeleteCustomerResponse{
		Count: int64(count),
	}, nil
}

func (c *CustomerHandlerStruct) GetAllCustomer(ctx context.Context, in *pb.GetAllCustomerRequest) (*pb.AllCustomerResponse, error) {
	// get from ES
	queryES := map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
	}
	totalRecord, err := helpers.QueryES(models.IndexCustomer, queryES)
	if err != nil {
		return nil, err
	}

	var result []*pb.CustomerModel
	listRecord := totalRecord["hits"].(map[string]interface{})["hits"].([]interface{})
	for _, customer := range listRecord {
		listTags := []string{}
		if customer.(map[string]interface{})["_source"].(map[string]interface{})["tags"] != nil {
			for _, item := range customer.(map[string]interface{})["_source"].(map[string]interface{})["tags"].([]interface{}) {
				listTags = append(listTags, item.(string))
			}
		}
		cus := &pb.CustomerModel{
			Id:         customer.(map[string]interface{})["_source"].(map[string]interface{})["id"].(string),
			Name:       customer.(map[string]interface{})["_source"].(map[string]interface{})["name"].(string),
			IdPersonal: customer.(map[string]interface{})["_source"].(map[string]interface{})["id_personal"].(string),
			Age:        customer.(map[string]interface{})["_source"].(map[string]interface{})["age"].(string),
			Tags:       listTags,
			Address:    customer.(map[string]interface{})["_source"].(map[string]interface{})["address"].(string),
			CreatedAt:  customer.(map[string]interface{})["_source"].(map[string]interface{})["created_at"].(string),
			UpdatedAt:  customer.(map[string]interface{})["_source"].(map[string]interface{})["updated_at"].(string),
		}
		result = append(result, cus)
	}

	return &pb.AllCustomerResponse{
		Customers: result,
	}, nil
}

func (c *CustomerHandlerStruct) AddTagsCustomer(ctx context.Context, in *pb.AddTagsCustomerRequest) (*pb.CustomerModel, error) {
	req := models.AddTags{
		ID:   in.Id,
		Tags: in.Tags,
	}

	cus, err := c.CustomerHandler.AddTagsCustomer(ctx, &req)
	if err != nil {
		return nil, err
	}

	// sync to ES
	go func() {
		queryES := map[string]interface{}{
			"id":         cus.ID,
			"tags":       cus.Tags,
			"updated_at": cus.UpdatedAt,
		}

		errUpdate := helpers.UpdateES(models.IndexCustomer, queryES, cus.ID)
		if errUpdate != nil {
			log.Println(errUpdate)
		}
	}()

	result := &pb.CustomerModel{
		Id:        cus.ID,
		Name:      cus.Name,
		Age:       cus.Age,
		Address:   cus.Address,
		Tags:      cus.Tags,
		CreatedAt: cus.CreatedAt,
		UpdatedAt: time.Now().Format(time.RFC3339),
	}
	return result, nil
}

func (c *CustomerHandlerStruct) DeleteTagsOfCustomer(ctx context.Context, in *pb.DeleteTagsOfCustomerRequest) (*pb.CustomerModel, error) {
	req := models.DeleteTags{
		ID:   in.Id,
		Tags: in.Tags,
	}
	cus, err := c.CustomerHandler.DeleteTagsOfCustomer(ctx, &req)
	if err != nil {
		return nil, err
	}
	return &pb.CustomerModel{
		Id:        cus.ID,
		Name:      cus.Name,
		Age:       cus.Age,
		Tags:      cus.Tags,
		Address:   cus.Address,
		CreatedAt: cus.CreatedAt,
		UpdatedAt: time.Now().Format(time.RFC3339),
	}, nil

}
