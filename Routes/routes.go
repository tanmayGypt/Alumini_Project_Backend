package routes

import (
	"encoding/json"
	models "my-go-backend/Models"
	database "my-go-backend/config"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func InitializeRoutes(router *mux.Router) {
	router.HandleFunc("/alumni", createAlumniProfile).Methods("POST")
	router.HandleFunc("/alumni", getAlumniProfiles).Methods("GET")
	router.HandleFunc("/alumni/{id}", getAlumniProfileByID).Methods("GET")
	router.HandleFunc("/alumni/{id}", updateAlumniProfile).Methods("PUT")
	router.HandleFunc("/alumni/{id}", deleteAlumniProfile).Methods("DELETE")
	router.HandleFunc("/event", createNewEvent).Methods("POST")
	router.HandleFunc("/achievement", addAchievements).Methods("POST")
	router.HandleFunc("/professionalInfo", addProfessionalInfo).Methods("POST")
	router.HandleFunc("/event", getEvents).Methods("GET")
}

func createAlumniProfile(w http.ResponseWriter, r *http.Request) {
	var alumni models.AlumniProfile
	if err := json.NewDecoder(r.Body).Decode(&alumni); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if table exists or create it if it doesn't
	if !database.DB.Migrator().HasTable(&models.AlumniProfile{}) {
		if err := database.DB.AutoMigrate(&models.AlumniProfile{}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if result := database.DB.Create(&alumni); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(alumni)
}

func getAlumniProfiles(w http.ResponseWriter, r *http.Request) {
	// Check if table exists or create it if it doesn't
	if !database.DB.Migrator().HasTable(&models.AlumniProfile{}) {
		if err := database.DB.AutoMigrate(&models.AlumniProfile{}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	var alumni []models.AlumniProfile
	if result := database.DB.Find(&alumni); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(alumni)
}

func getAlumniProfileByID(w http.ResponseWriter, r *http.Request) {
	// Check if table exists or create it if it doesn't
	if !database.DB.Migrator().HasTable(&models.AlumniProfile{}) {
		if err := database.DB.AutoMigrate(&models.AlumniProfile{}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var alumni models.AlumniProfile
	if result := database.DB.First(&alumni, id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(alumni)
}

func updateAlumniProfile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var alumni models.AlumniProfile
	if result := database.DB.First(&alumni, id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&alumni); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	database.DB.Save(&alumni)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(alumni)
}

func deleteAlumniProfile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if result := database.DB.Delete(&models.AlumniProfile{}, id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func addAchievements(w http.ResponseWriter, r *http.Request) {
	var achievement models.Achievement
	if err := json.NewDecoder(r.Body).Decode(&achievement); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Parse dateAchieved string into time.Time
	dateStr := r.FormValue("dateAchieved") // Assuming dateAchieved is a string in "YYYY-MM-DD" format
	dateTime, err := time.Parse("2006-01-02T15:04:05Z07:00", dateStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	achievement.DateAchieved = dateTime

	// Check if table exists or create it if it doesn't
	if !database.DB.Migrator().HasTable(&models.Achievement{}) {
		if err := database.DB.AutoMigrate(&models.Achievement{}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// Create the achievement record
	if result := database.DB.Create(&achievement); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(achievement)
}

func addProfessionalInfo(w http.ResponseWriter, r *http.Request) {
	var Data models.ProfessionalInformation
	if err := json.NewDecoder(r.Body).Decode(&Data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if table exists or create it if it doesn't
	if !database.DB.Migrator().HasTable(&models.ProfessionalInformation{}) {
		if err := database.DB.AutoMigrate(&models.ProfessionalInformation{}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if result := database.DB.Create(&Data); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Data)
}

func createNewEvent(w http.ResponseWriter, r *http.Request) {
	var event models.Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if table exists or create it if it doesn't
	if !database.DB.Migrator().HasTable(&models.Event{}) {
		if err := database.DB.AutoMigrate(&models.Event{}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if result := database.DB.Create(&event); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(event)
}

func getEvents(w http.ResponseWriter, r *http.Request) {
	// Check if table exists or create it if it doesn't
	if !database.DB.Migrator().HasTable(&models.Event{}) {
		if err := database.DB.AutoMigrate(&models.Event{}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	var events []models.Event
	if result := database.DB.Find(&events); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(events)
}
