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
	print(courseId.Hex())
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	course, err := data.Find(courseId)
	if err != nil {
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
	rw.WriteHeader(http.StatusNoContent)

	respErr := data.Delete(courseId)
	if respErr != nil {
		http.Error(rw, respErr.Error(), http.StatusInternalServerError)
		return
	}
	log.Print("Course was deleted successfully")
}

func GeneratePDF(rw http.ResponseWriter, r *http.Request) {
	id, _ := primitive.ObjectIDFromHex("64c0e325199e180b2d94d66d")
	course, _ := data.Find(id)
	println(course.Description)

	/*pdf := gofpdf.New("P", "mm", "A4", "")
	titleStr := "testtitle"
	pdf.SetTitle(titleStr, false)
	pdf.SetAuthor("Jules Verne", false)
	pdf.SetHeaderFunc(func() {
		// Arial bold 15
		pdf.SetFont("Arial", "B", 15)
		// Calculate width of title and position
		wd := pdf.GetStringWidth(titleStr) + 6
		pdf.SetX((210 - wd) / 2)
		// Colors of frame, background and text
		pdf.SetDrawColor(0, 80, 180)
		pdf.SetFillColor(230, 230, 0)
		pdf.SetTextColor(220, 50, 50)
		// Thickness of frame (1 mm)
		pdf.SetLineWidth(1)
		// Title
		pdf.CellFormat(wd, 9, titleStr, "1", 1, "C", true, 0, "")
		// Line break
		pdf.Ln(10)
	})
	pdf.SetFooterFunc(func() {
		// Position at 1.5 cm from bottom
		pdf.SetY(-15)
		// Arial italic 8
		pdf.SetFont("Arial", "I", 8)
		// Text color in gray
		pdf.SetTextColor(128, 128, 128)
		// Page number
		pdf.CellFormat(0, 10, fmt.Sprintf("Page %d", pdf.PageNo()),
			"", 0, "C", false, 0, "")
	})
	chapterTitle := func(chapNum int, titleStr string) {
		// 	// Arial 12
		pdf.SetFont("Arial", "", 12)
		// Background color
		pdf.SetFillColor(200, 220, 255)
		// Title
		pdf.CellFormat(0, 6, fmt.Sprintf("Chapter %d : %s", chapNum, titleStr),
			"", 1, "L", true, 0, "")
		// Line break
		pdf.Ln(4)
	}
	chapterBody := func(fileStr string) {
		// Read text file
		txtStr, err := ioutil.ReadFile(fileStr)
		if err != nil {
			pdf.SetError(err)
		}
		pdf.AddUTF8Font("dejavu", "", "DejaVuSansCondensed.ttf")
		pdf.AddUTF8Font("dejavu", "I", "DejaVuSansCondensed-Oblique.ttf")

		// Times 12
		pdf.SetFont("dejavu", "", 14)
		// Output justified text
		pdf.MultiCell(0, 5, string(txtStr), "", "", false)
		// Line break
		pdf.Ln(-1)
		// Mention in italics
		pdf.SetFont("dejavu", "I", 0)
		pdf.Cell(0, 5, "(end of excerpt)")
	}
	printChapter := func(chapNum int, titleStr, fileStr string) {
		pdf.AddPage()
		chapterTitle(chapNum, titleStr)
		chapterBody(fileStr)
	}*/

	var m map[string][]data.Subchapter

	for _, id := range course.Chapters {
		chapter, _ := data.FindChapter(id)
		print(chapter.Name)
		subchapterid := chapter.Subchapters[0]
		subchapter, _ := data.FindSubchapter(subchapterid)
		res := append(m[chapter.Name], subchapter)
		print(res)
	}
	/*
	   printChapter(1, "A RUNAWAY REEF", "test.txt")
	   printChapter(2, "THE PROS AND CONS", "test1.txt")
	   err := pdf.OutputFileAndClose("hello.pdf")

	   	if err != nil {
	   		fmt.Print(err.Error())
	   	}
	*/
}
