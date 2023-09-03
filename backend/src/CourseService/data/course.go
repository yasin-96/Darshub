package data

import (
	"fmt"
	"time"

	"darshub.dev/src/UserService/config"
	"github.com/phpdave11/gofpdf"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Course struct {
	ID          primitive.ObjectID   `json:"_id" bson:"_id"`
	Name        string               `json:"name" bson:"name"`
	Description string               `json:"description" bson:"description"`
	Duration    string               `json:"duration" bson:"duration"`
	Level       int32                `json:"level" bson:"level"`
	Chapters    []primitive.ObjectID `json:"chapters" bson:"chapters"`
	Author      string               `json:"author" bson:"author"`
	Released    time.Time            `json:"released" bson:"released"`
	LastUpdate  time.Time            `json:"lastUpdate" bson:"lastUpdate"`
}

type CreateCourseRequest struct {
	Name        string               `json:"name"`
	Description string               `json:"description"`
	Duration    string               `json:"duration"`
	Level       int                  `json:"level"`
	Chapters    []primitive.ObjectID `json:"chapters"`
	Author      string               `json:"author"`
	Released    time.Time            `json:"released"`
	LastUpdate  time.Time            `json:"lastUpdate"`
}

type UpdateCourseRequest struct {
	Name        string               `json:"name"`
	Description string               `json:"description"`
	Duration    string               `json:"duration"`
	Level       int                  `json:"level"`
	Chapters    []primitive.ObjectID `json:"content"`
	Author      string               `json:"author"`
	LastUpdate  time.Time            `json:"lastUpdate"`
}

func Create(course *CreateCourseRequest) error {
	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	_, err := client.Database("darshub").Collection("course").InsertOne(ctx, course)
	return err
}

func Find(courseId primitive.ObjectID) (Course, error) {
	var course Course
	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	err := client.Database("darshub").Collection("course").FindOne(ctx, bson.M{"_id": courseId}).Decode(&course)
	if err != nil {
		return Course{}, err
	}
	return course, nil
}

func GetAllCourses() ([]Course, error) {
	var courses []Course
	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	cur, err := client.Database("darshub").Collection("course").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	cur.All(ctx, &courses)
	return courses, nil
}

func Update(courseId primitive.ObjectID, updatedCourse *UpdateCourseRequest) (Course, error) {
	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	update := bson.M{
		"name":        updatedCourse.Name,
		"description": updatedCourse.Description,
		"duration":    updatedCourse.Duration,
		"level":       updatedCourse.Level,
		"chapters":    updatedCourse.Chapters,
		"author":      updatedCourse.Author,
		"lastUpdate":  updatedCourse.LastUpdate,
	}

	_, err := client.Database("darshub").Collection("course").ReplaceOne(ctx, bson.M{"_id": courseId}, update)
	if err != nil {
		return Course{}, err
	}
	course, respErr := Find(courseId)
	if respErr != nil {
		return Course{}, err
	}
	return course, nil
}

func Delete(courseId primitive.ObjectID) error {
	ctx, cancel, client := config.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	_, err := client.Database("darshub").Collection("course").DeleteOne(ctx, bson.M{"_id": courseId})
	return err
}

func GeneratePDF(courseId primitive.ObjectID) error {
	course, err := Find(courseId)
	if err != nil {
		return err
	}

	x := make(map[string][]Subchapter)

	for _, chapterId := range course.Chapters {

		chapter, _ := FindChapter(chapterId)

		for _, subchapterid := range chapter.Subchapters {
			subchapter, _ := FindSubchapter(subchapterid)
			x[chapter.Name] = append(x[chapter.Name], subchapter)
		}
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	courseTitle := course.Name
	pdf.SetTitle(courseTitle, false)
	pdf.SetAuthor(course.Author, false)
	pdf.SetHeaderFunc(func() {
		// Arial bold 15
		pdf.SetFont("Arial", "B", 15)
		// Calculate width of title and position
		wd := pdf.GetStringWidth(courseTitle) + 6
		pdf.SetX((210 - wd) / 2)
		// Colors of frame, background and text
		pdf.SetDrawColor(0, 80, 180)
		pdf.SetFillColor(230, 230, 0)
		pdf.SetTextColor(220, 50, 50)
		// Thickness of frame (1 mm)
		pdf.SetLineWidth(1)
		// Title
		pdf.CellFormat(wd, 9, courseTitle, "1", 1, "C", true, 0, "")
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
	subchapterTitle := func(subchappNum int, subchapTitle string) {
		// 	// Arial 12
		pdf.SetFont("Arial", "", 9)
		// Background color
		pdf.SetFillColor(200, 220, 255)
		// Title
		pdf.CellFormat(0, 6, fmt.Sprintf("Subchapter %d : %s", subchappNum, subchapTitle),
			"", 1, "L", true, 0, "")
		// Line break
		pdf.Ln(4)
	}
	chapterBody := func(content string) {
		// Read text file

		pdf.AddUTF8Font("dejavu", "", "DejaVuSansCondensed.ttf")
		pdf.AddUTF8Font("dejavu", "I", "DejaVuSansCondensed-Oblique.ttf")

		// Times 12
		pdf.SetFont("dejavu", "", 14)
		// Output justified text
		pdf.MultiCell(0, 5, content, "", "", false)
		// Line break
		pdf.Ln(-1)
		// Mention in italics
		pdf.SetFont("dejavu", "I", 0)
	}

	chapterCounter := 1
	subchapterCounter := 1

	printSubchapter := func(subhchapter Subchapter, subchapNum int) {
		subchapterTitle(subchapNum, subhchapter.Name)
		chapterBody(subhchapter.Content)
	}

	printChapter := func(chapNum int, titleStr string, subchapters []Subchapter) {
		pdf.AddPage()
		chapterTitle(chapNum, titleStr)
		for _, subchapter := range subchapters {
			printSubchapter(subchapter, subchapterCounter)
			subchapterCounter++
		}
	}

	for chapterTitle, subchapters := range x {
		printChapter(chapterCounter, chapterTitle, subchapters)
		chapterCounter++
	}

	_ = pdf.OutputFileAndClose(fmt.Sprintf("%s.pdf", course.Name))

	return nil
}
