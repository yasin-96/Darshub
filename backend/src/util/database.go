package util

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetConnection() (context.Context, context.CancelFunc, *mongo.Client) {
	client, err := mongo.NewClient(options.Client().
		ApplyURI(os.Getenv("MONGO_CONNECTION_STRING")))

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
