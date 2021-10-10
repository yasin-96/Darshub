package handlers

import (
	"log"
	"net/http"

	"dev.azure.com/learn-website-orga/_git/learn-website/backend/src/data"
)

type Users struct {
	l *log.Logger
}

func NewUserStruct(l *log.Logger) *Users {
	return &Users{l}
}

func (u *Users) AddTestProduct(rw http.ResponseWriter, r *http.Request) {
	prod := &data.User{}

	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal", http.StatusBadRequest)
	}
	data.AddTestProduct(prod)
}
