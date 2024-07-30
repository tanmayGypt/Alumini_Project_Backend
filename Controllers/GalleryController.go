package controllers

import (
	"encoding/json"
	"fmt"
	models "my-go-backend/Models"
	database "my-go-backend/config"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Album struct {
	ImageTitle string `json:"image_title"`
	ImageLink  string `json:"image_link"`
	ImageCount int64  `json:"image_count"`
}

func AddImage(w http.ResponseWriter, r *http.Request) {
	var img models.Gallery
	if err := json.NewDecoder(r.Body).Decode(&img); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if table exists or create it if it doesn't
	if !database.DB.Migrator().HasTable(&models.Gallery{}) {
		if err := database.DB.AutoMigrate(&models.Gallery{}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if result := database.DB.Create(&img); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(img)
}

func GetAllImages(w http.ResponseWriter, r *http.Request) {
	// Check if table exists or create it if it doesn't
	if !database.DB.Migrator().HasTable(&models.Gallery{}) {
		if err := database.DB.AutoMigrate(&models.Gallery{}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	var img []models.Gallery
	if result := database.DB.Find(&img); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(img)
}

func UpdateImage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var img models.Gallery
	if result := database.DB.First(&img, id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&img); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	database.DB.Save(&img)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(img)
}

func DeleteImage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if result := database.DB.Delete(&models.Gallery{}, id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func GetAlbums(w http.ResponseWriter, r *http.Request) {
	var albums []Album
	err := database.DB.Raw(`
	WITH RandomizedImages AS (
			SELECT image_title, image_link,
			       ROW_NUMBER() OVER (PARTITION BY image_title ORDER BY RANDOM()) as rn
			FROM galleries
		),
		CountedImages AS (
			SELECT image_title, COUNT(*) as image_count
			FROM galleries
			GROUP BY image_title
		)
		SELECT r.image_title, r.image_link, c.image_count
		FROM RandomizedImages r
		JOIN CountedImages c ON r.image_title = c.image_title
		WHERE r.rn = 1
`).Scan(&albums).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(albums)
}

func GetImagesByTitle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	title := params["title"]
	fmt.Println("Received Title:", title)
	var images []models.Gallery
	if result := database.DB.Where("image_title = ?", title).Find(&images).Error; result != nil {
		http.Error(w, result.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(images)
}
