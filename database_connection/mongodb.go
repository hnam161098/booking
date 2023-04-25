package database_connection

import (
	"context"
	"fmt"
	"grpc/config"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var MongoDB MongoInstance

func init() {
	opts := options.Client().ApplyURI(config.MONGODB_ADDRESS["HOST"])
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database(config.MONGODB_ADDRESS["DATABASE"])
	MongoDB = MongoInstance{
		Client: client,
		Db:     db,
	}

	fmt.Println("CONNECT MONGODB SUCCESS!")
}
