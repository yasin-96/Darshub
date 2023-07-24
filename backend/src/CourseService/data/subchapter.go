package data

import (
	"darshub.dev/src/UserService/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Subchapter struct {
	ID      primitive.ObjectID `bson:"_id"`
	Content []byte            `bson:"content"`
	Listing []string           `bson:"listing"`
}

type CreateSubchapterRequest struct {
	ID      primitive.ObjectID `json:"id"`
	Content []byte             `json:"content"`
	Listing []string           `json:"listing"`
}

type UpdateSubchapterRequest struct {
	Content []byte   `bson:"content"`
	Listing []string `bson:"listing"`
}

func CreateSubchapter(subchapter *CreateSubchapterRequest) error {
	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	subchapter.ID = primitive.NewObjectID()
	_, err := client.Database("darshub").Collection("subchapter").InsertOne(ctx, subchapter)
	return err
}

func FindSubchapter(subchapterId primitive.ObjectID) (Subchapter, error) {
	var subchapter Subchapter
	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	err := client.Database("darshub").Collection("subchapter").FindOne(ctx, bson.M{"_id": subchapterId}).Decode(&subchapter)
	if err != nil {
		return Subchapter{}, err
	}
	return subchapter, nil
}

func UpdateSubchapter(subchapterId primitive.ObjectID, updatedSubchapter *UpdateSubchapterRequest) (Subchapter, error) {
	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	update := bson.M{
		"content": updatedSubchapter.Content,
		"listing": updatedSubchapter.Listing,
	}

	_, err := client.Database("darshub").Collection("subchapter").ReplaceOne(ctx, bson.M{"_id": subchapterId}, update)
	if err != nil {
		return Subchapter{}, err
	}
	chapter, err := FindSubchapter(subchapterId)
	if err != nil {
		return Subchapter{}, err
	}
	return chapter, nil
}

func DeleteSubchapter(subchapterId primitive.ObjectID) error {
	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	_, err := client.Database("darshub").Collection("subchapter").DeleteOne(ctx, bson.M{"_id": subchapterId})
	return err
}
