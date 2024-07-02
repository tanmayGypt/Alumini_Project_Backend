package database

import (
	"database/sql"
	"log"
	models "my-go-backend/Models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func DatabaseConnector() {
	// First, check if the database exists and create it if it doesn't
	dsn := "root:87654321@tcp(127.0.0.1:3306)/"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("failed to connect to MySQL server: ", err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS alumini_db")
	if err != nil {
		log.Fatal("failed to create database: ", err)
	}

	// Now, connect to the newly created or existing database
	dsn = "root:87654321@tcp(127.0.0.1:3306)/alumini_db?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
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
	err = DB.AutoMigrate(&models.AlumniAttending{})
	if err != nil {
		log.Fatalf("failed to migrate AlumniAttending: %v", err)
	}
}
