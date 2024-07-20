package database

import (
	"database/sql"
	"log"
	models "my-go-backend/Models"
	"os"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func DatabaseConnector() {
	// Retrieve environment variables for the database
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	sslMode := os.Getenv("SSL_MODE")

	if dbName == "" || dbUser == "" || dbPass == "" || dbHost == "" || dbPort == "" || sslMode == "" {
		log.Fatal("Missing required database environment variables")
	}
	// First, check if the database exists and create it if it doesn't
	serverDSN := "postgres://" + dbUser + ":" + dbPass + "@" + dbHost + ":" + dbPort + "/postgres?sslmode=" + sslMode
	db, err := sql.Open("postgres", serverDSN)
	if err != nil {
		log.Fatal("failed to connect to MySQL server: ", err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	if err != nil {
		log.Fatal("failed to create database: ", err)
	}

	// Now, connect to the newly created or existing database
	dsn := "postgresql://" + dbUser + ":" + dbPass + "@" + dbHost + ":" + dbPort + "/" + dbName + "?sslmode=" + sslMode
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}
	// First, migrate the AlumniProfile table
	err = DB.AutoMigrate(&models.AlumniProfile{})
	if err != nil {
		log.Fatalf("failed to migrate AlumniProfile: %v", err)
	}

	// Then migrate the related tables
	err = DB.AutoMigrate(&models.ProfessionalInformation{}, &models.Achievement{}, &models.InterestHobby{}, &models.Event{})
	if err != nil {
		log.Fatalf("failed to migrate related tables: %v", err)
	}

	// Finally, migrate the AlumniAttending table
	err = DB.AutoMigrate(&models.AlumniAttending{}, &models.InterviewExperience{})
	if err != nil {
		log.Fatalf("failed to migrate AlumniAttending: %v", err)
	}
}
