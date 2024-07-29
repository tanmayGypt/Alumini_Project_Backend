package controllers

import (
	"encoding/json"
	models "my-go-backend/Models"
	database "my-go-backend/config"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func AddAlumniForEvent(w http.ResponseWriter, r *http.Request) {
	var Data models.AlumniAttending
	if err := json.NewDecoder(r.Body).Decode(&Data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if table exists or create it if it doesn't
	if !database.DB.Migrator().HasTable(&models.AlumniAttending{}) {
		if err := database.DB.AutoMigrate(&models.AlumniAttending{}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if result := database.DB.Create(&Data); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Data)
}

func UpdateAlumniAttending(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var info models.AlumniAttending
	if result := database.DB.First(&info, id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&info); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	database.DB.Save(&info)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(info)
}
func DeleteAlumniAttending(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if result := database.DB.Delete(&models.AlumniAttending{}, id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func GetAlumniByEventID(w http.ResponseWriter, r *http.Request) {
	// Check if table exists or create it if it doesn't
	if !database.DB.Migrator().HasTable(&models.AlumniAttending{}) {
		if err := database.DB.AutoMigrate(&models.AlumniAttending{}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	params := mux.Vars(r)
	eventID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid EventID", http.StatusBadRequest)
		return
	}

	var alumni []models.AlumniProfile
	err = database.DB.Model(&models.AlumniAttending{}).
		Select("alumni_profiles.*").
		Joins("JOIN alumni_profiles ON alumni_profiles.alumni_id = alumni_attendings.alumni_id").
		Where("alumni_attendings.event_id = ?", eventID).
		Find(&alumni).Error

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(alumni)
}

func GetEventsByAlumniID(w http.ResponseWriter, r *http.Request) {
	// Check if table exists or create it if it doesn't
	if !database.DB.Migrator().HasTable(&models.AlumniAttending{}) {
		if err := database.DB.AutoMigrate(&models.AlumniAttending{}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	params := mux.Vars(r)
	alumniID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid AlumniID", http.StatusBadRequest)
		return
	}

	var events []models.Event
	err = database.DB.Model(&models.AlumniAttending{}).
		Select("events.*").
		Joins("JOIN events ON events.event_id = alumni_attendings.event_id").
		Where("alumni_attendings.alumni_id = ?", alumniID).
		Find(&events).Error

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(events)
}
type AlumniResponse struct {
	FirstName string
	LastName  string
	Position  string
	Title string
}

func GetAlumniAttending(w http.ResponseWriter, r *http.Request) {
	var alumniResponses []AlumniResponse

	err := database.DB.Table("alumni_profiles").
		Select("alumni_profiles.first_name, alumni_profiles.last_name, alumni_attendings.position, events.title").
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