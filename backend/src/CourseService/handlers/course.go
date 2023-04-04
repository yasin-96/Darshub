package handlers

import (
	"log"
	"net/http"
	"reflect"
	"time"

	"darshub.dev/src/CourseService/data"
	"darshub.dev/src/util"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertCourse(rw http.ResponseWriter, r *http.Request) {
	course := &data.CreateCourseRequest{}

	if r.Body == nil {
		http.Error(rw, "Body is empty", http.StatusBadRequest)
		return
	}

	err := util.FromJSON(course, r.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	course.Released = time.Now()
	course.LastUpdate = time.Now()
	respErr := data.Create(course)
	if respErr != nil {
		http.Error(rw, respErr.Error(), http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusCreated)
}

func FindCourse(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if vars == nil || vars["courseId"] == "" {
		http.Error(rw, "No course id provided in path variable", http.StatusBadRequest)
		return
	}

	courseId, err := primitive.ObjectIDFromHex(vars["courseId"])
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	course, err := data.Find(courseId)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	if reflect.ValueOf(course).IsZero() {
		http.Error(rw, "Course with the given id was not found", http.StatusNotFound)
		return
	}
	parseErr := util.ToJSON(course, rw)
	if parseErr != nil {
		http.Error(rw, parseErr.Error(), http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusOK)
}

func GetAllCourses(rw http.ResponseWriter, r *http.Request) {
	courses, err := data.GetAllCourses()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	parseErr := util.ToJSON(courses, rw)
	if parseErr != nil {
		http.Error(rw, parseErr.Error(), http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusOK)
	util.ToJSON(courses, rw)
}

func UpdateCourse(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if vars == nil || vars["courseId"] == "" {
		http.Error(rw, "No course id provided in path variable", http.StatusBadRequest)
		return
	}

	updatedCourse := &data.UpdateCourseRequest{}
	courseId, err := primitive.ObjectIDFromHex(vars["courseId"])
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	parseErr := util.FromJSON(updatedCourse, r.Body)
	if parseErr != nil {
		http.Error(rw, parseErr.Error(), http.StatusBadRequest)
		return
	}
	updatedCourse.LastUpdate = time.Now()
	course, respErr := data.Update(courseId, updatedCourse)
	if respErr != nil {
		http.Error(rw, respErr.Error(), http.StatusInternalServerError)
		return
	}
	if reflect.ValueOf(course).IsZero() {
		http.Error(rw, "The course with the given id was not found", http.StatusNotFound)
		return
	}
	rw.WriteHeader(http.StatusOK)
	util.ToJSON(course, rw)
}

func DeleteCourse(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if vars == nil || vars["courseId"] == "" {
		http.Error(rw, "No course id was provided in the path variable", http.StatusBadRequest)
		return
	}

	courseId, err := primitive.ObjectIDFromHex(vars["courseId"])
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusNoContent)

	respErr := data.Delete(courseId)
	if respErr != nil {
		http.Error(rw, respErr.Error(), http.StatusInternalServerError)
		return
	}
	log.Print("Course was deleted successfully")
}
