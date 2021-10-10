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
	Password   []byte             `bson:"password"`
	First_Name string             `bson:"first_name"`
	Last_Name  string             `bson:"last_name"`
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

type UserRequest struct {
	ID         primitive.ObjectID `json:"id"`
	Password   string             `json:"password"`
	First_Name string             `json:"first_name"`
	Last_Name  string             `json:"last_name"`
	Birthday   time.Time          `json:"birthday"`
	Avatar     string             `json:"avatar"`
	Email      string             `json:"email"`
	TelNr      string             `json:"telNr"`
	Company    string             `json:"company"`
	Occupation string             `json:"occupation"`
	School     string             `json:"school"`
	Subject    string             `json:"subject"`
	Country    string             `json:"country"`
}

func AddTestProduct(userRequest *UserRequest) (primitive.ObjectID, error) {
	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	userRequest.ID = primitive.NewObjectID()

	res, err := client.Database("darshub").Collection("user").InsertOne(ctx, userRequest.toUser())
	if err != nil {
		log.Printf("Could not save Product: %v", err)
	}
	oid := res.InsertedID.(primitive.ObjectID)
	return oid, nil
}

func (userRequest *UserRequest) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(userRequest)
}

func getEncryptedPassword(password []byte) []byte {
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}

	return hashedPassword
}

func (userRequest *UserRequest) toUser() User {
	user := User{}
	user.ID = userRequest.ID
	user.Password = getEncryptedPassword([]byte(userRequest.Password))
	user.First_Name = userRequest.First_Name
	user.Last_Name = userRequest.Last_Name
	user.Birthday = userRequest.Birthday
	user.Avatar = userRequest.Avatar
	user.Email = userRequest.Email
	user.TelNr = userRequest.TelNr
	user.Company = userRequest.Company
	user.Occupation = userRequest.Occupation
	user.School = userRequest.School
	user.Subject = userRequest.Subject
	user.Country = userRequest.Country

	return user
}
