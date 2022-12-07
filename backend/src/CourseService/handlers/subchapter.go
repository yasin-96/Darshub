package handlers

import (
	"log"
	"net/http"
	"reflect"

	"dev.azure.com/learn-website-orga/_git/learn-website/src/CourseService/data"
	"dev.azure.com/learn-website-orga/_git/learn-website/src/util"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertSubchapter(rw http.ResponseWriter, r *http.Request) {

	if r.Body == nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	subchapter := &data.CreateSubchapterRequest{}

	err := util.FromJSON(subchapter, r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	data.CreateSubchapter(subchapter)
	rw.WriteHeader(http.StatusCreated)
}

func FindSubchapter(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if vars == nil || vars["subchapterId"] == "" {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	subchapterId, err := primitive.ObjectIDFromHex(vars["subchapterId"])
	if err != nil {
		log.Print(err)
		rw.WriteHeader(http.StatusInternalServerError)
	}
	subchapter := data.FindSubchapter(subchapterId)
	if reflect.ValueOf(subchapter).IsZero() {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte("The subchapter with the given id does not exist"))
	}
	parseErr := util.ToJSON(subchapter, rw)
	if parseErr != nil {
		log.Print(parseErr)
		rw.WriteHeader(http.StatusInternalServerError)
	}
}

func UpdateSubchapter(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if vars == nil || vars["subchapterId"] == "" {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	updatedSubchapter := &data.UpdateSubchapterRequest{}
	chapterId, err := primitive.ObjectIDFromHex(vars["subchapterId"])
	if err != nil {
		log.Print(err)
		return
	}
	parseErr := util.FromJSON(updatedSubchapter, r.Body)
	if parseErr != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	subchapter := data.UpdateSubchapter(chapterId, updatedSubchapter)
	if reflect.ValueOf(subchapter).IsZero() {
		rw.WriteHeader(http.StatusInternalServerError)
	}
	util.ToJSON(subchapter, rw)
}

func DeleteSubchapter(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if vars == nil || vars["subchapterId"] == "" {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	subchapterId, err := primitive.ObjectIDFromHex(vars["subchapterId"])
	if err != nil {
		log.Print(err)
		return
	}
	data.DeleteSubchapter(subchapterId)
	rw.WriteHeader(http.StatusNoContent)
}
