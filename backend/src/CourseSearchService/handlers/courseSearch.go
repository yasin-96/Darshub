package handlers

import (
	"fmt"
	"log"
	"net/http"
	"reflect"

	courseSearchService "darshub.dev/src/CourseSearchService/data"
	"darshub.dev/src/util"
)

func SearchCourse(rw http.ResponseWriter, r *http.Request) {
	searchTerm := r.URL.Query().Get("course")
	fmt.Println("Search term:", searchTerm)
	respData := courseSearchService.SearchCourse(searchTerm)

	if reflect.ValueOf(respData).IsZero() {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte("No Data found"))
		return
	}
	parseErr := util.ToJSON(respData, rw)
	if parseErr != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		log.Print(parseErr)
	}
}
