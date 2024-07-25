package main

import (
	"fmt"
	"log"
	"net/http"

	routes "my-go-backend/Routes"
	database "my-go-backend/config"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Error loading .env file, proceeding with environment variables: ", err)
	}
	database.DatabaseConnector()
	// Migrate the schema

	router := mux.NewRouter()
	routes.InitializeRoutes(router)
	// Set up the CORS middleware
	corsOptions := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}), // Allow all origins, change "*" to specific origins if needed
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	// Apply the CORS middleware to your router
	fmt.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", corsOptions(router)))
}
