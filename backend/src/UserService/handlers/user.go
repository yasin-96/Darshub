package handlers

import (
	"log"
	"net/http"
	"reflect"

	"dev.azure.com/learn-website-orga/_git/learn-website/backend/src/UserService/data"
	"dev.azure.com/learn-website-orga/_git/learn-website/backend/src/util"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RegisterUser(rw http.ResponseWriter, r *http.Request) {
	user := &data.UserRequest{}

	err := util.FromJSON(user, r.Body)
	if err != nil {
		log.Fatal(err)
	}

	data.Create(user)
	rw.WriteHeader(http.StatusCreated)
}

func FindById(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := primitive.ObjectIDFromHex(vars["userId"])
	if err != nil {
		log.Fatal(err)
	}
	user := data.FindById(userId)
	if reflect.ValueOf(user).IsZero() {
		rw.WriteHeader(http.StatusNotFound)
	}
	rw.WriteHeader(http.StatusOK)
	parseErr := util.ToJSON(user, rw)
	if parseErr != nil {
		log.Fatal(err)
	}
}

func Login(rw http.ResponseWriter, r *http.Request) {
	loginRequest := &LoginRequest{}
	err := util.FromJSON(loginRequest, r.Body)
	if err != nil {
		log.Fatal(err)
	}

	user := data.Find(loginRequest.Email)

	if data.CheckIfPasswordsMatch(user, loginRequest.Password) {
		rw.WriteHeader(http.StatusOK)
	}
	rw.WriteHeader(http.StatusUnauthorized)
}
