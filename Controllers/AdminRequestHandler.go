package controllers

import (
	"encoding/json"
	"fmt"
	models "my-go-backend/Models"
	database "my-go-backend/config"
	"net/http"
	"strings"
)

func GetAllRequests(w http.ResponseWriter, r *http.Request) {
	table := r.URL.Query().Get("table")

	var result interface{}

	switch strings.ToLower(table) {
	case "alumni":
		var alumniProfiles []models.AlumniProfile
		if err := database.DB.Debug().Where("is_approved = ? AND status = ?", false, "alumni").Find(&alumniProfiles).Error; err != nil {
			http.Error(w, "Error fetching student profiles", http.StatusInternalServerError)
			return
		}
		fmt.Println(alumniProfiles)
		result = alumniProfiles
	case "student":
		var studentprofiles []models.AlumniProfile
		if err := database.DB.Debug().Where("is_approved = ? AND status = ?", false, "student").Find(&studentprofiles).Error; err != nil {
			http.Error(w, "Error fetching student profiles", http.StatusInternalServerError)
			return
		}
		fmt.Println(studentprofiles)
		result = studentprofiles
	case "achievements":
		var data []struct {
			models.Achievement
			FirstName string
			LastName  string
			Branch    string
			BatchYear int64
			Email     string
		}

		if err := database.DB.Debug().Table("achievements").
			Select("achievements.*, alumni_profiles.first_name, alumni_profiles.last_name, alumni_profiles.branch, alumni_profiles.batch_year, alumni_profiles.email").
			Joins("JOIN alumni_profiles ON achievements.alumni_id = alumni_profiles.alumni_id").
			Where("achievements.is_approved = ?", false).
			Scan(&data).Error; err != nil {
			http.Error(w, "Error fetching achievements", http.StatusInternalServerError)
			return
		}
		fmt.Println(data)
		result = data
	case "professional_information":
		var data []struct {
			models.ProfessionalInformation
			FirstName string
			LastName  string
			Branch    string
			BatchYear int64
			Email     string
		}

		if err := database.DB.Debug().Table("professional_informations").
			Select("professional_informations.*, alumni_profiles.first_name, alumni_profiles.last_name, alumni_profiles.branch, alumni_profiles.batch_year, alumni_profiles.email").
			Joins("JOIN alumni_profiles ON professional_informations.alumni_id = alumni_profiles.alumni_id").
			Where("professional_informations.is_approved = ?", false).
			Scan(&data).Error; err != nil {
			http.Error(w, "Error fetching professional information", http.StatusInternalServerError)
			return
		}
		fmt.Println(data)
		result = data
	// Add more cases for other tables as needed
	default:
		http.Error(w, "Invalid table name", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)

}

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Type       string
		ID         int64
		IsApproved bool
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	var err error
	switch input.Type {
	case "achievement":
		if input.IsApproved {
			err = database.DB.Model(&models.Achievement{}).Where("id = ?", input.ID).
				Update("is_approved", input.IsApproved).Error
		} else {
			err = database.DB.Where("id = ?", input.ID).
				Delete(&models.Achievement{}).Error
		}
	case "professional":
		if input.IsApproved {
			err = database.DB.Model(&models.ProfessionalInformation{}).Where("id = ?", input.ID).
				Update("is_approved", input.IsApproved).Error
		} else {
			err = database.DB.Where("id = ?", input.ID).
				Delete(&models.ProfessionalInformation{}).Error
		}
	case "alumni":
		if input.IsApproved {
			err = database.DB.Model(&models.AlumniProfile{}).Where("alumni_id = ?", input.ID).
				Update("is_approved", input.IsApproved).Error
		} else {
			err = database.DB.Where("alumni_id = ?", input.ID).
				Delete(&models.AlumniProfile{}).Error
		}
	default:
		http.Error(w, "Unknown type", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Approval status updated successfully")
}
