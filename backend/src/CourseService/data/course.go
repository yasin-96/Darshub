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

type CreateCourseRequest struct {
	ID          primitive.ObjectID   `json:"id"`
	Name        string               `json:"name"`
	Description string               `json:"description"`
	Duration    time.Time            `json:"duration"`
	Level       string               `json:"level"`
	Content     []primitive.ObjectID `json:"content"`
	Author      string               `json:"author"`
	Released    time.Time            `json:"released"`
	LastUpdate  time.Time            `json:"lastUpdate"`
}

type UpdateCourseRequest struct {
	Name        string               `json:"name"`
	Description string               `json:"description"`
	Duration    time.Time            `json:"duration"`
	Level       string               `json:"level"`
	Content     []primitive.ObjectID `json:"content"`
	Author      string               `json:"author"`
	Released    time.Time            `json:"released"`
	LastUpdate  time.Time            `json:"lastUpdate"`
}

func Create(course *CreateCourseRequest) {
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

func Update(courseId primitive.ObjectID, updatedCourse *UpdateCourseRequest) Course {
	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	update := bson.M{
		"name":        updatedCourse.Name,
		"description": updatedCourse.Description,
		"duration":    updatedCourse.Duration,
		"level":       updatedCourse.Level,
		"content":     updatedCourse.Content,
		"author":      updatedCourse.Author,
		"released":    updatedCourse.Released,
		"lastUpdate":  updatedCourse.LastUpdate,
	}

	_, err := client.Database("darshub").Collection("course").ReplaceOne(ctx, bson.M{"_id": courseId}, update)
	if err != nil {
		log.Fatal(err)
	}

	return Find(courseId)
}

func Delete(courseId primitive.ObjectID) {
	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	_, err := client.Database("darshub").Collection("course").DeleteOne(ctx, bson.M{"_id": courseId})
	if err != nil {
		log.Fatal(err)
	}
}