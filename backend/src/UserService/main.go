package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	userHandler "darshub.dev/src/UserService/handlers"
	"darshub.dev/src/UserService/middleware"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/nicholasjackson/env"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":8080", "Bind address for the server")
var allowedMethods = "OPTIONS,POST,PUT,DELETE,GET"
var allowedHeaders = "Origin, Content-Type"

func main() {
	env.Parse()
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalf("Error loading the .env file: %v", err)
	}

	l := log.New(os.Stdout, "darshub-api-user-service", log.LstdFlags)

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet, http.MethodOptions).Subrouter()
	getRouter.Use(middleware.EnsureValidToken())
	getRouter.HandleFunc("/users/{userId}", userHandler.FindById)
	getRouter.HandleFunc("/user/{userId}/setInactive", userHandler.SetAccountInactive)
	getRouter.HandleFunc("/users", userHandler.GetAllUsers)

	postRouter := sm.Methods(http.MethodPost, http.MethodOptions).Subrouter()
	postRouter.Use(middleware.EnsureValidToken())
	postRouter.HandleFunc("/user", userHandler.RegisterUser)

	putRouter := sm.Methods(http.MethodPut, http.MethodOptions).Subrouter()
	putRouter.Use(middleware.EnsureValidToken())
	putRouter.HandleFunc("/user/{userId}", userHandler.UpdateUser)

	deleteRouter := sm.Methods(http.MethodDelete, http.MethodOptions).Subrouter()
	deleteRouter.Use(middleware.EnsureValidToken())
	deleteRouter.HandleFunc("/user/{userId}", userHandler.DeleteUser)

	//middleware
	sm.Use(corsMiddleware)

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

	l.Println("Starting server on port 8080")

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
		w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5173")
		w.Header().Set("Access-Control-Allow-Methods", allowedMethods)
		w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusCreated)
			return
		}

		next.ServeHTTP(w, r)
	})

}
