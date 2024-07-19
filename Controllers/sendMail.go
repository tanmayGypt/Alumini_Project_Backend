package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	helper "my-go-backend/Helper"
	models "my-go-backend/Models"
	database "my-go-backend/config"
	"net/http"

	"gorm.io/gorm"
)

type EmailRequest struct {
	Email string `json:"email"`
}

func SendEmail(w http.ResponseWriter, r *http.Request) {
	var emailReq EmailRequest
	err := json.NewDecoder(r.Body).Decode(&emailReq)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	email := emailReq.Email
	fmt.Printf("Received email: %s\n", email)

	// Check if email exists in AlumniProfile
	var alumni models.AlumniProfile
	err = database.DB.Where("email = ?", email).First(&alumni).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Email not found in AlumniProfile", http.StatusNotFound)
			return
		}
		log.Printf("Error searching for email: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	err = helper.SendEmail(email)
	if err != nil {
		log.Printf("Error sending email: %v\n", err)
		http.Error(w, "Failed to send email", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Email received"))
}
