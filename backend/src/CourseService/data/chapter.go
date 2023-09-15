package data

import (
	"darshub.dev/src/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Chapter struct {
	ID          primitive.ObjectID   `bson:"_id"`
	Name        string               `bson:"name"`
	Description string               `bson:"description"`
	Skills      string               `bson:"skills"`
	Subchapters []primitive.ObjectID `bson:"subchapters"`
}

type CreateChapterRequest struct {
	Name        string               `json:"name"`
	Description string               `json:"description"`
	Skills      string               `json:"skills"`
	Subchapters []primitive.ObjectID `json:"subchapters"`
}

type UpdateChapterRequest struct {
	Name        string               `json:"name"`
	Description string               `json:"description"`
	Skills      string               `json:"skills"`
	Subchapters []primitive.ObjectID `json:"subchapters"`
}

func CreateChapter(chapterRequest *CreateChapterRequest) error {
	ctx, cancel, client := util.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	_, err := client.Database("darshub").Collection("chapter").InsertOne(ctx, chapterRequest)
	if err != nil {
		return err
	}
	return nil
}

func FindChapter(chapterId primitive.ObjectID) (Chapter, error) {
	var chapter Chapter
	ctx, cancel, client := util.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	err := client.Database("darshub").Collection("chapter").FindOne(ctx, bson.M{"_id": chapterId}).Decode(&chapter)
	if err != nil {
		return Chapter{}, err
	}
	return chapter, nil
}

func UpdateChapter(chapterId primitive.ObjectID, updatedChapter *UpdateChapterRequest) (Chapter, error) {
	ctx, cancel, client := util.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	update := bson.M{
		"name":        updatedChapter.Name,
		"description": updatedChapter.Description,
		"skills":      updatedChapter.Skills,
		"subchapters": updatedChapter.Subchapters,
	}

	_, err := client.Database("darshub").Collection("chapter").ReplaceOne(ctx, bson.M{"_id": chapterId}, update)
	if err != nil {
		return Chapter{}, err
	}
	chapter, err := FindChapter(chapterId)
	if err != nil {
		return Chapter{}, err
	}
	return chapter, nil
}

func DeleteChapter(chapterId primitive.ObjectID) error {
	ctx, cancel, client := util.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	_, err := client.Database("darshub").Collection("chapter").DeleteOne(ctx, bson.M{"_id": chapterId})
	return err
}
