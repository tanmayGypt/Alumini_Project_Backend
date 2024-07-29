package controllers

import (
	"encoding/json"
	database "my-go-backend/config"
	"net/http"
	"time"
)

type AlumniResponse struct {
	FirstName     string
	LastName      string
	Position      string
	Title         string
	EventDateTime time.Time
	Location      string
}

type AlumniDirectory struct {
	FirstName      string
	LastName       string
	Email          string
	Branch         string
	MobileNo       string
	CurrentCompany string
}

func GetAlumniAttending(w http.ResponseWriter, r *http.Request) {
	var alumniResponses []AlumniResponse

	err := database.DB.Table("alumni_profiles").
		Select("alumni_profiles.first_name, alumni_profiles.last_name, alumni_attendings.position, events.title, events.event_date_time, events.location").
		Joins("JOIN alumni_attendings ON alumni_profiles.alumni_id = alumni_attendings.alumni_id").
		Joins("JOIN events ON alumni_attendings.event_id = events.event_id").
		Where("alumni_profiles.status = ?", "alumni").
		Scan(&alumniResponses).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(alumniResponses)
}
