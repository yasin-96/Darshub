package data

import (
	"time"

	"darshub.dev/src/UserService/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Course struct {
	ID          primitive.ObjectID   `json:"id" bson:"_id"`
	Name        string               `json:"name" bson:"name"`
	Description string               `json:"description" bson:"description"`
	Duration    string               `json:"duration" bson:"duration"`
	Level       int32                `json:"level" bson:"level"`
	Chapters    []primitive.ObjectID `json:"chapters" bson:"chapters"`
	AuthorId    string               `json:"authorId" bson:"authorId"`
	Released    time.Time            `json:"released" bson:"released"`
	LastUpdate  time.Time            `json:"lastUpdate" bson:"lastUpdate"`
}

type CreateCourseRequest struct {
	Name        string               `json:"name"`
	Description string               `json:"description"`
	Duration    string               `json:"duration"`
	Level       int                  `json:"level"`
	Chapters    []primitive.ObjectID `json:"chapters"`
	AuthorId    string               `json:"authorId"`
	Released    time.Time            `json:"released"`
	LastUpdate  time.Time            `json:"lastUpdate"`
}

type UpdateCourseRequest struct {
	Name        string               `json:"name"`
	Description string               `json:"description"`
	Duration    string               `json:"duration"`
	Level       int                  `json:"level"`
	Chapters    []primitive.ObjectID `json:"content"`
	AuthorId    string               `json:"authorId"`
	LastUpdate  time.Time            `json:"lastUpdate"`
}

func Create(course *CreateCourseRequest) error {
	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	_, err := client.Database("darshub").Collection("course").InsertOne(ctx, course)
	return err
}

func Find(courseId primitive.ObjectID) (Course, error) {
	var course Course
	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	err := client.Database("darshub").Collection("course").FindOne(ctx, bson.M{"_id": courseId}).Decode(&course)
	if err != nil {
		return Course{}, err
	}
	return course, nil
}

func GetCoursesByAuthorId(authorId string) ([]Course, error) {
	var courses []Course
	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	filter := bson.D{{Key: "authorid", Value: authorId}}
	cur, err := client.Database("darshub").Collection("course").Find(ctx, filter)
	if err != nil {
		return []Course{}, err
	}
	cur.All(ctx, &courses)
	return courses, nil
}

func GetAllCourses() ([]Course, error) {
	var courses []Course
	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	cur, err := client.Database("darshub").Collection("course").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	cur.All(ctx, &courses)
	return courses, nil
}

func Update(courseId primitive.ObjectID, updatedCourse *UpdateCourseRequest) (Course, error) {
	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	update := bson.M{
		"name":        updatedCourse.Name,
		"description": updatedCourse.Description,
		"duration":    updatedCourse.Duration,
		"level":       updatedCourse.Level,
		"chapters":    updatedCourse.Chapters,
		"authorId":    updatedCourse.AuthorId,
		"lastUpdate":  updatedCourse.LastUpdate,
	}

	_, err := client.Database("darshub").Collection("course").ReplaceOne(ctx, bson.M{"_id": courseId}, update)
	if err != nil {
		return Course{}, err
	}
	course, respErr := Find(courseId)
	if respErr != nil {
		return Course{}, err
	}
	return course, nil
}

func Delete(courseId primitive.ObjectID) error {
	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	_, err := client.Database("darshub").Collection("course").DeleteOne(ctx, bson.M{"_id": courseId})
	return err
}
