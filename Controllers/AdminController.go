package controllers

import (
	"encoding/json"
	"fmt"
	models "my-go-backend/Models"
	database "my-go-backend/config"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type AlumniResponse struct {
	AttendID      int64
	EventID       int64
	AlumniID      int64
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
		Select("alumni_profiles.first_name, alumni_profiles.last_name, alumni_profiles.alumni_id, alumni_attendings.position, alumni_attendings.attend_id,events.event_id, events.title, events.event_date_time, events.location").
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

func CreateAdminNetworking(w http.ResponseWriter, r *http.Request) {
	var input struct {
		AlumniID      int64
		Position      string
		Title         string
		Description   string
		EventType     string
		ModeOfEvent   string
		Location      string
		EventDateTime time.Time
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create the Event
	event := models.Event{
		Title:         input.Title,
		Description:   input.Description,
		EventType:     input.EventType,
		ModeOfEvent:   input.ModeOfEvent,
		Location:      input.Location,
		EventDateTime: input.EventDateTime,
	}

	// Create AlumniAttending
	alumniAttending := models.AlumniAttending{
		AlumniID: input.AlumniID,
		Position: input.Position,
	}

	// Use a transaction to ensure both inserts are successful
	tx := database.DB.Begin()
	if err := tx.Create(&event).Error; err != nil {
		tx.Rollback()
		http.Error(w, "Failed to create event", http.StatusInternalServerError)
		return
	}

	alumniAttending.EventID = event.EventID
	if err := tx.Create(&alumniAttending).Error; err != nil {
		tx.Rollback()
		http.Error(w, "Failed to create alumni attending", http.StatusInternalServerError)
		return
	}

	tx.Commit()
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Event and AlumniAttending created successfully")
}

func GetAdminAlumniAttendingByAlumniID(w http.ResponseWriter, r *http.Request) {
	alumniIDStr := mux.Vars(r)["alumni_id"]

	var alumniResponses []AlumniResponse

	err := database.DB.Table("alumni_profiles").
		Select("alumni_profiles.first_name, alumni_profiles.last_name, alumni_profiles.alumni_id, alumni_attendings.position, alumni_attendings.attend_id,events.event_id, events.title, events.event_date_time, events.location").
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
	var input struct {
		AlumniID      int64
		Position      string
		Title         string
		Description   string
		EventType     string
		ModeOfEvent   string
		Location      string
		EventDateTime time.Time
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get the event ID from the URL parameters (assuming you pass it in the URL)
	vars := mux.Vars(r)
	eventID := vars["id"]

	var event models.Event
	if err := database.DB.First(&event, eventID).Error; err != nil {
		http.Error(w, "Event not found", http.StatusNotFound)
		return
	}

	// Update fields if they are present in the request body
	if input.Title != "" {
		event.Title = input.Title
	}
	if input.Description != "" {
		event.Description = input.Description
	}
	if input.EventType != "" {
		event.EventType = input.EventType
	}
	if input.ModeOfEvent != "" {
		event.ModeOfEvent = input.ModeOfEvent
	}
	if input.Location != "" {
		event.Location = input.Location
	}
	if !input.EventDateTime.IsZero() {
		event.EventDateTime = input.EventDateTime
	}

	// Update AlumniAttending record if necessary
	var alumniAttending models.AlumniAttending
	if err := database.DB.Where("event_id = ? AND alumni_id = ?", event.EventID, input.AlumniID).First(&alumniAttending).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			alumniAttending = models.AlumniAttending{
				EventID:  event.EventID,
				AlumniID: input.AlumniID,
				Position: input.Position,
			}
			if err := database.DB.Create(&alumniAttending).Error; err != nil {
				http.Error(w, "Failed to create alumni attending", http.StatusInternalServerError)
				return
			}
		} else {
			http.Error(w, "Failed to find alumni attending", http.StatusInternalServerError)
			return
		}
	} else {
		if input.Position != "" {
			alumniAttending.Position = input.Position
		}
		if err := database.DB.Save(&alumniAttending).Error; err != nil {
			http.Error(w, "Failed to update alumni attending", http.StatusInternalServerError)
			return
		}
	}

	if err := database.DB.Save(&event).Error; err != nil {
		http.Error(w, "Failed to update event", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Event and AlumniAttending updated successfully")
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

// Added By Me

func GetAlumniAchievements(w http.ResponseWriter, r *http.Request) {
	var data []struct {
		AchievementID int64
		AlumniID     int64
		FirstName    string
		LastName     string
		Branch       string
		BatchYear    int64
		Title        string
		Description  string
		DateAchieved time.Time
	}
	if err := database.DB.Table("achievements").
		Select("alumni_profiles.first_name, alumni_profiles.last_name, alumni_profiles.branch, alumni_profiles.batch_year, achievements.*").
		Joins("JOIN alumni_profiles ON alumni_profiles.alumni_id = achievements.alumni_id").
		Scan(&data).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
