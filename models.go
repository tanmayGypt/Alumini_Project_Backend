package main

import (
	"time"

	"gorm.io/gorm"
)
type AlumniProfile struct {
	gorm.Model
	AlumniID               int                    `gorm:"primaryKey;autoIncrement"`
	FirstName              string
	LastName               string
	Branch                 string
	BatchYear              int
	MobileNo               string                 `gorm:"unique"`
	Email                  string                 `gorm:"unique"`
	EnrollmentNo           string                 `gorm:"unique"`
	Tenth                  string
	Xllth                  string
	Degree                 string
	GithubProfile          *string
	LeetCodeProfile        *string
	LinkedInProfile        *string
	CodeforceProfile       *string
	CodeChefProfile        *string
	ProfilePicture         []byte
	ProfessionalInformation []ProfessionalInformation `gorm:"foreignKey:AlumniID"`
	Achievements           []Achievement             `gorm:"foreignKey:AlumniID"`
	InterestsHobbies       []InterestHobby           `gorm:"foreignKey:AlumniID"`
	AlumniAttending        []AlumniAttending         `gorm:"foreignKey:AlumniID"`
}

type ProfessionalInformation struct {
	gorm.Model
	ProfID      int            `gorm:"primaryKey;autoIncrement"`
	AlumniID    int
	CompanyName string
	Position    string
	Duration    string
	Alumni      AlumniProfile  `gorm:"foreignKey:AlumniID;references:AlumniID"`
}

type Achievement struct {
	gorm.Model
	AchievementID int            `gorm:"primaryKey;autoIncrement"`
	AlumniID      int
	Title         string
	Description   string
	DateAchieved  time.Time
	Alumni        AlumniProfile  `gorm:"foreignKey:AlumniID;references:AlumniID"`
}

type InterestHobby struct {
	gorm.Model
	InterestID    int            `gorm:"primaryKey;autoIncrement"`
	AlumniID      int
	InterestHobby string
	Alumni        AlumniProfile  `gorm:"foreignKey:AlumniID;references:AlumniID"`
}

type Event struct {
	gorm.Model
	EventID       int            `gorm:"primaryKey;autoIncrement"`
	Title         string
	Description   string
	EventType     string
	ModeOfEvent   string
	Location      string
	EventDateTime time.Time
	AlumniAttending []AlumniAttending `gorm:"foreignKey:EventID"`
}

type AlumniAttending struct {
	gorm.Model
	AttendID      int            `gorm:"primaryKey;autoIncrement"`
	EventID       int
	AlumniID      int
	Event         Event          `gorm:"foreignKey:EventID;references:EventID"`
	Alumni        AlumniProfile  `gorm:"foreignKey:AlumniID;references:AlumniID"`
}
