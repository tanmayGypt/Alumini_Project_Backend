package database

import (
	"log"
	models "my-go-backend/Models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func DatabaseConnector() {
	dsn := "root:Mohit@2005@tcp(127.0.0.1:3306)/alumini_database?charset=utf8mb4&parseTime=True&loc=Local"
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
