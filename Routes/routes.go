package routes

import (
	controllers "my-go-backend/Controllers"
	middleware "my-go-backend/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeRoutes(router *mux.Router) {
	// Authentication routes
	router.HandleFunc("/register", controllers.Register).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/signup", controllers.Signup).Methods("POST")
	registerHandler := http.HandlerFunc(controllers.Register)
	router.Handle("/verifyOTP", middleware.OTPVerify(registerHandler)).Methods("POST")

	// Alumni routes
	alumniRouter := router.PathPrefix("/alumni").Subrouter()
	alumniRouter.Use(middleware.JWTVerify)
	alumniRouter.HandleFunc("", controllers.CreateAlumniProfile).Methods("POST")
	alumniRouter.HandleFunc("", controllers.GetAlumniProfiles).Methods("GET")
	alumniRouter.HandleFunc("/{id}", controllers.GetAlumniProfileByID).Methods("GET")
	alumniRouter.HandleFunc("/{id}", controllers.UpdateAlumniProfile).Methods("PUT")
	alumniRouter.HandleFunc("/{id}", controllers.DeleteAlumniProfile).Methods("DELETE")

	// Event routes
	eventRouter := router.PathPrefix("/event").Subrouter()
	eventRouter.Use(middleware.JWTVerify)
	eventRouter.HandleFunc("", controllers.CreateNewEvent).Methods("POST")
	eventRouter.HandleFunc("", controllers.GetEvents).Methods("GET")
	eventRouter.HandleFunc("/{id}", controllers.UpdateEvent).Methods("PUT")
	eventRouter.HandleFunc("/{id}", controllers.DeleteEvent).Methods("DELETE")
	eventRouter.HandleFunc("/{id}", controllers.GetEventByID).Methods("GET")

	// Professional Info routes
	professionalInfoRouter := router.PathPrefix("/professionalInfo").Subrouter()
	professionalInfoRouter.Use(middleware.JWTVerify)
	professionalInfoRouter.HandleFunc("", controllers.AddProfessionalInfo).Methods("POST")
	professionalInfoRouter.HandleFunc("/{id}", controllers.UpdateProfessionalInfo).Methods("PUT")
	professionalInfoRouter.HandleFunc("/{id}", controllers.DeleteProfessionalInfo).Methods("DELETE")
	professionalInfoRouter.HandleFunc("/{id}", controllers.GetAllProfessionalInfo).Methods("GET")
	professionalInfoRouter.HandleFunc("", controllers.GetProfessionalInfos).Methods("GET")

	// Achievements routes
	achievementRouter := router.PathPrefix("/achievement").Subrouter()
	achievementRouter.Use(middleware.JWTVerify)
	achievementRouter.HandleFunc("", controllers.AddAchievements).Methods("POST")
	achievementRouter.HandleFunc("/{id}", controllers.UpdateAchievementInfo).Methods("PUT")
	achievementRouter.HandleFunc("/{id}", controllers.GetAllAchievementByAlumniID).Methods("GET")
	achievementRouter.HandleFunc("/{id}", controllers.DeleteAchievement).Methods("DELETE")
	achievementRouter.HandleFunc("", controllers.GetAchievements).Methods("GET")

	// Interest/Hobbies routes
	interestHobbiesRouter := router.PathPrefix("/interesthobbies").Subrouter()
	interestHobbiesRouter.Use(middleware.JWTVerify)
	interestHobbiesRouter.HandleFunc("", controllers.AddInterestHobby).Methods("POST")
	interestHobbiesRouter.HandleFunc("/{id}", controllers.UpdateInterestHobby).Methods("PUT")
	interestHobbiesRouter.HandleFunc("/{id}", controllers.DeleteInterestHobby).Methods("DELETE")
	interestHobbiesRouter.HandleFunc("/alumni/{id}", controllers.GetAllInterestHobbiesByAlumniID).Methods("GET")

	// Interview Experience routes
	interviewExperienceRouter := router.PathPrefix("/interviewexperiences").Subrouter()
	interviewExperienceRouter.Use(middleware.JWTVerify)
	interviewExperienceRouter.HandleFunc("", controllers.AddInterviewExperience).Methods("POST")
	interviewExperienceRouter.HandleFunc("/{id}", controllers.UpdateInterviewExperience).Methods("PUT")
	interviewExperienceRouter.HandleFunc("/{id}", controllers.DeleteInterviewExperience).Methods("DELETE")
	interviewExperienceRouter.HandleFunc("/alumni/{id}", controllers.GetAllInterviewExperienceByAlumniID).Methods("GET")
	interviewExperienceRouter.HandleFunc("", controllers.GetInterviewExperiences).Methods("GET")

	// Alumni Attending routes
	alumniAttendingRouter := router.PathPrefix("/alumniattending").Subrouter()
	alumniAttendingRouter.Use(middleware.JWTVerify)
	alumniAttendingRouter.HandleFunc("", controllers.AddAlumniForEvent).Methods("POST")
	alumniAttendingRouter.HandleFunc("/{id}", controllers.UpdateAlumniAttending).Methods("PUT")
	alumniAttendingRouter.HandleFunc("/{id}", controllers.DeleteAlumniAttending).Methods("DELETE")
	alumniAttendingRouter.HandleFunc("/event/{id}", controllers.GetAlumniByEventID).Methods("GET")
	alumniAttendingRouter.HandleFunc("/alumni/{id}", controllers.GetEventsByAlumniID).Methods("GET")

	// testing route
	router.HandleFunc("/test", controllers.SendEmail).Methods("POST")

}
