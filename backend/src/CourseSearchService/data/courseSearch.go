package data

import (
	"context"
	"fmt"
	"log"

	courseService "dev.azure.com/learn-website-orga/_git/learn-website/src/CourseService/data"
	"dev.azure.com/learn-website-orga/_git/learn-website/src/UserService/config"
	"go.mongodb.org/mongo-driver/bson"
)

func SearchCourse(searchTerm string) []courseService.Course {
	var courses []courseService.Course

	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	col := client.Database("darshub").Collection("course")
	searchStage := bson.A{bson.D{{"$search", bson.D{{"index", "name"}, {"text", bson.D{{"path", "name"}, {"query", searchTerm}}}}}}}
	//projectStage := bson.D{{Key: "$project", Value: bson.D{{Key: "name", Value: 1}, {Key: "_id", Value: 0}}}}
	//opts := options.Aggregate().SetMaxTime(5 * time.Second)
	cursor, err := col.Aggregate(ctx, searchStage)
	if err != nil {
		print("atlas search failed")
		log.Printf("test %s", err)
	}
	if err = cursor.All(context.TODO(), &courses); err != nil {
	}
	println(len(courses))
	//println("testo", courses)
	for _, result := range courses {
		fmt.Println("out", result.Name)
	}
	return courses
}
