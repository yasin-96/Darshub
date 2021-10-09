package data

import (
	"encoding/json"
	"io"
	"log"

	"dev.azure.com/learn-website-orga/_git/learn-website/backend/src/config"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
}

func AddTestProduct(p *Product) (primitive.ObjectID, error) {
	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	p.ID = primitive.NewObjectID()

	res, err := client.Database("darshub").Collection("user").InsertOne(ctx, p)
	if err != nil {
		log.Printf("Could not save Product: %v", err)
	}
	oid := res.InsertedID.(primitive.ObjectID)
	return oid, nil

}
func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}
