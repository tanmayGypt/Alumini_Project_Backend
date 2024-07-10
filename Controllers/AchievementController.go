package controllers

import (
	"encoding/json"
	models "my-go-backend/Models"
	database "my-go-backend/config"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func AddAchievements(w http.ResponseWriter, r *http.Request) {
	var achievement models.Achievement
	if err := json.NewDecoder(r.Body).Decode(&achievement); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

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

func UpdateAchievementInfo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var info models.Achievement
	if result := database.DB.First(&info, id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&info); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	database.DB.Save(&info)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(info)
}
func GetAllAchievementByAlumniID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var achievementInfos []models.Achievement
	result := database.DB.Where("alumni_id = ?", id).Find(&achievementInfos)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(achievementInfos)
}

func DeleteAchievement(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if result := database.DB.Delete(&models.Achievement{}, id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
func GetAchievements(w http.ResponseWriter, r *http.Request) {
	// Check if table exists or create it if it doesn't
	if !database.DB.Migrator().HasTable(&models.Achievement{}) {
		if err := database.DB.AutoMigrate(&models.Achievement{}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	var achievements []models.Achievement
	if result := database.DB.Find(&achievements); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(achievements)
}
