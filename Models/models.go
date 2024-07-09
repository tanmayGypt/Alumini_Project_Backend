package models

import (
	"time"
)

type AlumniProfile struct {
	AlumniID         int64 `gorm:"primaryKey;autoIncrement"`
	FirstName        string
	LastName         string
	Branch           string
	BatchYear        int64
	MobileNo         string `gorm:"unique"`
	Email            string `gorm:"unique"`
	EnrollmentNo     string `gorm:"unique"`
	Tenth            string
	Xllth            string
	Degree           string
	GithubProfile    *string
	LeetCodeProfile  *string
	LinkedInProfile  *string
	CodeforceProfile *string
	CodeChefProfile  *string
	InstagramProfile *string
	TwitterProfile   *string
	ProfilePicture   string
	// ProfessionalInformation []ProfessionalInformation `gorm:"foreignKey:AlumniID"`
	// Achievements            []Achievement             `gorm:"foreignKey:AlumniID"`
	// InterestsHobbies        []InterestHobby           `gorm:"foreignKey:AlumniID"`
	// AlumniAttending         []AlumniAttending         `gorm:"foreignKey:AlumniID"`
	// InterviewExperiences    []InterviewExperience     `gorm:"foreignKey:AlumniID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ProfessionalInformation struct {
	ProfID      int64 `gorm:"primaryKey;autoIncrement"`
	AlumniID    int64 `gorm:"index"`
	CompanyName string
	Position    string
	Duration    string
	Alumni      AlumniProfile `gorm:"foreignKey:AlumniID;references:AlumniID;constraint:OnDelete:CASCADE"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Achievement struct {
	AchievementID int64 `gorm:"primaryKey;autoIncrement"`
	AlumniID      int64 `gorm:"index"`
	Title         string
	Description   string
	DateAchieved  time.Time
	Alumni        AlumniProfile `gorm:"foreignKey:AlumniID;references:AlumniID;constraint:OnDelete:CASCADE"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type InterestHobby struct {
	InterestID    int64 `gorm:"primaryKey;autoIncrement"`
	AlumniID      int64 `gorm:"index"`
	InterestHobby string
	Alumni        AlumniProfile `gorm:"foreignKey:AlumniID;references:AlumniID;constraint:OnDelete:CASCADE"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Event struct {
	EventID       int64 `gorm:"primaryKey;autoIncrement"`
	Title         string
	Description   string
	EventType     string
	ModeOfEvent   string
	Location      string
	EventDateTime time.Time
	// AlumniAttending []AlumniAttending `gorm:"foreignKey:EventID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AlumniAttending struct {
	AttendID  int64         `gorm:"primaryKey;autoIncrement"`
	EventID   int64         `gorm:"index"`
	AlumniID  int64         `gorm:"index"`
	Event     Event         `gorm:"foreignKey:EventID;references:EventID;constraint:OnDelete:CASCADE"`
	Alumni    AlumniProfile `gorm:"foreignKey:AlumniID;references:AlumniID;constraint:OnDelete:CASCADE"`
	CreatedAt time.Time
}

type InterviewExperience struct {
	ExperienceID  int64 `gorm:"primaryKey;autoIncrement"`
	AlumniID      int64
	CompanyName   string
	JobTitle      string
	Description   string
	InterviewDate time.Time
	Alumni        AlumniProfile `gorm:"foreignKey:AlumniID;references:AlumniID"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
