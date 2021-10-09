package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"dev.azure.com/learn-website-orga/_git/learn-website/backend/src/handlers"
	"github.com/gorilla/mux"
	"github.com/nicholasjackson/env"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")

func main() {

	env.Parse()

	l := log.New(os.Stdout, "products-api ", log.LstdFlags)

	client, err := mongo.NewClient(options.Client().
		ApplyURI("mongodb+srv://dhub:OiPe7pU8kxaIVhBx@dhcluster001.c17aj.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))

	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	darshubDb := client.Database("darshub")
	userCollection := darshubDb.Collection("user")
	_, error := userCollection.InsertOne(ctx, bson.D{
		{Key: "title", Value: "Dev"},
	})

	if error != nil {
		log.Fatal(error)
	}

	// create the handlers
	ph := handlers.NewProducts(l)

	// create a new serve mux and register the handlers
	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ph.GetProducts)

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
	go func() {
		l.Println("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	context, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(context)
}
