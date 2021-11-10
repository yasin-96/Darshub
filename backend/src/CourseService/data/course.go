package data

import (
	"log"
	"time"

	"dev.azure.com/learn-website-orga/_git/learn-website/backend/src/UserService/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Course struct {
	ID          primitive.ObjectID   `bson:"_id"`
	Name        string               `bson:"name"`
	Description string               `bson:"description"`
	Duration    time.Time            `bson:"duration"`
	Level       string               `bson:"level"`
	Content     []primitive.ObjectID `bson:"content"`
	Author      string               `bson:"author"`
	Released    time.Time            `bson:"released"`
	LastUpdate  time.Time            `bson:"lastUpdate"`
}

func Create(course *Course) {
	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	course.ID = primitive.NewObjectID()

	_, err := client.Database("darshub").Collection("course").InsertOne(ctx, course)
	if err != nil {
		log.Printf("Could not save course: %v", err)
	}
}

func Find(courseId primitive.ObjectID) Course {
	var course Course
	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	err := client.Database("darshub").Collection("course").FindOne(ctx, bson.M{"_id": courseId}).Decode(&course)
	if err != nil {
		log.Fatal(err)
	}
	return course
}
