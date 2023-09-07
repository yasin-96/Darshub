package data

import (
	"fmt"
	"time"

	"darshub.dev/src/UserService/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/phpdave11/gofpdf"
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
	Markdown    bool                 `json:"markdown" bson:"markdown"`
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
	Markdown    bool                 `json:"markdown"`
}

type UpdateCourseRequest struct {
	Name        string               `json:"name"`
	Description string               `json:"description"`
	Duration    string               `json:"duration"`
	Level       int                  `json:"level"`
	Chapters    []primitive.ObjectID `json:"content"`
	Author      string               `json:"author"`
	LastUpdate  time.Time            `json:"lastUpdate"`
	Markdown    bool                 `json:"markdown"`
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
		"markdown":    updatedCourse.Markdown,
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
	//courseTitle := course.Name

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
		pdf.SetFont("Helvetica", "", 20)
		chapterTitle(chapNum, titleStr)
		for _, subchapter := range subchapters {
			if course.Markdown {
				_, lineHt := pdf.GetFontSize()
				html := pdf.HTMLBasicNew()
				convertedHTML := convertMarkdown(subchapter.Content)
				println(string(convertedHTML))
				html.Write(lineHt, string(convertedHTML))
				err := pdf.OutputFileAndClose(fmt.Sprintf("%s.pdf", course.Name))
				if err != nil {
					println(err)
				}
				return
			}
			printSubchapter(subchapter, subchapterCounter)
			subchapterCounter++
		}
	}

	for chapterTitle, subchapters := range x {
		printChapter(chapterCounter, chapterTitle, subchapters)
		chapterCounter++
	}

	return nil
}

func convertMarkdown(subchapterContent string) []byte {
	extensions := parser.CommonExtensions | parser.OrderedListStart | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse([]byte(subchapterContent))

	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}
