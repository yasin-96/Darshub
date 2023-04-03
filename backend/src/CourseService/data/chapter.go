package data

import (
	"log"

	"darshub.dev/src/UserService/config"
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
	Name        string `json:"name"`
	Description string `json:"description"`
	Skills      string `json:"skills"`
}

type UpdateChapterRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Skills      string `json:"skills"`
}

func CreateChapter(chapterRequest *CreateChapterRequest) {
	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	_, err := client.Database("darshub").Collection("chapter").InsertOne(ctx, chapterRequest)
	if err != nil {
		log.Printf("Could not save chapter: %v", err)
	}
}

func FindChapter(chapterId primitive.ObjectID) Chapter {
	var chapter Chapter
	log.Printf(chapterId.String())
	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	err := client.Database("darshub").Collection("chapter").FindOne(ctx, bson.M{"_id": chapterId}).Decode(&chapter)
	if err != nil {
		log.Print(err)

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
		log.Print(err)
		return Chapter{}
	}

	return FindChapter(chapterId)
}

func DeleteChapter(chapterId primitive.ObjectID) {
	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	_, err := client.Database("darshub").Collection("chapter").DeleteOne(ctx, bson.M{"_id": chapterId})
	if err != nil {
		log.Print(err)
		return
	}
	log.Print("Chapter was deleted successfully")
}

func (chapterRequest *CreateChapterRequest) toChapter() Chapter {
	chapter := Chapter{}
	chapter.Description = chapterRequest.Description
	chapter.Name = chapterRequest.Name
	chapter.Skills = chapterRequest.Skills
	return chapter
}
