package data

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

type AppMetadata struct {
	Courses         []string               `json:"courses"`
	CoursesProgress map[string]interface{} `json:"coursesProgress"`
}

type CourseProgressRequest struct {
	UserId       string `json:"userId"`
	CourseId     string `json:"courseId"`
	ChapterId    string `json:"chapterId"`
	SubchapterId string `json:"subchapterId"`
}

type Chapterdata struct {
	ChapterId    string `json:"chapterId"`
	SubchapterId string `json:"subchapterId"`
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
	user, findUserErr := findAuth0User(userId)
	if findUserErr != nil {
		return findUserErr
	}
	metadata, getMetadataErr := getUserAppMetadata(userId)
	if getMetadataErr != nil {
		return getMetadataErr
	}

	registeredCourses := metadata.Courses
	registeredCourses = append(registeredCourses, courseId)
	updateduser := management.User{}
	metadataMap := make(map[string]interface{})
	metadataMap["courses"] = &registeredCourses
	updateduser.AppMetadata = &metadataMap
	updateErr := client.User.Update(context.TODO(), user.GetID(), &updateduser)
	if err != nil {
		return updateErr
	}
	return nil
}

func SetCourseProgress(courseProgressRequest *CourseProgressRequest) error {
	client, clientErr := util.GetManagementClient()
	if clientErr != nil {
		return clientErr
	}
	user, getUserErr := findAuth0User(courseProgressRequest.UserId)
	if getUserErr != nil {
		return getUserErr
	}
	metadata, getMetadataErr := getUserAppMetadata(courseProgressRequest.UserId)
	if getMetadataErr != nil {
		return getMetadataErr
	}

	chapterdata := Chapterdata{}
	chapterdata.ChapterId = courseProgressRequest.ChapterId
	chapterdata.SubchapterId = courseProgressRequest.SubchapterId
	metadata.CoursesProgress[courseProgressRequest.CourseId] = chapterdata

	newUser := management.User{}
	metadataMap := make(map[string]interface{})
	metadataMap["coursesProgress"] = &metadata.CoursesProgress
	newUser.AppMetadata = &metadataMap

	err := client.User.Update(context.TODO(), user.GetID(), &newUser)
	if err != nil {
		println(err.Error())
	}

	return nil
}

func getUserAppMetadata(userId string) (AppMetadata, error) {
	user, getUserErr := findAuth0User(userId)
	if getUserErr != nil {
		return AppMetadata{}, getUserErr
	}
	metadata := AppMetadata{}
	mapstructure.Decode(user.GetAppMetadata(), &metadata)

	return metadata, nil
}

func findAuth0User(userId string) (*management.User, error) {
	url := "https://dev-l726rl1d8x1rw7du.eu.auth0.com/api/v2/users/auth0%7C" + userId
	method := "GET"
	token, err := util.GetAccesToken()
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	user := management.User{}
	json.Unmarshal(body, &user)

	return &user, nil
}
