package handlers

import (
	"log"
	"net/http"
	"reflect"

	"darshub.dev/src/UserService/data"
	"darshub.dev/src/util"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RegisterUser(rw http.ResponseWriter, r *http.Request) {

	if r.Body == http.NoBody {
		http.Error(rw, "Request is not valid.", http.StatusBadRequest)
		return
	}

	user := &data.UserRequest{}

	err := util.FromJSON(user, r.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	respErr := data.Create(user)
	if respErr != nil {
		http.Error(rw, respErr.Error(), http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusCreated)
}

func RegisterUserToCourse(rw http.ResponseWriter, r *http.Request) {
	if r.Body == http.NoBody {
		http.Error(rw, "Request is not valid.", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)

	if vars == nil || vars["userId"] == "" {
		http.Error(rw, "Np user id provided in path variable", http.StatusBadRequest)
		return
	}

	userId := vars["userId"]
	request := &data.CourseRegisterRequest{}

	parseErr := util.FromJSON(request, r.Body)
	if parseErr != nil {
		http.Error(rw, parseErr.Error(), http.StatusBadRequest)
		return
	}

	err := data.RegisterToCourse(userId, request.CourseId)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusCreated)
}

func GetAllUsers(rw http.ResponseWriter, r *http.Request) {
	users, err := data.GetAllUsers()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	parseErr := util.ToJSON(users, rw)
	if parseErr != nil {
		log.Print(parseErr)
	}
}

func FindById(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if vars == nil || vars["userId"] == "" {
		http.Error(rw, "No user id provided in path variable", http.StatusBadRequest)
		return
	}

	userId, err := primitive.ObjectIDFromHex(vars["userId"])
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	user := data.FindById(userId)
	if reflect.ValueOf(user).IsZero() {
		http.Error(rw, "The user with the given id was not found", http.StatusNotFound)
		return
	}

	parseErr := util.ToJSON(user, rw)
	if parseErr != nil {
		http.Error(rw, parseErr.Error(), http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusOK)
}

func FindUserAuth0(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if vars == nil || vars["userId"] == "" {
		http.Error(rw, "There was no user id provided in the path variable", http.StatusBadRequest)
		return
	}

	userId := vars["userId"]

	user, err := data.FindUserAuth0(userId)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	parseErr := util.ToJSON(user, rw)
	if parseErr != nil {
		http.Error(rw, parseErr.Error(), http.StatusInternalServerError)
		return
	}
}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if vars == nil || vars["userId"] == "" {
		http.Error(rw, "There was no user id provided in the path variable", http.StatusBadRequest)
		return
	}

	updatedUser := &data.UpdateUserRequest{}
	userId, err := primitive.ObjectIDFromHex(vars["userId"])

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	parseErr := util.FromJSON(updatedUser, r.Body)
	if parseErr != nil {
		http.Error(rw, parseErr.Error(), http.StatusBadRequest)
		return
	}

	changedUser, respErr := data.UpdateUser(userId, updatedUser)
	if respErr != nil {
		http.Error(rw, respErr.Error(), http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusOK)
	util.ToJSON(changedUser, rw)
}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if vars == nil || vars["userId"] == "" {
		http.Error(rw, "No user id provided in the path variable", http.StatusBadRequest)
		return
	}

	userId, err := primitive.ObjectIDFromHex(vars["userId"])
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	respErr := data.DeleteUser(userId)
	if respErr != nil {
		http.Error(rw, respErr.Error(), http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusNoContent)
}

func SetAccountInactive(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if vars == nil || vars["userId"] == "" {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	userId, err := primitive.ObjectIDFromHex(vars["userId"])
	if err != nil {
		log.Print("Could not parse userId: ", err)
	}

	data.SetAccountInactive(userId)
}
