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

func InsertCourseCategory(rw http.ResponseWriter, r *http.Request) {

	if r.Body == nil {
		http.Error(rw, "Body is empty", http.StatusBadRequest)
	}

	courseCategory := &data.CreateCourseCategoryRequest{}
	err := util.FromJSON(courseCategory, r.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	respErr := data.CreateCourseCategory(courseCategory)
	if respErr != nil {
		http.Error(rw, respErr.Error(), http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusCreated)
}

func FindCourseCategory(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if vars == nil || vars["courseCategoryId"] == "" {
		http.Error(rw, "No course category id was provided in the path variable", http.StatusBadRequest)
		return
	}

	courseCategoryId, err := primitive.ObjectIDFromHex(vars["courseCategoryId"])
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	courseCategory := data.FindCourseCategory(courseCategoryId)
	if reflect.ValueOf(courseCategory).IsZero() {
		http.Error(rw, "The course category with the given id was not found", http.StatusNotFound)
		return
	}

	parseErr := util.ToJSON(courseCategory, rw)
	if parseErr != nil {
		http.Error(rw, parseErr.Error(), http.StatusInternalServerError)
		return
	}
}

func GetAllCourseCategoryNames(rw http.ResponseWriter, r *http.Request) {
	courseCategories := data.FindAllCourses()
	courseCategoryNames := make([]string, 0)
	for _, v := range courseCategories {
		courseCategoryNames = append(courseCategoryNames, v.Name)
	}
	parseErr := util.ToJSON(courseCategoryNames, rw)
	if parseErr != nil {
		http.Error(rw, parseErr.Error(), http.StatusInternalServerError)
	}
}

func UpdateCourseCategory(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if vars == nil || vars["courseCategoryId"] == "" {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	updatedCourseCategory := &data.UpdateCourseCategoryRequest{}

	courseCategoryId, err := primitive.ObjectIDFromHex(vars["courseCategoryId"])
	if err != nil {
		log.Print(err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	parseErr := util.FromJSON(updatedCourseCategory, r.Body)
	if parseErr != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
		return
	}

	course := data.UpdateCourseCategory(courseCategoryId, updatedCourseCategory)
	util.ToJSON(course, rw)
}

func DeleteCourseCategory(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if vars == nil || vars["courseCategoryId"] == "" {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	courseCategoryId, err := primitive.ObjectIDFromHex(vars["courseCategoryId"])
	if err != nil {
		return
	}
	respErr := data.DeleteCourseCategory(courseCategoryId)
	if respErr != nil {
		http.Error(rw, respErr.Error(), http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusNoContent)
}
