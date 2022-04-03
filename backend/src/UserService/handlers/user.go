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
	if r.Body == nil {
		log.Println("Request is not valid.")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	user := &data.UserRequest{}

	err := util.FromJSON(user, r.Body)
	if err != nil {
		log.Println("Request body could not parsed")
		log.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	dbResponse := data.Create(user)
	rw.WriteHeader(http.StatusCreated)
	rw.Write([]byte(dbResponse))
}

func FindById(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if vars == nil || vars["userId"] == "" {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	userId, err := primitive.ObjectIDFromHex(vars["userId"])
	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	user := data.FindById(userId)
	if reflect.ValueOf(user).IsZero() {
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	rw.WriteHeader(http.StatusOK)
	parseErr := util.ToJSON(user, rw)
	if parseErr != nil {
		log.Fatal(err)
	}
}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if vars == nil || vars["userId"] == "" {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	updatedUser := &data.UpdateUserRequest{}
	userId, err := primitive.ObjectIDFromHex(vars["userId"])

	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	parseErr := util.FromJSON(updatedUser, r.Body)
	if parseErr != nil {
		// http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
		log.Println("Request body could not parsed")
		log.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	changedUser := data.UpdateUser(userId, updatedUser)

	// if course == nil {
	// 	rw.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }

	rw.WriteHeader(http.StatusOK)
	util.ToJSON(changedUser, rw)
}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if vars == nil || vars["userId"] == "" {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	userId, err := primitive.ObjectIDFromHex(vars["userId"])
	if err != nil {
		log.Println("Request body could not parsed")
		log.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	data.DeleteUser(userId)
	rw.WriteHeader(http.StatusNoContent)
}

func Login(rw http.ResponseWriter, r *http.Request) {

	if r.Body == nil {
		log.Println("Request is not valid.")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

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
