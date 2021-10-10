package data

import (
	"encoding/json"
	"io"
	"log"
	"time"

	"dev.azure.com/learn-website-orga/_git/learn-website/backend/src/config"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID         primitive.ObjectID `bson:"_id"`
	Password   string             `bson:"password"`
	First_Name string             `bson:"name"`
	Last_Name  string             `bson:"description"`
	Birthday   time.Time          `bson:"birthday"`
	Avatar     string             `bson:"avatar"`
	Email      string             `bson:"email"`
	TelNr      string             `bson:"telNr"`
	Company    string             `bson:"company"`
	Occupation string             `bson:"occupation"`
	School     string             `bson:"school"`
	Subject    string             `bson:"subject"`
	Country    string             `bson:"country"`
}

func AddTestProduct(u *User) (primitive.ObjectID, error) {
	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	u.ID = primitive.NewObjectID()
	hashedPw := getEncryptedPassword([]byte(u.Password))
	hashedStringPassword := toString(hashedPw)
	u.Password = hashedStringPassword

	res, err := client.Database("darshub").Collection("user").InsertOne(ctx, u)
	if err != nil {
		log.Printf("Could not save Product: %v", err)
	}
	oid := res.InsertedID.(primitive.ObjectID)
	return oid, nil
}
func (u *User) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(u)
}

func getEncryptedPassword(password []byte) []byte {
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}

	return hashedPassword
}

func toString(bytes []byte) string {
	return string(bytes)
}
