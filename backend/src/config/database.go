package config

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type DatabaseHelper struct {
	Context    context.Context
	Collection *mongo.Collection
}

var helper DatabaseHelper

/*func InitializeDbConnection() {
	client, err := mongo.NewClient(options.Client().
		ApplyURI("mongodb+srv://dhub:OiPe7pU8kxaIVhBx@dhcluster001.c17aj.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))

	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)

	darshubDb := client.Database("darshub")
	userCollection := darshubDb.Collection("user")
	setHelperFields(ctx, userCollection)
}*/

func setHelperFields(ctx context.Context, col *mongo.Collection) {
	helper.Context = ctx
	helper.Collection = col
}

func GetHelper() DatabaseHelper {
	return helper
}
