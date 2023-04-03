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
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

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

	if vars == nil || vars["courseCategoryId"] == "" {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	courseCategoryId, err := primitive.ObjectIDFromHex(vars["courseCategoryId"])
	if err != nil {
		log.Print(err)
		rw.WriteHeader(http.StatusInternalServerError)
	}

	courseCategory := data.FindCourseCategory(courseCategoryId)
	if reflect.ValueOf(courseCategory).IsZero() {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte("The course category with the given id was not found."))
		return
	}

	parseErr := util.ToJSON(courseCategory, rw)
	if parseErr != nil {
		log.Print(err)
		rw.WriteHeader(http.StatusInternalServerError)
	}
}

func GetAllCourseCategoryNames(rw http.ResponseWriter, r *http.Request) {
	println("test")
	courseCategories := data.FindAllCourses()
	courseCategoryNames := make([]string, 0)
	for _, v := range courseCategories {
		courseCategoryNames = append(courseCategoryNames, v.Name)
	}
	rw.WriteHeader(http.StatusOK)
	parseErr := util.ToJSON(courseCategoryNames, rw)
	if parseErr != nil {
		log.Print(parseErr)
		rw.WriteHeader(http.StatusInternalServerError)
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
	}

	course := data.UpdateCourseCategory(courseCategoryId, updatedCourseCategory)
	rw.WriteHeader(http.StatusOK)
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
		log.Print(err)
		return
	}
	data.DeleteCourseCategory(courseCategoryId)
	rw.WriteHeader(http.StatusNoContent)
}
