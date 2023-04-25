package handler

import (
	"context"
	"grpc/server/booking/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

type BookingHandler interface {
	CreateTicket(ctx context.Context, model *models.Ticket) (*models.Ticket, error)
	FindTicket(ctx context.Context, code string) (*models.Ticket, error)
}

func (m *DbManager) CreateTicket(ctx context.Context, model *models.Ticket) (*models.Ticket, error) {
	filter := bson.M{
		"guess_id":    model.GuessId,
		"guess_name":  model.GuessName,
		"id_personal": model.IdPersonal,
		"from":        model.From,
		"to":          model.To,
		"date":        model.Date,
		"plane_id":    model.PlaneID,
		"airport_id":  model.AirportID,
		"seat_name":   model.SeatName,
		"seat_id":     model.SeatID,
		"created_at":  model.CreatedAt,
		"updated_at":  model.UpdatedAt,
	}
	rs, err := m.collection.InsertOne(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	return &models.Ticket{
		ID:         rs.InsertedID.(primitive.ObjectID).Hex(),
		GuessName:  model.GuessName,
		GuessId:    model.GuessId,
		IdPersonal: model.IdPersonal,
		From:       model.From,
		To:         model.To,
		Date:       model.Date,
		PlaneID:    model.PlaneID,
		SeatID:     model.SeatID,
		SeatName:   model.SeatName,
		AirportID:  model.AirportID,
		CreatedAt:  model.CreatedAt,
		UpdatedAt:  model.UpdatedAt,
	}, nil
}

func (m *DbManager) FindTicket(ctx context.Context, code string) (*models.Ticket, error) {
	id, _ := primitive.ObjectIDFromHex(code)
	filter := bson.M{
		"_id": id,
	}
	var ticket models.Ticket
	err := m.collection.FindOne(context.TODO(), filter).Decode(&ticket)
	if err != nil {
		return nil, err
	}
	return &ticket, nil
}
