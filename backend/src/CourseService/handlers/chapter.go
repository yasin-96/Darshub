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

func InsertChapter(rw http.ResponseWriter, r *http.Request) {
	chapter := &data.CreateChapterRequest{}

	err := util.FromJSON(chapter, r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	data.CreateChapter(chapter)
	rw.WriteHeader(http.StatusCreated)
}

func FindChapter(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	chapterId, err := primitive.ObjectIDFromHex(vars["courseId"])
	if err != nil {
		log.Fatal(err)
	}
	chapter := data.FindChapter(chapterId)
	if reflect.ValueOf(chapter).IsZero() {
		rw.WriteHeader(http.StatusNotFound)
	}
	rw.WriteHeader(http.StatusOK)
	parseErr := util.ToJSON(chapter, rw)
	if parseErr != nil {
		log.Fatal(parseErr)
	}
}

func UpdateChapter(rw http.ResponseWriter, r *http.Request) {
	updatedChapter := &data.UpdateChapterRequest{}
	vars := mux.Vars(r)
	chapterId, err := primitive.ObjectIDFromHex(vars["chapterId"])
	if err != nil {
		log.Fatal(err)
	}

	parseErr := util.FromJSON(updatedChapter, r.Body)
	if parseErr != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	chapter := data.UpdateChapter(chapterId, updatedChapter)
	rw.WriteHeader(http.StatusOK)
	util.ToJSON(chapter, rw)
}

func DeleteChapter(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	chapterId, err := primitive.ObjectIDFromHex(vars["chapterId"])
	if err != nil {
		log.Fatal(err)
	}

	data.DeleteChapter(chapterId)
	rw.WriteHeader(http.StatusNoContent)
}
