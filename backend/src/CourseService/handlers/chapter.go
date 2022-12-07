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

func InsertChapter(rw http.ResponseWriter, r *http.Request) {
	chapterRequest := &data.CreateChapterRequest{}

	if r.Body == nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	err := util.FromJSON(chapterRequest, r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
		return
	}
	data.CreateChapter(chapterRequest)
	rw.WriteHeader(http.StatusCreated)

}

func FindChapter(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if vars == nil || vars["chapterId"] == "" {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	chapterId, err := primitive.ObjectIDFromHex(vars["chapterId"])
	if err != nil {
		log.Print(err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	chapter := data.FindChapter(chapterId)
	if reflect.ValueOf(chapter).IsZero() {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte("Course category with the given id does not exist"))
		return
	}
	parseErr := util.ToJSON(chapter, rw)
	if parseErr != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		log.Print(parseErr)
	}
}

func UpdateChapter(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if vars == nil || vars["chapterId"] == "" {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	updatedChapter := &data.UpdateChapterRequest{}

	chapterId, err := primitive.ObjectIDFromHex(vars["chapterId"])
	if err != nil {
		log.Print(err)
		return
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

	if vars == nil || vars["chapterId"] == "" {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	chapterId, err := primitive.ObjectIDFromHex(vars["chapterId"])
	if err != nil {
		log.Print(err)
		return
	}
	data.DeleteChapter(chapterId)
	rw.WriteHeader(http.StatusNoContent)
}
