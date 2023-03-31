package data

import (
	"log"
	"time"

	"dev.azure.com/learn-website-orga/_git/learn-website/src/UserService/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

//Abstraction to MVC, this are the service functions

type User struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	Password   []byte             `json:"-" bson:"password"` // will not in the response
	First_Name string             `json:"first_name" bson:"first_name"`
	Last_Name  string             `json:"last_name" bson:"last_name"`
	Birthday   time.Time          `json:"birthday" bson:"birthday"`
	Avatar     string             `json:"avatar" bson:"avatar"`
	Email      string             `json:"email" bson:"email"`
	TelNr      string             `json:"telNr" bson:"telNr"`
	Company    string             `json:"company" bson:"company"`
	Occupation string             `json:"occupation" bson:"occupation"`
	School     string             `json:"school" bson:"school"`
	Subject    string             `json:"subject" bson:"subject"`
	Country    string             `json:"country" bson:"country"`
	IsActive   bool               `json:"isActive" bson:"isActive"`
	Role       []int              `json:"role" bson:"role"`
}

type UserRequest struct {
	Password   string    `json:"password"`
	First_Name string    `json:"first_name"`
	Last_Name  string    `json:"last_name"`
	Birthday   time.Time `json:"birthday"`
	Avatar     string    `json:"avatar"`
	Email      string    `json:"email"`
	TelNr      string    `json:"tel_nr"`
	Company    string    `json:"company"`
	Occupation string    `json:"occupation"`
	School     string    `json:"school"`
	Subject    string    `json:"subject"`
	Country    string    `json:"country"`
	Role       []int     `bson:"role"`
}

type UpdateUserRequest struct {
	First_Name string    `json:"first_name"`
	Last_Name  string    `json:"last_name"`
	Birthday   time.Time `json:"birthday"`
	Avatar     string    `json:"avatar"`
	Email      string    `json:"email"`
	TelNr      string    `json:"tel_nr"`
	Company    string    `json:"company"`
	Occupation string    `json:"occupation"`
	School     string    `json:"school"`
	Subject    string    `json:"subject"`
	Country    string    `json:"country"`
}

func Create(userRequest *UserRequest) {
	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	user := userRequest.toUser()
	user.ID = primitive.NewObjectID()
	_, err := client.Database("darshub").Collection("user").InsertOne(ctx, user)
	if err != nil {
		log.Printf("Could not save the new user obj: %v\n", err)
		return
	}

	log.Println("User was created")
}

func GetAllUsers() []User {
	var users []User
	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	cur, err := client.Database("darshub").Collection("user").Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	cur.All(ctx, &users)

	return users
}

func Find(email string) User {
	if email == "" {
		return User{}
	}

	user := getUser(email)

	return user
}

func FindById(userId primitive.ObjectID) User {
	var user User
	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	err := client.Database("darshub").Collection("user").FindOne(ctx, bson.M{"_id": userId}).Decode(&user)
	if err != nil {
		log.Println(err)
		return User{}
	}
	return user
}

func UpdateUser(userId primitive.ObjectID, updatedUser *UpdateUserRequest) User {

	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	update := bson.M{
		"first_name": updatedUser.First_Name,
		"last_name":  updatedUser.Last_Name,
		"birthday":   updatedUser.Birthday,
		"avatar":     updatedUser.Avatar,
		"email":      updatedUser.Email,
		"company":    updatedUser.Company,
		"occupation": updatedUser.Occupation,
		"school":     updatedUser.School,
		"subject":    updatedUser.Subject,
		"country":    updatedUser.Country,
		"tel_nr":     updatedUser.TelNr,
	}

	_, err := client.Database("darshub").Collection("user").ReplaceOne(ctx, bson.M{"_id": userId}, update)
	if err != nil {
		log.Println(err)
		return User{}
	}
	return FindById(userId)
}

func DeleteUser(userId primitive.ObjectID) {
	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	_, err := client.Database("darshub").Collection("course").DeleteOne(ctx, bson.M{"_id": userId})
	if err != nil {
		log.Print(err)
		return
	}
	log.Print("User was deleted successfully")
}

func CheckIfPasswordsMatch(user User, password string) bool {
	err := bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	return err == nil
}

func SetAccountInactive(userId primitive.ObjectID) {
	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	userToBeDeactivated := FindById(userId)
	userToBeDeactivated.IsActive = false

	_, err := client.Database("darshub").Collection("user").ReplaceOne(ctx, bson.M{"_id": userToBeDeactivated.ID}, userToBeDeactivated)
	if err != nil {
		log.Println(err)
	}
}

func getUser(email string) User {
	var user User
	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	err := client.Database("darshub").Collection("user").FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return user
	}

	return user
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
	user.IsActive = true
	user.Subject = userRequest.Subject
	user.Country = userRequest.Country
	return user
}
