package handler

import (
	"grpc/database_connection"

	"go.mongodb.org/mongo-driver/mongo"
)

type DbManager struct {
	collection *mongo.Collection
}

func NewDBmanager() BookingHandler {
	var coll = database_connection.MongoDB.Db.Collection("booking")
	return &DbManager{
		collection: coll,
	}
}
