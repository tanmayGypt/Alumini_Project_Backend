package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func main() {
	DatabaseConnector()
	// Migrate the schema
	
	router := mux.NewRouter()
	InitializeRoutes(router)

	fmt.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
