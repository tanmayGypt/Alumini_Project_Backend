package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	helper "my-go-backend/Helper"
	"net/http"
	"os"
)

type Contactform struct {
	Name    string
	Email   string
	Subject string
    Contact string
	Message string
}

func ContactUSHandler(w http.ResponseWriter, r *http.Request) {

	emailUser := os.Getenv("EMAIL_USER")

	var data Contactform
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	emailBody := fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
    <title>New Contact Us Form Submission</title>
    <style>
        body {
            font-family: Arial, sans-serif;
        }
        .container {
            padding: 20px;
        }
        .header {
            background-color: #f8f9fa;
            padding: 10px 20px;
            border-bottom: 1px solid #dee2e6;
        }
        .content {
            margin: 20px 0;
        }
        .footer {
            background-color: #f8f9fa;
            padding: 10px 20px;
            border-top: 1px solid #dee2e6;
            text-align: center;
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h2>New Contact Us Form Submission</h2>
        </div>
        <div class="content">
            <p><strong>Name:</strong> %s</p>
            <p><strong>Email:</strong> %s</p>
            <p><strong>Contact:</strong> %s</p>
            <p><strong>Subject:</strong> %s</p>
            <p><strong>Message:</strong></p>
            <p>%s</p>
        </div>
        <div class="footer">
            <p>This message was sent from the BPIT Alumni Website Contact Us form.</p>
        </div>
    </div>
</body>
</html>`, data.Name, data.Email,data.Contact,data.Subject, data.Message)
	err := helper.SendEmail(emailUser, "New Contact Us Form Submission from BPIT Alumni Website", emailBody)
	if err != nil {
		log.Printf("Error sending email: %v\n", err)
		http.Error(w, "Failed to send email", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "mail sent successfully",
	})
}
