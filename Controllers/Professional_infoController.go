package controllers

import (
	"encoding/json"
	models "my-go-backend/Models"
	database "my-go-backend/config"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func AddProfessionalInfo(w http.ResponseWriter, r *http.Request) {
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

func UpdateProfessionalInfo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var info models.ProfessionalInformation
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
func DeleteProfessionalInfo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if result := database.DB.Delete(&models.ProfessionalInformation{}, id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func GetAllProfessionalInfo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var professionalInfos []models.ProfessionalInformation
	result := database.DB.Where("alumni_id = ?", id).Find(&professionalInfos)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(professionalInfos)
}

func GetProfessionalInfos(w http.ResponseWriter, r *http.Request) {
	// Check if table exists or create it if it doesn't
	if !database.DB.Migrator().HasTable(&models.ProfessionalInformation{}) {
		if err := database.DB.AutoMigrate(&models.ProfessionalInformation{}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	var data []models.ProfessionalInformation
	if result := database.DB.Find(&data); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
