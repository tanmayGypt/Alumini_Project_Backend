package main

import (
	"fmt"
	"log"
	"net/http"
)

// homePage function will handle requests to the root URL.
func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

// handleRequests function will set up the routes and start the server.
func handleRequests() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

// main function is the entry point of the program.
func main() {
    handleRequests()
}
