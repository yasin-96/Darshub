package data

import (
	"log"
	"time"

	"dev.azure.com/learn-website-orga/_git/learn-website/backend/src/UserService/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	gomail "gopkg.in/gomail.v2"
)

//Abstraction to MVC, this are the service functions

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
	IsActive   bool               `bson:is_active"`
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
	Bio        string             `json:"bio"`
}

type UpdateUserRequest struct {
	Password   string    `json:"password"`
	First_Name string    `json:"first_name"`
	Last_Name  string    `json:"last_name"`
	Birthday   time.Time `json:"birthday"`
	Avatar     string    `json:"avatar"`
	Email      string    `json:"email"`
	TelNr      string    `json:"telNr"`
	Company    string    `json:"company"`
	Occupation string    `json:"occupation"`
	School     string    `json:"school"`
	Subject    string    `json:"subject"`
	Country    string    `json:"country"`
}

type ForgotPasswordRequest struct {
	Email string `bson:"email"`
}

func Create(userRequest *UserRequest) {
	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	userRequest.ID = primitive.NewObjectID()

	_, err := client.Database("darshub").Collection("user").InsertOne(ctx, userRequest.toUser())
	if err != nil {
		log.Printf("Could not save Product: %v", err)
	}
}

func Find(email string) User {
	var user User
	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	err := client.Database("darshub").Collection("user").FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	return user
}

func FindById(userId primitive.ObjectID) User {
	var user User
	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	err := client.Database("darshub").Collection("user").FindOne(ctx, bson.M{"_id": userId}).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}

	return user
}

func UpdateUser(userId primitive.ObjectID, updatedUser *UpdateUserRequest) User {
	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	update := bson.M{
		"password":   updatedUser.Password,
		"first_name": updatedUser.First_Name,
		"last_name":  updatedUser.Last_Name,
		"birthday":   updatedUser.Birthday,
		"avatar":     updatedUser.Avatar,
		"email":      updatedUser.Email,
		"tel_nr":     updatedUser.TelNr,
		"company":    updatedUser.Company,
		"occupation": updatedUser.Occupation,
		"school":     updatedUser.School,
		"subject":    updatedUser.Subject,
		"country":    updatedUser.Country,
	}

	_, err := client.Database("darshub").Collection("course").ReplaceOne(ctx, bson.M{"_id": userId}, update)
	if err != nil {
		log.Fatal(err)
	}

	return FindById(userId)
}

func DeleteUser(userId primitive.ObjectID) {
	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	_, err := client.Database("darshub").Collection("course").DeleteOne(ctx, bson.M{"_id": userId})
	if err != nil {
		log.Fatal(err)
	}
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

func ForgotPassword(email string) {
	msg := gomail.NewMessage()
	msg.SetHeader("From", "<paste your gmail account here>") //darshub email
	msg.SetHeader("To", email)
	msg.SetHeader("Subject", "Forgot password request")
	msg.SetBody("text/html", "<b>Forgot password request</b><br><p>Click on this linkt to reset your password: <a>link to password reset</a></p>")
	n := gomail.NewDialer("email host", 587, "darshub email", "email password")
	if err := n.DialAndSend(msg); err != nil {
		log.Fatal(err)
	}
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
