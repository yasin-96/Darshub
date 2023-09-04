package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"darshub.dev/src/UserService/config"
	"darshub.dev/src/util"
	"github.com/mitchellh/mapstructure"
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

type CourseRegisterRequest struct {
	CourseId string `json:"courseId"`
}

type Auth0User struct {
	CreatedAt     string `mapstructure:"created_at"`
	Email         string
	EmailVerified bool `mapstructure:"email_verified"`
	Name          string
	Nickname      string
	Picture       string
	UpdatedAt     string            `mapstructure:"updated_at"`
	UserId        string            `mapstructure:"user_id"`
	UserMetadata  map[string]string `mapstructure:"user_metadata"`
	Username      string
	AppMetadata   map[string][]string `mapstructure:"app_metadata"`
	LastLogin     string              `mapstructure:"last_login"`
}

var token string = ""

func Create(userRequest *UserRequest) error {
	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	user := userRequest.toUser()
	user.ID = primitive.NewObjectID()
	_, err := client.Database("darshub").Collection("user").InsertOne(ctx, user)
	return err
}

func GetAllUsers() ([]User, error) {
	var users []User
	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	cur, err := client.Database("darshub").Collection("user").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	cur.All(ctx, &users)

	return users, nil
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

func UpdateUser(userId primitive.ObjectID, updatedUser *UpdateUserRequest) (User, error) {

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
		return User{}, err
	}
	user := FindById(userId)
	return user, nil
}

func DeleteUser(userId primitive.ObjectID) error {
	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	_, err := client.Database("darshub").Collection("course").DeleteOne(ctx, bson.M{"_id": userId})
	return err
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

func RegisterToCourse(userId string, courseId string) error {
	tokenErr := checkIfTokenIsSet()
	if tokenErr != nil {
		return tokenErr
	}
	auth0Prefix := fmt.Sprintf("auth0|%s", userId)

	formattedUrl := fmt.Sprintf("https://dev-l726rl1d8x1rw7du.eu.auth0.com/api/v2/users/%s", auth0Prefix)
	method := "PATCH"
	registeredCourses, err := getRegisteredCourses(userId)
	if err != nil {
		return err
	}
	registeredCourses = append(registeredCourses, courseId)

	jsonArray, marshalErr := json.Marshal(registeredCourses)
	if marshalErr != nil {
		return marshalErr
	}

	payload := strings.NewReader(fmt.Sprintf(`{"app_metadata": {"courses": %s}}`, string(jsonArray)))

	client := &http.Client{}
	req, reqErr := http.NewRequest(method, formattedUrl, payload)

	if reqErr != nil {
		return reqErr
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	res, respErr := client.Do(req)
	if respErr != nil {
		return respErr
	}
	defer res.Body.Close()

	return nil
}

func checkIfTokenIsSet() error {
	if len(token) == 0 {
		retrievedToken, err := util.GetManagementAPIToken()
		if err != nil {
			return err
		}
		token = retrievedToken
	}
	return nil
}

func findAuth0User(userId string) (Auth0User, error) {
	tokenErr := checkIfTokenIsSet()
	if tokenErr != nil {
		return Auth0User{}, tokenErr
	}
	auth0Prefix := fmt.Sprintf("auth0|%s", userId)

	url := fmt.Sprintf("https://dev-l726rl1d8x1rw7du.eu.auth0.com/api/v2/users/%s", auth0Prefix)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return Auth0User{}, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return Auth0User{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return Auth0User{}, err
	}

	var result map[string]interface{}
	var user Auth0User
	json.Unmarshal([]byte(string(body)), &result)
	mapstructure.Decode(result, &user)
	return user, nil
}

func getRegisteredCourses(userId string) ([]string, error) {
	user, err := findAuth0User(userId)
	if err != nil {
		return nil, err
	}
	courses := user.AppMetadata["courses"]
	return courses, nil
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
