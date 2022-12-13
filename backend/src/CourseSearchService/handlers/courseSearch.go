package handlers

import (
	"fmt"
	"net/http"

	courseSearchService "dev.azure.com/learn-website-orga/_git/learn-website/src/CourseSearchService/data"
)

func SearchCourse(rw http.ResponseWriter, r *http.Request) {
	searchTerm := r.URL.Query().Get("course")
	fmt.Println("Search term:", searchTerm)
	courseSearchService.SearchCourse(searchTerm)
}
