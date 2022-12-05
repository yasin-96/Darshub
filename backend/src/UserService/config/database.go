package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetConnection() (context.Context, context.CancelFunc, *mongo.Client) {
	client, err := mongo.NewClient(options.Client().
		ApplyURI("mongodb+srv://dhub:OiPe7pU8kxaIVhBx@dhcluster001.c17aj.mongodb.net/?retryWrites=true&w=majority"))

	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return ctx, cancel, client
}
