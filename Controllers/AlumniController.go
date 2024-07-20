package controllers

import (
	"encoding/json"
	models "my-go-backend/Models"
	database "my-go-backend/config"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateAlumniProfile(w http.ResponseWriter, r *http.Request) {
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

func GetAlumniProfiles(w http.ResponseWriter, r *http.Request) {
	// Check if table exists or create it if it doesn't
	if !database.DB.Migrator().HasTable(&models.AlumniProfile{}) {
		if err := database.DB.AutoMigrate(&models.AlumniProfile{}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// Parse the status parameter from the query string
	status := r.URL.Query().Get("status")

	var alumni []models.AlumniProfile
	query := database.DB

	// Modify the query based on the status parameter
	if status == "student" {
		query = query.Where("status = ?", "student")
	} else if status == "alumni" {
		query = query.Where("status = ?", "alumni")
	}

	// Fetch the data
	if result := query.Find(&alumni); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(alumni)
}

func GetAlumniProfileByID(w http.ResponseWriter, r *http.Request) {
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

func UpdateAlumniProfile(w http.ResponseWriter, r *http.Request) {
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

func DeleteAlumniProfile(w http.ResponseWriter, r *http.Request) {
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
