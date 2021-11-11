package handlers

import (
	"log"
	"net/http"

	"dev.azure.com/learn-website-orga/_git/learn-website/backend/src/CourseService/data"
	"dev.azure.com/learn-website-orga/_git/learn-website/backend/src/util"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertCourse(rw http.ResponseWriter, r *http.Request) {
	course := &data.Course{}

	err := util.FromJSON(course, r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	data.Create(course)
	rw.WriteHeader(http.StatusCreated)
}

func FindCourse(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := primitive.ObjectIDFromHex(vars["courseId"])
	if err != nil {
		log.Fatal(err)
	}
	course := data.Find(userId)
	rw.WriteHeader(http.StatusOK)
	parseErr := util.ToJSON(course, rw)
	if parseErr != nil {
		log.Fatal(parseErr)
	}
}
