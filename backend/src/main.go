package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	courseHandler "dev.azure.com/learn-website-orga/_git/learn-website/backend/src/CourseService/handlers"
	userHandler "dev.azure.com/learn-website-orga/_git/learn-website/backend/src/UserService/handlers"
	"github.com/gorilla/mux"
	"github.com/nicholasjackson/env"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")

func main() {

	env.Parse()

	l := log.New(os.Stdout, "products-api ", log.LstdFlags)

	// create the handlers

	// create a new serve mux and register the handlers
	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/user/{userId}", userHandler.FindById)
	getRouter.HandleFunc("/course/{courseId}", courseHandler.FindCourse)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/user", userHandler.RegisterUser)
	postRouter.HandleFunc("/session", userHandler.Login)
	postRouter.HandleFunc("/course", courseHandler.InsertCourse)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/course/{courseId}", courseHandler.UpdateCourse)

	// create a new server
	s := http.Server{
		Addr:         *bindAddress,      // configure the bind address
		Handler:      sm,                // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server

	l.Println("Starting server on port 9090")

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
