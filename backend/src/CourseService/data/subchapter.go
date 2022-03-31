package data

import (
	"log"

	"dev.azure.com/learn-website-orga/_git/learn-website/backend/src/UserService/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Subchapter struct {
	ID      primitive.ObjectID `bson:"_id"`
	Content string             `bson:"content"`
	Listing []string           `bson:"listing"`
}

type CreateSubchapterRequest struct {
	ID      primitive.ObjectID `json:"id"`
	Content string             `json:"content"`
	Listing []string           `json:"listing"`
}

type UpdateSubchapterRequest struct {
	Content string   `bson:"content"`
	Listing []string `bson:"listing"`
}

func CreateSubchapter(subchapter *CreateSubchapterRequest) {
	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	subchapter.ID = primitive.NewObjectID()
	_, err := client.Database("darshub").Collection("chapter").InsertOne(ctx, subchapter)
	if err != nil {
		log.Printf("Could not save chapter: %v", err)
	}
}

func FindSubchapter(subchapterId primitive.ObjectID) Subchapter {
	var subchapter Subchapter
	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	err := client.Database("darshub").Collection("subchapter").FindOne(ctx, bson.M{"_id": subchapterId}).Decode(&subchapter)
	if err != nil {
		log.Fatal(err)
	}
	return subchapter
}

func UpdateSubchapter(subchapterId primitive.ObjectID, updatedSubchapter *UpdateSubchapterRequest) Chapter {
	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	update := bson.M{
		"content": updatedSubchapter.Content,
		"listing": updatedSubchapter.Listing,
	}

	_, err := client.Database("darshub").Collection("subchapter").ReplaceOne(ctx, bson.M{"_id": subchapterId}, update)
	if err != nil {
		log.Fatal(err)
	}

	return FindChapter(subchapterId)
}

func DeleteSubchapter(subchapterId primitive.ObjectID) {
	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	_, err := client.Database("darshub").Collection("subchapter").DeleteOne(ctx, bson.M{"_id": subchapterId})
	if err != nil {
		log.Fatal(err)
	}
}
