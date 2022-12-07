package data

import (
	"log"
	"time"

	"dev.azure.com/learn-website-orga/_git/learn-website/src/UserService/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Course struct {
	ID          primitive.ObjectID   `json:"id" bson:"_id"`
	Name        string               `json:"name" bson:"name"`
	Description string               `json:"description" bson:"description"`
	Duration    time.Time            `json:"duration" bson:"duration"`
	Level       string               `json:"level" bson:"level"`
	Content     []primitive.ObjectID `json:"content" bson:"content"`
	Author      string               `json:"author" bson:"author"`
	Released    time.Time            `json:"released" bson:"released"`
	LastUpdate  time.Time            `json:"lastUpdate" bson:"lastUpdate"`
}

type CreateCourseRequest struct {
	Name        string               `json:"name"`
	Description string               `json:"description"`
	Duration    string               `json:"duration"`
	Level       int                  `json:"level"`
	Content     []primitive.ObjectID `json:"content"`
	Author      string               `json:"author"`
	Released    time.Time            `json:"released"`
	LastUpdate  time.Time            `json:"lastUpdate"`
}

type UpdateCourseRequest struct {
	Name        string               `json:"name"`
	Description string               `json:"description"`
	Duration    string               `json:"duration"`
	Level       int                  `json:"level"`
	Content     []primitive.ObjectID `json:"content"`
	Author      string               `json:"author"`
	LastUpdate  time.Time            `json:"lastUpdate"`
}

func Create(course *CreateCourseRequest) {
	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

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
		log.Print(err)
		return Course{}
	}
	return course
}

func GetAllCourses() []Course {
	var courses []Course
	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	cur, err := client.Database("darshub").Collection("course").Find(ctx, bson.D{})
	if err != nil {
		log.Print(err)
	}
	cur.All(ctx, &courses)
	return courses
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
		"lastUpdate":  updatedCourse.LastUpdate,
	}

	_, err := client.Database("darshub").Collection("course").ReplaceOne(ctx, bson.M{"_id": courseId}, update)
	if err != nil {
		log.Print(err)
		return Course{}
	}

	return Find(courseId)
}

func Delete(courseId primitive.ObjectID) {
	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	_, err := client.Database("darshub").Collection("course").DeleteOne(ctx, bson.M{"_id": courseId})
	if err != nil {
		log.Print(err)
		return
	}
	log.Print("Course was successfully deleted")
}
