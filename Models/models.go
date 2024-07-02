package models

import (
	"time"

	"gorm.io/gorm"
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
	ProfilePicture   []byte
	// ProfessionalInformation []ProfessionalInformation `gorm:"foreignKey:AlumniID"`
	// Achievements            []Achievement             `gorm:"foreignKey:AlumniID"`
	// InterestsHobbies        []InterestHobby           `gorm:"foreignKey:AlumniID"`
	// AlumniAttending         []AlumniAttending         `gorm:"foreignKey:AlumniID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type ProfessionalInformation struct {
	ProfID      int64 `gorm:"primaryKey;autoIncrement"`
	AlumniID    int64
	CompanyName string
	Position    string
	Duration    string
	Alumni      AlumniProfile `gorm:"foreignKey:AlumniID;references:AlumniID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type Achievement struct {
	AchievementID int64 `gorm:"primaryKey;autoIncrement"`
	AlumniID      int64
	Title         string
	Description   string
	DateAchieved  time.Time
	Alumni        AlumniProfile `gorm:"foreignKey:AlumniID;references:AlumniID"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type InterestHobby struct {
	InterestID    int64 `gorm:"primaryKey;autoIncrement"`
	AlumniID      int64
	InterestHobby string
	Alumni        AlumniProfile `gorm:"foreignKey:AlumniID;references:AlumniID"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
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
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type AlumniAttending struct {
	AttendID  int64 `gorm:"primaryKey;autoIncrement"`
	EventID   int64
	AlumniID  int64
	Event     Event         `gorm:"foreignKey:EventID;references:EventID"`
	Alumni    AlumniProfile `gorm:"foreignKey:AlumniID;references:AlumniID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
