package data

import (
	"context"

	courseService "darshub.dev/src/CourseService/data"
	"darshub.dev/src/UserService/config"
	"go.mongodb.org/mongo-driver/bson"
)

func SearchCourse(searchTerm string) ([]courseService.Course, error) {
	var courses []courseService.Course

	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	col := client.Database("darshub").Collection("course")
	searchStage := bson.A{bson.D{{"$search", bson.D{{"index", "adn_autocomplete"}, {"autocomplete", bson.D{{"path", "name"}, {"query", searchTerm}}}}}}}
	//searchStage := bson.A{bson.D{{"$search", bson.D{{"index", "adn_autocomplete"}, {"text", bson.D{{"path", "name"}, {"query", searchTerm}}}}}}}
	//projectStage := bson.D{{Key: "$project", Value: bson.D{{Key: "name", Value: 1}, {Key: "_id", Value: 0}}}}
	//opts := options.Aggregate().SetMaxTime(5 * time.Second)
	cursor, err := col.Aggregate(ctx, searchStage)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &courses); err != nil {
		return nil, err
	}
	return courses, nil
}
