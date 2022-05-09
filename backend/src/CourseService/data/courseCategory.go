package data

import (
	"log"

	"dev.azure.com/learn-website-orga/_git/learn-website/src/UserService/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CourseCategory struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	Skills      string             `bson:"skills"`
}

type CreateCourseCategoryRequest struct {
	ID          primitive.ObjectID `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Skills      string             `json:"skills"`
}

type UpdateCourseCategoryRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Skills      string `json:"skills"`
}

func CreateCourseCategory(courseCategory *CreateCourseCategoryRequest) {
	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	courseCategory.ID = primitive.NewObjectID()
	_, err := client.Database("darshub").Collection("courseCategory").InsertOne(ctx, courseCategory)
	if err != nil {
		log.Printf("Could not save course: %v", err)
	}
}

func FindCourseCategory(courseCategoryId primitive.ObjectID) CourseCategory {
	var courseCategory CourseCategory
	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	err := client.Database("darshub").Collection("course").FindOne(ctx, bson.M{"_id": courseCategoryId}).Decode(&courseCategory)
	if err != nil {
		log.Fatal(err)
	}
	return courseCategory
}

func UpdateCourseCategory(courseCategoryId primitive.ObjectID, updatedCourseCategory *UpdateCourseCategoryRequest) CourseCategory {
	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	update := bson.M{
		"name":        updatedCourseCategory.Name,
		"description": updatedCourseCategory.Description,
		"status":      updatedCourseCategory.Skills,
	}

	_, err := client.Database("darshub").Collection("course").ReplaceOne(ctx, bson.M{"_id": courseCategoryId}, update)
	if err != nil {
		log.Fatal(err)
	}

	return FindCourseCategory(courseCategoryId)
}

func DeleteCourseCategory(courseCategoryId primitive.ObjectID) {
	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	_, err := client.Database("darshub").Collection("courseCategory").DeleteOne(ctx, bson.M{"_id": courseCategoryId})
	if err != nil {
		log.Fatal(err)
	}
}
