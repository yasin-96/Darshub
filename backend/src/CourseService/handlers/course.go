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
	courseId, err := primitive.ObjectIDFromHex(vars["courseId"])
	if err != nil {
		log.Fatal(err)
	}
	course := data.Find(courseId)
	if reflect.ValueOf(course).IsZero() {
		rw.WriteHeader(http.StatusBadRequest)
	}
	rw.WriteHeader(http.StatusOK)
	parseErr := util.ToJSON(course, rw)
	if parseErr != nil {
		log.Fatal(parseErr)
	}
}

func UpdateCourse(rw http.ResponseWriter, r *http.Request) {
	updatedCourse := &data.UpdateCourseRequest{}
	vars := mux.Vars(r)
	courseId, err := primitive.ObjectIDFromHex(vars["courseId"])
	if err != nil {
		log.Fatal(err)
	}

	parseErr := util.FromJSON(updatedCourse, r.Body)
	if parseErr != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	course := data.Update(courseId, updatedCourse)
	rw.WriteHeader(http.StatusOK)
	util.ToJSON(course, rw)
}

func deleteCourse(rw http.ResponseWriter, r *http.Request) {

}
