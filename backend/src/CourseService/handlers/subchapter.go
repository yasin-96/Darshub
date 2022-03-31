package handlers

import (
	"log"
	"net/http"
	"reflect"

	"dev.azure.com/learn-website-orga/_git/learn-website/backend/src/CourseService/data"
	"dev.azure.com/learn-website-orga/_git/learn-website/backend/src/util"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertSubchapter(rw http.ResponseWriter, r *http.Request) {
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
	subchapterId, err := primitive.ObjectIDFromHex(vars["courseId"])
	if err != nil {
		log.Fatal(err)
	}
	subchapter := data.FindSubchapter(subchapterId)
	if reflect.ValueOf(subchapter).IsZero() {
		rw.WriteHeader(http.StatusNotFound)
	}
	rw.WriteHeader(http.StatusOK)
	parseErr := util.ToJSON(subchapter, rw)
	if parseErr != nil {
		log.Fatal(parseErr)
	}
}

func UpdateSubchapter(rw http.ResponseWriter, r *http.Request) {
	updatedSubchapter := &data.UpdateSubchapterRequest{}
	vars := mux.Vars(r)
	chapterId, err := primitive.ObjectIDFromHex(vars["subchapterId"])
	if err != nil {
		log.Fatal(err)
	}

	parseErr := util.FromJSON(updatedSubchapter, r.Body)
	if parseErr != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	subchapter := data.UpdateSubchapter(chapterId, updatedSubchapter)
	rw.WriteHeader(http.StatusOK)
	util.ToJSON(subchapter, rw)
}

func DeleteSubchapter(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	subchapterId, err := primitive.ObjectIDFromHex(vars["subchapterId"])
	if err != nil {
		log.Fatal(err)
	}

	data.DeleteSubchapter(subchapterId)
	rw.WriteHeader(http.StatusNoContent)
}
