package data

import (
	"log"

	"darshub.dev/src/util"
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
	Name        string `json:"name"`
	Description string `json:"description"`
	Skills      string `json:"skills,omitempty"`
}

type UpdateCourseCategoryRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Skills      string `json:"skills"`
}

func CreateCourseCategory(courseCategory *CreateCourseCategoryRequest) error {
	ctx, cancel, client := util.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	_, err := client.Database("darshub").Collection("course_category").InsertOne(ctx, courseCategory)
	return err
}

func FindCourseCategory(courseCategoryId primitive.ObjectID) CourseCategory {
	var courseCategory CourseCategory
	ctx, cancel, client := util.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	err := client.Database("darshub").Collection("course_Category").FindOne(ctx, bson.M{"_id": courseCategoryId}).Decode(&courseCategory)
	if err != nil {
		return CourseCategory{}
	}
	return courseCategory
}

func FindAllCourses() []CourseCategory {
	var courseCategories []CourseCategory
	ctx, cancel, client := util.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	cur, err := client.Database("darshub").Collection("course_category").Find(ctx, bson.M{})
	if err != nil {
		log.Print(err)
	}
	cur.All(ctx, &courseCategories)
	return courseCategories
}

func UpdateCourseCategory(courseCategoryId primitive.ObjectID, updatedCourseCategory *UpdateCourseCategoryRequest) CourseCategory {
	ctx, cancel, client := util.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	update := bson.M{
		"name":        updatedCourseCategory.Name,
		"description": updatedCourseCategory.Description,
		"skills":      updatedCourseCategory.Skills,
	}

	_, err := client.Database("darshub").Collection("course_category").ReplaceOne(ctx, bson.M{"_id": courseCategoryId}, update)
	if err != nil {
		log.Print(err)
	}

	return FindCourseCategory(courseCategoryId)
}

func DeleteCourseCategory(courseCategoryId primitive.ObjectID) error {
	ctx, cancel, client := util.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	_, err := client.Database("darshub").Collection("course_category").DeleteOne(ctx, bson.M{"_id": courseCategoryId})
	if err != nil {
		return err
	}
	return nil
}
