package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	helper "my-go-backend/Helper"
	"net/http"
	"os"
)

type Feedbackform struct {
	Name     string
	Email    string
	Feedback string
}

func FeedbackHandler(w http.ResponseWriter, r *http.Request) {

	emailUser := os.Getenv("EMAIL_USER")

	var data Feedbackform
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	emailBody := fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
    <title>New Feedback Form Submission</title>
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
            <h2>New Feedback Form Submission</h2>
        </div>
        <div class="content">
            <p><strong>Name:</strong> %s</p>
            <p><strong>Email:</strong> %s</p>
            <p><strong>Feedback:</strong></p>
            <p>%s</p>
        </div>
        <div class="footer">
            <p>This message was sent from the BPIT Alumni Website Feedback form.</p>
        </div>
    </div>
</body>
</html>`, data.Name, data.Email, data.Feedback)
	err := helper.SendEmail(emailUser, "New Feedback Form Submission from BPIT Alumni Website", emailBody)
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
