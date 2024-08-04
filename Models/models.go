package models

import (
	"time"
)

type AlumniProfile struct {
	AlumniID                int64 `gorm:"primaryKey;autoIncrement"`
	FirstName               string
	LastName                string
	Fathername              string
	Password                string
	Status                  string `json:"status" validate:"required,oneof=student alumni"`
	Branch                  string
	BatchYear               int64
	MobileNo                string `gorm:"unique"`
	Email                   string `gorm:"unique"`
	EnrollmentNo            string `gorm:"unique"`
	IsVerified              bool
	IsApproved              bool
	Code                    string
	ExpiresAt               time.Time
	Tenth                   string
	Xllth                   string
	Degree                  string
	GithubProfile           *string
	LeetCodeProfile         *string
	LinkedInProfile         *string
	CodeforceProfile        *string
	CodeChefProfile         *string
	InstagramProfile        *string
	TwitterProfile          *string
	GeeksForGeeksProfile    *string
	CodingNinjasProfile     *string
	ProfilePicture          string
	ProfessionalInformation []ProfessionalInformation `gorm:"foreignKey:AlumniID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Achievements            []Achievement             `gorm:"foreignKey:AlumniID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	InterestsHobbies        []InterestHobby           `gorm:"foreignKey:AlumniID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	AlumniAttending         []AlumniAttending         `gorm:"foreignKey:AlumniID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	InterviewExperience     []InterviewExperience     `gorm:"foreignKey:AlumniID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt               time.Time
	UpdatedAt               time.Time
}

type ProfessionalInformation struct {
	ProfID      int64 `gorm:"primaryKey;autoIncrement"`
	AlumniID    int64 `gorm:"index"`
	CompanyName string
	Position    string
	StartDate   time.Time
	EndDate     time.Time
	IsApproved  bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Achievement struct {
	AchievementID int64 `gorm:"primaryKey;autoIncrement"`
	AlumniID      int64 `gorm:"index"`
	Title         string
	Description   string
	DateAchieved  time.Time
	IsApproved    bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type InterestHobby struct {
	InterestID    int64 `gorm:"primaryKey;autoIncrement"`
	AlumniID      int64 `gorm:"index"`
	InterestHobby string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Event struct {
	EventID         int64 `gorm:"primaryKey;autoIncrement"`
	Title           string
	Description     string
	EventType       string
	ModeOfEvent     string
	Location        string
	EventDateTime   time.Time
	AlumniAttending []AlumniAttending `gorm:"foreignKey:EventID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type AlumniAttending struct {
	AttendID  int64 `gorm:"primaryKey;autoIncrement"`
	EventID   int64 `gorm:"index"`
	AlumniID  int64 `gorm:"index"`
	Position  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type InterviewExperience struct {
	ExperienceID  int64 `gorm:"primaryKey;autoIncrement"`
	AlumniID      int64
	CompanyName   string
	JobTitle      string
	Description   string
	InterviewDate time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type ResetPassword struct {
	ResetID   int64     `gorm:"primaryKey;autoIncrement"`
	Email     string    `gorm:"not null;unique"`
	Code      string    `gorm:"not null"`
	ExpiresAt time.Time `gorm:"not null"`
}

type Gallery struct {
	ImageID          int64  `gorm:"primaryKey;autoIncrement"`
	ImageLink        string `gorm:"not null;unique"`
	ImageTitle       string `gorm:"not null;"`
	ImageDescription string
}
