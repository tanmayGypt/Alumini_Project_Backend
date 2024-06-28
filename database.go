package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var DB *gorm.DB
var err error


func DatabaseConnector()  {
	dsn := "root:root@tcp(127.0.0.1:3306)/alumni_db?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}
	DB.AutoMigrate(&AlumniProfile{}, &ProfessionalInformation{}, &AlumniAttending{})
}