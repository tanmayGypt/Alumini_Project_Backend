package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	helper "my-go-backend/Helper"
	"net/http"
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

	err = helper.SendEmail(email)
	if err != nil {
		log.Printf("Error sending email: %v\n", err)
		http.Error(w, "Failed to send email", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Email received"))
}
