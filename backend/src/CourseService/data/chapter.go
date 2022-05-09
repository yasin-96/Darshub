package data

import (
	"log"

	"dev.azure.com/learn-website-orga/_git/learn-website/src/UserService/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Chapter struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	Skills      string             `bson:"skills"`
}

type CreateChapterRequest struct {
	ID          primitive.ObjectID `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Skills      string             `json:"skills"`
}

type UpdateChapterRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Skills      string `json:"skills"`
}

func CreateChapter(chapter *CreateChapterRequest) {
	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	chapter.ID = primitive.NewObjectID()
	_, err := client.Database("darshub").Collection("chapter").InsertOne(ctx, chapter)
	if err != nil {
		log.Printf("Could not save chapter: %v", err)
	}
}

func FindChapter(chapterId primitive.ObjectID) Chapter {
	var chapter Chapter
	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	err := client.Database("darshub").Collection("course").FindOne(ctx, bson.M{"_id": chapterId}).Decode(&chapter)
	if err != nil {
		log.Fatal(err)
	}
	return chapter
}

func UpdateChapter(chapterId primitive.ObjectID, updatedChapter *UpdateChapterRequest) Chapter {
	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	update := bson.M{
		"name":        updatedChapter.Name,
		"description": updatedChapter.Description,
		"skills":      updatedChapter.Skills,
	}

	_, err := client.Database("darshub").Collection("chapter").ReplaceOne(ctx, bson.M{"_id": chapterId}, update)
	if err != nil {
		log.Fatal(err)
	}

	return FindChapter(chapterId)
}

func DeleteChapter(chapterId primitive.ObjectID) {
	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	_, err := client.Database("darshub").Collection("chapter").DeleteOne(ctx, bson.M{"_id": chapterId})
	if err != nil {
		log.Fatal(err)
	}
}
