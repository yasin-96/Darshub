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

func InsertCourseCategory(rw http.ResponseWriter, r *http.Request) {
	courseCategory := &data.CreateCourseCategoryRequest{}

	err := util.FromJSON(courseCategory, r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	data.CreateCourseCategory(courseCategory)
	rw.WriteHeader(http.StatusCreated)
}

func FindCourseCategory(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	courseCategoryId, err := primitive.ObjectIDFromHex(vars["courseCategoryId"])
	if err != nil {
		log.Fatal(err)
	}

	courseCategory := data.FindCourseCategory(courseCategoryId)
	if reflect.ValueOf(courseCategory).IsZero() {
		rw.WriteHeader(http.StatusNotFound)
	}

	rw.WriteHeader(http.StatusOK)
	parseErr := util.ToJSON(courseCategory, rw)
	if parseErr != nil {
		log.Fatal(err)
	}
}

func UpdateCourseCategory(rw http.ResponseWriter, r *http.Request) {
	updatedCourseCategory := &data.UpdateCourseCategoryRequest{}
	vars := mux.Vars(r)
	courseCategoryId, err := primitive.ObjectIDFromHex(vars["courseId"])
	if err != nil {
		log.Fatal(err)
	}

	parseErr := util.FromJSON(updatedCourseCategory, r.Body)
	if parseErr != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	course := data.UpdateCourseCategory(courseCategoryId, updatedCourseCategory)
	rw.WriteHeader(http.StatusOK)
	util.ToJSON(course, rw)
}

func DeleteCourseCategory(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	courseId, err := primitive.ObjectIDFromHex(vars["courseId"])
	if err != nil {
		log.Fatal(err)
	}

	data.Delete(courseId)
	rw.WriteHeader(http.StatusNoContent)
}
