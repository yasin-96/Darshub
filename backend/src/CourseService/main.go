package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	courseHandler "darshub.dev/src/CourseService/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var allowedMethods = "OPTIONS,POST,PUT,DELETE,GET"
var allowedHeaders = "Origin, Content-Type"

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		println("Enviroment variables could not be loaded.")
	}

	var bindAddress = os.Getenv("COURSE_SERVICE_PORT")
	l := log.New(os.Stdout, "darshub-api-course-service", log.LstdFlags)

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet, http.MethodOptions).Subrouter()
	getRouter.HandleFunc("/course/{courseId}", courseHandler.FindCourse)
	getRouter.HandleFunc("/courseCategory/{courseCategoryId}", courseHandler.FindCourseCategory)
	getRouter.HandleFunc("/chapter/{chapterId}", courseHandler.FindChapter)
	getRouter.HandleFunc("/subchapter/{subchapterId}", courseHandler.FindSubchapter)
	getRouter.HandleFunc("/courses", courseHandler.GetAllCourses)
	getRouter.HandleFunc("/courseCategoryNames/all", courseHandler.GetAllCourseCategoryNames)

	postRouter := sm.Methods(http.MethodPost, http.MethodOptions).Subrouter()
	postRouter.HandleFunc("/course", courseHandler.InsertCourse)
	postRouter.HandleFunc("/courseCategory", courseHandler.InsertCourseCategory)
	postRouter.HandleFunc("/chapter", courseHandler.InsertChapter)
	postRouter.HandleFunc("/subchapter", courseHandler.InsertSubchapter)
	postRouter.HandleFunc("/course/{courseId}/register", courseHandler.RegisterUserToCourse)

	putRouter := sm.Methods(http.MethodPut, http.MethodOptions).Subrouter()
	putRouter.HandleFunc("/course/{courseId}", courseHandler.UpdateCourse)
	putRouter.HandleFunc("/courseCategory/{courseCategoryId}", courseHandler.UpdateCourseCategory)
	putRouter.HandleFunc("/chapter/{chapterId}", courseHandler.UpdateChapter)
	putRouter.HandleFunc("/subchapter/{subchapterId}", courseHandler.UpdateSubchapter)

	deleteRouter := sm.Methods(http.MethodDelete, http.MethodOptions).Subrouter()
	deleteRouter.HandleFunc("/course/{courseId}", courseHandler.DeleteCourse)
	deleteRouter.HandleFunc("/courseCategory/{courseCategoryId}", courseHandler.DeleteCourseCategory)
	deleteRouter.HandleFunc("/chapter/{chapterId}", courseHandler.DeleteChapter)
	deleteRouter.HandleFunc("/subchapter/{subchapterId}", courseHandler.DeleteSubchapter)

	//middleware
	sm.Use(corsMiddleware)

	// create a new server
	s := http.Server{
		Addr:         bindAddress,       // configure the bind address
		Handler:      sm,                // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server

	l.Println("Starting server on port 8081")

	errLis := s.ListenAndServe()
	if errLis != nil {
		l.Printf("Error starting server: %s\n", errLis)
		os.Exit(1)
	}

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}

// All Request with Options will be ignored
// TODO Specify the origins ans methods from const variable
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", os.Getenv("FRONTEND_IP"))
		w.Header().Set("Access-Control-Allow-Methods", allowedMethods)
		w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusCreated)
			return
		}

		next.ServeHTTP(w, r)
	})

}
