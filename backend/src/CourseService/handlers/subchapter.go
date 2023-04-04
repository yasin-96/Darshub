package handlers

import (
	"log"
	"net/http"
	"reflect"

	"darshub.dev/src/CourseService/data"
	"darshub.dev/src/util"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertSubchapter(rw http.ResponseWriter, r *http.Request) {

	if r.Body == nil {
		http.Error(rw, "Body is empty", http.StatusBadRequest)
		return
	}

	subchapter := &data.CreateSubchapterRequest{}

	err := util.FromJSON(subchapter, r.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	respErr := data.CreateSubchapter(subchapter)
	if respErr != nil {
		http.Error(rw, respErr.Error(), http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusCreated)
	log.Print("Subchapter was created successfully")
}

func FindSubchapter(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if vars == nil || vars["subchapterId"] == "" {
		http.Error(rw, "Subchapter id was not provided in the path variable", http.StatusBadRequest)
		return
	}

	subchapterId, err := primitive.ObjectIDFromHex(vars["subchapterId"])
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	subchapter, err := data.FindSubchapter(subchapterId)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	if reflect.ValueOf(subchapter).IsZero() {
		http.Error(rw, "The sbchapter with the given id was not found", http.StatusNotFound)
		return
	}
	parseErr := util.ToJSON(subchapter, rw)
	if parseErr != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusOK)
}

func UpdateSubchapter(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if vars == nil || vars["subchapterId"] == "" {
		http.Error(rw, "There was no subchapter id provided in the path variable", http.StatusBadRequest)
		return
	}

	updatedSubchapter := &data.UpdateSubchapterRequest{}
	chapterId, err := primitive.ObjectIDFromHex(vars["subchapterId"])
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	parseErr := util.FromJSON(updatedSubchapter, r.Body)
	if parseErr != nil {
		http.Error(rw, parseErr.Error(), http.StatusBadRequest)
		return
	}

	subchapter, err := data.UpdateSubchapter(chapterId, updatedSubchapter)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	if reflect.ValueOf(subchapter).IsZero() {
		http.Error(rw, "The subchapter with the given id was not found", http.StatusNotFound)
		return
	}
	util.ToJSON(subchapter, rw)
	rw.WriteHeader(http.StatusOK)
}

func DeleteSubchapter(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if vars == nil || vars["subchapterId"] == "" {
		http.Error(rw, "There was no subchapter id provided in the path variable", http.StatusBadRequest)
		return
	}

	subchapterId, err := primitive.ObjectIDFromHex(vars["subchapterId"])
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	respErr := data.DeleteSubchapter(subchapterId)
	if respErr != nil {
		http.Error(rw, respErr.Error(), http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusNoContent)
}
