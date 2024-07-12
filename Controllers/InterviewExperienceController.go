package controllers

import (
	"encoding/json"
	models "my-go-backend/Models"
	database "my-go-backend/config"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func AddInterviewExperience(w http.ResponseWriter, r *http.Request) {
	var interviewExperience models.InterviewExperience
	if err := json.NewDecoder(r.Body).Decode(&interviewExperience); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !database.DB.Migrator().HasTable(&models.InterviewExperience{}) {
		if err := database.DB.AutoMigrate(&models.InterviewExperience{}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if result := database.DB.Create(&interviewExperience); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(interviewExperience)
}

func UpdateInterviewExperience(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var interviewExperience models.InterviewExperience
	if result := database.DB.First(&interviewExperience, id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&interviewExperience); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	database.DB.Save(&interviewExperience)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(interviewExperience)
}

func DeleteInterviewExperience(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if result := database.DB.Delete(&models.InterviewExperience{}, id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func GetAllInterviewExperienceByAlumniID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var interviewExperiences []models.InterviewExperience
	if result := database.DB.Where("alumni_id = ?", id).Find(&interviewExperiences); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(interviewExperiences)
}

func GetInterviewExperiences(w http.ResponseWriter, r *http.Request) {
	// Check if table exists or create it if it doesn't
	if !database.DB.Migrator().HasTable(&models.InterviewExperience{}) {
		if err := database.DB.AutoMigrate(&models.InterviewExperience{}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	var res []models.InterviewExperience
	if result := database.DB.Find(&res); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
