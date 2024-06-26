package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

const (
	DBUsername = "root"
	DBPassword = "87654321"
	DBHost     = "localhost"
	DBPort     = "3306"
	DBName     = "alumni_db"
)

// initialize database connection
func InitializeDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DBUsername, DBPassword, DBHost, DBPort, DBName) // Data source name

	// open database
	database, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}

	// assign database connection to global variable
	db = database

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	fmt.Println("Connected to database!")
}

// close database connection
func CloseDB() {
	if db != nil {
		db.Close()
		fmt.Println("Database connection closed.")
	}
}
