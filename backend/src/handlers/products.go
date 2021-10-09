package handlers

import (
	"log"
	"net/http"

	"dev.azure.com/learn-website-orga/_git/learn-website/backend/src/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) AddTestProduct(rw http.ResponseWriter, r *http.Request) {
	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal", http.StatusBadRequest)
	}

	data.AddTestProduct(prod)

}
