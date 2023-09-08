package data

import (
	"context"
	"fmt"
	"time"

	"darshub.dev/src/util"
	"github.com/auth0/go-auth0/management"
	"github.com/mitchellh/mapstructure"
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

type CourseRegisterRequest struct {
	UserId string `json:"userId"`
}

type AppMetadataCourses struct {
	Courses []string `json:"courses"`
}

func Create(course *CreateCourseRequest) error {
	ctx, cancel, client := util.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	_, err := client.Database("darshub").Collection("course").InsertOne(ctx, course)
	return err
}

func Find(courseId primitive.ObjectID) (Course, error) {
	var course Course
	ctx, cancel, client := util.GetConnection()
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
	ctx, cancel, client := util.GetConnection()
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
	ctx, cancel, client := util.GetConnection()
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
	ctx, cancel, client := util.GetConnection()
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
	ctx, cancel, client := util.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	_, err := client.Database("darshub").Collection("course").DeleteOne(ctx, bson.M{"_id": courseId})
	return err
}

func RegisterUserToCourse(userId string, courseId string) error {
	client, err := util.GetManagementClient()
	if err != nil {
		return err
	}
	registeredCourses, err := getRegisteredCourses(userId)
	if err != nil {
		return err
	}
	registeredCourses = append(registeredCourses, courseId)
	updateduser := management.User{}
	updateduser.AppMetadata = &map[string]interface{}{
		"courses": registeredCourses,
	}
	updateErr := client.User.Update(context.TODO(), fmt.Sprintf("auth0|%s", userId), &updateduser)
	if err != nil {
		return updateErr
	}
	return nil
}

func getRegisteredCourses(userId string) ([]string, error) {
	user, err := findAuth0User(userId)
	if err != nil {
		return nil, err
	}

	if user.AppMetadata == nil {
		return []string{}, nil
	}
	appMetadata := user.GetAppMetadata()
	appMetadataCourses := &AppMetadataCourses{}
	decodeErr := mapstructure.Decode(appMetadata, appMetadataCourses)
	if decodeErr != nil {
		return nil, err
	}
	return appMetadataCourses.Courses, nil
}

func findAuth0User(userId string) (*management.User, error) {
	client, err := util.GetManagementClient()
	if err != nil {
		return nil, err
	}
	user, getUserErr := client.User.Read(context.TODO(), fmt.Sprintf("auth0|%s", userId))
	if getUserErr != nil {
		return nil, getUserErr
	}
	return user, nil
}
