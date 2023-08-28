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

func InsertChapter(rw http.ResponseWriter, r *http.Request) {
	chapterRequest := &data.CreateChapterRequest{}

	if r.Body == nil {
		http.Error(rw, "Body is empty", http.StatusBadRequest)
		return
	}

	err := util.FromJSON(chapterRequest, r.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	respErr := data.CreateChapter(chapterRequest)
	if respErr != nil {
		http.Error(rw, respErr.Error(), http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusCreated)
	log.Print("Chapter was created successfully")
}

func FindChapter(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if vars == nil || vars["chapterId"] == "" {
		http.Error(rw, "No chapter id provided in path variable", http.StatusBadRequest)
		return
	}

	chapterId, err := primitive.ObjectIDFromHex(vars["chapterId"])
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	chapter, respErr := data.FindChapter(chapterId)
	if respErr != nil {
		http.Error(rw, respErr.Error(), http.StatusInternalServerError)
		return
	}
	if reflect.ValueOf(chapter).IsZero() {
		http.Error(rw, "Chapter with the given id was not found", http.StatusNotFound)
		return
	}
	parseErr := util.ToJSON(chapter, rw)
	if parseErr != nil {
		http.Error(rw, parseErr.Error(), http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusOK)
}

func UpdateChapter(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if vars == nil || vars["chapterId"] == "" {
		http.Error(rw, "No chapter id provided in path variable", http.StatusBadRequest)
		return
	}

	updatedChapter := &data.UpdateChapterRequest{}

	chapterId, err := primitive.ObjectIDFromHex(vars["chapterId"])
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	parseErr := util.FromJSON(updatedChapter, r.Body)
	if parseErr != nil {
		http.Error(rw, parseErr.Error(), http.StatusBadRequest)
		return
	}

	chapter, err := data.UpdateChapter(chapterId, updatedChapter)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	util.ToJSON(chapter, rw)
	rw.WriteHeader(http.StatusOK)
}

func DeleteChapter(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if vars == nil || vars["chapterId"] == "" {
		http.Error(rw, "No chapter id provided in path variable", http.StatusBadRequest)
		return
	}

	chapterId, err := primitive.ObjectIDFromHex(vars["chapterId"])
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	respErr := data.DeleteChapter(chapterId)
	if respErr != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusNoContent)
}
