package handlers

import (
	"log"
	"net/http"

	"dev.azure.com/learn-website-orga/_git/learn-website/backend/src/data"
)

type Users struct {
	l *log.Logger
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUserStruct(l *log.Logger) *Users {
	return &Users{l}
}

func (u *Users) RegisterUser(rw http.ResponseWriter, r *http.Request) {
	user := &data.UserRequest{}

	err := data.FromJSON(user, r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	data.Create(user)
	rw.WriteHeader(http.StatusCreated)
}

func (u *Users) Login(rw http.ResponseWriter, r *http.Request) {
	loginRequest := &LoginRequest{}
	err := data.FromJSON(loginRequest, r.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(loginRequest.Email)

	user := data.Find(loginRequest.Email)

	if data.CheckIfPasswordsMatch(user, loginRequest.Password) {
		rw.WriteHeader(http.StatusOK)
	}
	rw.WriteHeader(http.StatusUnauthorized)
}
