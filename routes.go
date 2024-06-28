package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func InitializeRoutes(router *mux.Router) {
	router.HandleFunc("/alumni", createAlumniProfile).Methods("POST")
	router.HandleFunc("/alumni", getAlumniProfiles).Methods("GET")
	router.HandleFunc("/alumni/{id}", getAlumniProfileByID).Methods("GET")
	router.HandleFunc("/alumni/{id}", updateAlumniProfile).Methods("PUT")
	router.HandleFunc("/alumni/{id}", deleteAlumniProfile).Methods("DELETE")
}

func createAlumniProfile(w http.ResponseWriter, r *http.Request) {
	var alumni AlumniProfile
	if err := json.NewDecoder(r.Body).Decode(&alumni); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if table exists or create it if it doesn't
	if !DB.Migrator().HasTable(&AlumniProfile{}) {
		if err := DB.AutoMigrate(&AlumniProfile{}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if result := DB.Create(&alumni); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(alumni)
}

func getAlumniProfiles(w http.ResponseWriter, r *http.Request) {
	// Check if table exists or create it if it doesn't
	if !DB.Migrator().HasTable(&AlumniProfile{}) {
		if err := DB.AutoMigrate(&AlumniProfile{}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	var alumni []AlumniProfile
	if result := DB.Find(&alumni); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(alumni)
}

func getAlumniProfileByID(w http.ResponseWriter, r *http.Request) {
	// Check if table exists or create it if it doesn't
	if !DB.Migrator().HasTable(&AlumniProfile{}) {
		if err := DB.AutoMigrate(&AlumniProfile{}); err != nil {
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

	var alumni AlumniProfile
	if result := DB.First(&alumni, id); result.Error != nil {
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
	var alumni AlumniProfile
	if result := DB.First(&alumni, id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&alumni); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	DB.Save(&alumni)
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
	if result := DB.Delete(&AlumniProfile{}, id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
