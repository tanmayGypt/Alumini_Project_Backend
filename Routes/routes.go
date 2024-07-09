package routes

import (
	controllers "my-go-backend/Controllers"

	"github.com/gorilla/mux"
)

func InitializeRoutes(router *mux.Router) {
	// routes for Alumni Table
	router.HandleFunc("/alumni", controllers.CreateAlumniProfile).Methods("POST")
	router.HandleFunc("/alumni", controllers.GetAlumniProfiles).Methods("GET")
	router.HandleFunc("/alumni/{id}", controllers.GetAlumniProfileByID).Methods("GET")
	router.HandleFunc("/alumni/{id}", controllers.UpdateAlumniProfile).Methods("PUT")
	router.HandleFunc("/alumni/{id}", controllers.DeleteAlumniProfile).Methods("DELETE")
	// routes for event Table
	router.HandleFunc("/event", controllers.CreateNewEvent).Methods("POST")
	router.HandleFunc("/event", controllers.GetEvents).Methods("GET")
	router.HandleFunc("/event/{id}", controllers.UpdateEvent).Methods("PUT")
	router.HandleFunc("/event/{id}", controllers.DeleteEvent).Methods("DELETE")
	router.HandleFunc("/event/{id}", controllers.GetEventByID).Methods("GET")
	// routes for professionalInfo Table
	router.HandleFunc("/professionalInfo", controllers.AddProfessionalInfo).Methods("POST")
	router.HandleFunc("/professionalInfo/{id}", controllers.UpdateProfessionalInfo).Methods("PUT")
	router.HandleFunc("/professionalInfo/{id}", controllers.DeleteProfessionalInfo).Methods("DELETE")
	router.HandleFunc("/professionalInfo/{id}", controllers.GetAllProfessionalInfo).Methods("GET")
	// routes for achievements table
	router.HandleFunc("/achievement", controllers.AddAchievements).Methods("POST")
	router.HandleFunc("/achievement/{id}", controllers.UpdateAchievementInfo).Methods("PUT")
	router.HandleFunc("/achievement/{id}", controllers.GetAllAchievementByAlumniID).Methods("GET")
	router.HandleFunc("/achievement/{id}", controllers.DeleteAchievement).Methods("DELETE")
	// routes for interest/hobbies table
	router.HandleFunc("/interesthobbies", controllers.AddInterestHobby).Methods("POST")
	router.HandleFunc("/interesthobbies/{id}", controllers.UpdateInterestHobby).Methods("PUT")
	router.HandleFunc("/interesthobbies/{id}", controllers.DeleteInterestHobby).Methods("DELETE")
	router.HandleFunc("/interesthobbies/alumni/{id}", controllers.GetAllInterestHobbiesByAlumniID).Methods("GET")
	// routes for interviewexperience table
	router.HandleFunc("/interviewexperiences", controllers.AddInterviewExperience).Methods("POST")
	router.HandleFunc("/interviewexperiences/{id}", controllers.UpdateInterviewExperience).Methods("PUT")
	router.HandleFunc("/interviewexperiences/{id}", controllers.DeleteInterviewExperience).Methods("DELETE")
	router.HandleFunc("/interviewexperiences/alumni/{id}", controllers.GetAllInterviewExperienceByAlumniID).Methods("GET")
	// routes for AlumniAttending Table
	router.HandleFunc("/alumniattending", controllers.AddAlumniForEvent).Methods("POST")
	router.HandleFunc("/alumniattending/{id}", controllers.UpdateAlumniAttending).Methods("PUT")
	router.HandleFunc("/alumniattending/{id}", controllers.DeleteAlumniAttending).Methods("DELETE")
	// for testing Authentication -->   Some Error need to handle for now
	router.HandleFunc("/login",controllers.HandleMicrosoftLogin)
	router.HandleFunc("/callback",controllers.HandleMicrosoftCallback)
}
