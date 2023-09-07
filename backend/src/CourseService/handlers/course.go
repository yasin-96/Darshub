package handlers

import (
	"log"
	"net/http"
	"reflect"
	"time"

	"darshub.dev/src/CourseService/data"
	"darshub.dev/src/util"
	"github.com/gomarkdown/markdown"
	"github.com/gorilla/mux"
	"github.com/phpdave11/gofpdf"
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
	print(courseId.Hex())
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	course, err := data.Find(courseId)
	if err != nil {
		println("error here")
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	parseErr := util.ToJSON(course, rw)
	if parseErr != nil {
		http.Error(rw, parseErr.Error(), http.StatusInternalServerError)
		return
	}
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

	respErr := data.Delete(courseId)
	if respErr != nil {
		http.Error(rw, respErr.Error(), http.StatusInternalServerError)
		return
	}
	log.Print("Course was deleted successfully")
}

func GeneratePDF(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if vars == nil || vars["courseId"] == "" {
		http.Error(rw, "No course id was provided in the path variable", http.StatusBadRequest)
		return
	}

	courseId, parseErr := primitive.ObjectIDFromHex(vars["courseId"])
	if parseErr != nil {
		http.Error(rw, parseErr.Error(), http.StatusInternalServerError)
		return
	}

	err := data.GeneratePDF(courseId)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusNotFound)
		return
	}
}

func GenPDF(rw http.ResponseWriter, r *http.Request) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	html := pdf.HTMLBasicNew()
	htmlStr := `---

	## Chapter outcomes
	At the end of this page you should be able to:
	- Explain how the Cloud Native Computing Foundation (CNCF) defines the Cloud Native terms and their core characteristics.
	- Describe monolith vs. microservices architecture approaches with examples as well as the benefits and drawbacks of both approaches.
	- Describe different autoscaling options in Cloud Native environments.
	- Describe the concept and benefits of serverless computing.
	- Explain community and governance 
	- Explain the roles and personas that exist within Cloud Native environments and their corresponding tasks?
	- Describe which open standards are available in Cloud Native world and what they are used for?
	
	---`

	res := convertMarkdown(htmlStr)
	println(string(res))
	pdf.SetFont("Helvetica", "", 20)
	_, lineHt := pdf.GetFontSize()
	html.Write(lineHt, string(res))
	err := pdf.OutputFileAndClose("course.pdf")
	if err != nil {
		println(err)
	}
}

func convertMarkdown(subchapterContent string) []byte {
	return markdown.ToHTML([]byte(subchapterContent), nil, nil)
}
