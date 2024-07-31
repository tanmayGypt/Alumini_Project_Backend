package controllers

import (
	"encoding/json"
	"fmt"
	models "my-go-backend/Models"
	database "my-go-backend/config"
	"net/http"
	"time"

	"github.com/gorilla/mux"
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

func CreateAdminAlumniAttending(w http.ResponseWriter, r *http.Request) {
	var input struct {
		AlumniID int64  `json:"alumni_id"`
		EventID  int64  `json:"event_id"`
		Position string `json:"position"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	err := database.DB.Table("alumni_attendings").Create(&models.AlumniAttending{
		AlumniID:  input.AlumniID,
		EventID:   input.EventID,
		Position:  input.Position,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Alumni attending record created successfully"})
}

func GetAdminAlumniAttendingByAlumniID(w http.ResponseWriter, r *http.Request) {
	alumniIDStr := mux.Vars(r)["alumni_id"]

	var alumniResponses []struct {
		FirstName     string    `json:"first_name"`
		LastName      string    `json:"last_name"`
		Position      string    `json:"position"`
		Title         string    `json:"title"`
		EventDateTime time.Time `json:"event_date_time"`
		Location      string    `json:"location"`
	}

	err := database.DB.Table("alumni_profiles").
		Select("alumni_profiles.first_name, alumni_profiles.last_name, alumni_attendings.position, events.title, events.event_date_time, events.location").
		Joins("JOIN alumni_attendings ON alumni_profiles.alumni_id = alumni_attendings.alumni_id").
		Joins("JOIN events ON alumni_attendings.event_id = events.event_id").
		Where("alumni_profiles.alumni_id = ?", alumniIDStr).
		Scan(&alumniResponses).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(alumniResponses) == 0 {
		http.Error(w, "No attending records found for the given alumni ID", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(alumniResponses)
}

func GetAdminAlumniAttendingByID(w http.ResponseWriter, r *http.Request) {
	alumniIDStr := mux.Vars(r)["alumni_id"]
	eventIDStr := mux.Vars(r)["event_id"]

	var alumniAttending models.AlumniAttending

	err := database.DB.Table("alumni_attendings").
		Where("alumni_id = ? AND event_id = ?", alumniIDStr, eventIDStr).
		First(&alumniAttending).Error
	if err != nil {
		http.Error(w, "Alumni attending record not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(alumniAttending)
}

func UpdateAdminAlumniAttending(w http.ResponseWriter, r *http.Request) {
	alumniIDStr := mux.Vars(r)["alumni_id"]
	eventIDStr := mux.Vars(r)["event_id"]

	var input struct {
		Position string `json:"position"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	var alumniAttending models.AlumniAttending

	err := database.DB.Table("alumni_attendings").
		Where("alumni_id = ? AND event_id = ?", alumniIDStr, eventIDStr).
		First(&alumniAttending).Error
	if err != nil {
		http.Error(w, "Alumni attending record not found", http.StatusNotFound)
		return
	}

	err = database.DB.Model(&alumniAttending).Updates(models.AlumniAttending{
		Position:  input.Position,
		UpdatedAt: time.Now(),
	}).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Alumni attending record updated successfully"})
}

func DeleteAdminAlumniAttending(w http.ResponseWriter, r *http.Request) {
	alumniIDStr := mux.Vars(r)["alumni_id"]
	eventIDStr := mux.Vars(r)["event_id"]

	err := database.DB.Table("alumni_attendings").
		Where("alumni_id = ? AND event_id = ?", alumniIDStr, eventIDStr).
		Delete(&models.AlumniAttending{}).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Alumni attending record deleted successfully"})
}
