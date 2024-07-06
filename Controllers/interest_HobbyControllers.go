package controllers

import (
	"encoding/json"
	models "my-go-backend/Models"
	database "my-go-backend/config"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func AddInterestHobby(w http.ResponseWriter, r *http.Request) {
	var interestHobby models.InterestHobby
	if err := json.NewDecoder(r.Body).Decode(&interestHobby); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !database.DB.Migrator().HasTable(&models.InterestHobby{}) {
		if err := database.DB.AutoMigrate(&models.InterestHobby{}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if result := database.DB.Create(&interestHobby); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(interestHobby)
}

func UpdateInterestHobby(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var interestHobby models.InterestHobby
	if result := database.DB.First(&interestHobby, id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&interestHobby); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	database.DB.Save(&interestHobby)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(interestHobby)
}

func DeleteInterestHobby(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if result := database.DB.Delete(&models.InterestHobby{}, id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func GetAllInterestHobbiesByAlumniID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var interestHobbies []models.InterestHobby
	if result := database.DB.Where("alumni_id = ?", id).Find(&interestHobbies); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(interestHobbies)
}
