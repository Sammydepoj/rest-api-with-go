package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sammydepoj/golang-rest-api/dbconfig"
	"github.com/sammydepoj/golang-rest-api/internal/handlers"
	"github.com/sammydepoj/golang-rest-api/internal/routes"
	"github.com/sammydepoj/golang-rest-api/internal/store"
	"github.com/sammydepoj/golang-rest-api/serverconfig"
)

func main() {
	//Load config
	config, err := serverconfig.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
		fmt.Printf("Error loading config: %v\n", err)
		return
	}

	//connect to DB
	db := dbconfig.ConnectDB(config.DatabaseUrl)
	defer db.Close()

	queries := store.New(db)
	//create handler
	handler := handlers.NewHandlers(db, queries)

	// set up the http server
	mux := http.NewServeMux()

	// setup routes
	routes.SetupRoutes(mux, handler)

	serverAddr := fmt.Sprintf(":%s", config.ServerPort)
	server := &http.Server{
		Addr:    serverAddr,
		Handler: mux,
	}
	fmt.Printf("Starting server on port %s...\n", config.ServerPort)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Error starting server: %v", err)
		fmt.Printf("Error starting server: %v\n", err)
		return
	}

}
