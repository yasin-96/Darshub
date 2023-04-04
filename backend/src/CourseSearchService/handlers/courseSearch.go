package handlers

import (
	"fmt"
	"net/http"
	"reflect"

	courseSearchService "darshub.dev/src/CourseSearchService/data"
	"darshub.dev/src/util"
)

func SearchCourse(rw http.ResponseWriter, r *http.Request) {
	searchTerm := r.URL.Query().Get("course")
	fmt.Println("Search term:", searchTerm)
	respData, err := courseSearchService.SearchCourse(searchTerm)
	if err == nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	if reflect.ValueOf(respData).IsZero() {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("There were no courses found for your search term"))
	}
	parseErr := util.ToJSON(respData, rw)
	if parseErr != nil {
		http.Error(rw, parseErr.Error(), http.StatusInternalServerError)
	}
}
