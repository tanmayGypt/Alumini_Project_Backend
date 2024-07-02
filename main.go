package main

import (
	"fmt"
	"log"
	"net/http"

	routes "my-go-backend/Routes"
	database "my-go-backend/config"

	"github.com/gorilla/mux"
)

func main() {
	database.DatabaseConnector()
	// Migrate the schema

	router := mux.NewRouter()
	routes.InitializeRoutes(router)

	fmt.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
