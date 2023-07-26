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

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

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

	/*if user.Password == "" ||
		user.First_Name == "" ||
		user.Last_Name == "" ||
		&user.Birthday == new(time.Time) ||
		user.Email == "" ||
		user.TelNr == "" {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Request is not valid. Required informations are missing"))
		return
	}*/

	respErr := data.Create(user)
	if respErr != nil {
		http.Error(rw, respErr.Error(), http.StatusInternalServerError)
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
		http.Error(rw, "Np user id provided in path variable", http.StatusBadRequest)
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

func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if vars == nil || vars["userId"] == "" {
		http.Error(rw, "There was no user is provided in the path variable", http.StatusBadRequest)
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

func Login(rw http.ResponseWriter, r *http.Request) {

	if r.Body == http.NoBody {
		log.Println("Request is not valid. Req Body is not valid or missing information")
		http.Error(rw, "Request is not valid.", http.StatusBadRequest)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	loginRequest := &LoginRequest{}

	err := util.FromJSON(loginRequest, r.Body)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	if loginRequest.Email == "" || loginRequest.Password == "" {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	user := data.Find(loginRequest.Email)

	if reflect.ValueOf(user).IsZero() {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte("No account found with the entered email address."))
		return
	}

	if !data.CheckIfPasswordsMatch(user, loginRequest.Password) {
		rw.WriteHeader(http.StatusUnauthorized)
		rw.Write([]byte("The password which was entered is incorrect."))
		return
	}

	rw.WriteHeader(http.StatusOK)
	util.ToJSON(user, rw)
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
