package handlers

import (
	"log"
	"net/http"
	"reflect"
	"time"

	"dev.azure.com/learn-website-orga/_git/learn-website/src/CourseService/data"
	"dev.azure.com/learn-website-orga/_git/learn-website/src/util"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertCourse(rw http.ResponseWriter, r *http.Request) {
	course := &data.CreateCourseRequest{}

	if r.Body == nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	err := util.FromJSON(course, r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}
	course.Released = time.Now()
	course.LastUpdate = time.Now()
	data.Create(course)
	rw.WriteHeader(http.StatusCreated)
}

func FindCourse(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if vars == nil || vars["courseId"] == "" {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	courseId, err := primitive.ObjectIDFromHex(vars["courseId"])
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
	}
	course := data.Find(courseId)
	if reflect.ValueOf(course).IsZero() {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte("Course with the given id was not found"))
	}
	parseErr := util.ToJSON(course, rw)
	if parseErr != nil {
		log.Print(parseErr)
		rw.WriteHeader(http.StatusInternalServerError)
	}
}

func GetAllCourses(rw http.ResponseWriter, r *http.Request) {
	courses := data.GetAllCourses()
	rw.WriteHeader(http.StatusOK)
	parseErr := util.ToJSON(courses, rw)
	if parseErr != nil {
		log.Print(parseErr)
		rw.WriteHeader(http.StatusInternalServerError)
	}
}

func UpdateCourse(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if vars == nil || vars["courseId"] == "" {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	updatedCourse := &data.UpdateCourseRequest{}
	courseId, err := primitive.ObjectIDFromHex(vars["courseId"])
	if err != nil {
		log.Print(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	parseErr := util.FromJSON(updatedCourse, r.Body)
	if parseErr != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
		return
	}
	updatedCourse.LastUpdate = time.Now()
	course := data.Update(courseId, updatedCourse)
	if reflect.ValueOf(course).IsZero() {
		rw.WriteHeader(http.StatusNotFound)
		return
	}
	util.ToJSON(course, rw)
}

func DeleteCourse(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if vars == nil || vars["courseId"] == "" {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	courseId, err := primitive.ObjectIDFromHex(vars["courseId"])
	if err != nil {
		log.Print(err)
		return
	}
	data.Delete(courseId)
	rw.WriteHeader(http.StatusNoContent)
}
