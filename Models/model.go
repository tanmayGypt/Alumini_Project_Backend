package models

import "time"

type AlumniProfile struct {
	AlumniId             int
	FirstName            string
	LastName             string
	Branch               string
	Batch                int
	EnrollmentNo         string
	MobileNo             string
	Email                string
	Percentage_In_10th   string
	Percentage_In_12th   string
	HighestQualification string
	GithubProfile        *string
	LinkedInProfile      *string
	LeetCodeProfile      *string
	CodeForceProfile     *string
	CodeChefProfile      *string
	ProfilePicture       *[]byte
}

type ProfessionalInformation struct {
	ProfID      int
	AlumniId    int
	CompanyName string
	Designation string
	Duration    string
}

type Achievements struct {
	AchievementsID int
	AlumniId       int
	Title          string
	Description    *string
	DateAchieved   time.Time
}

type InterestsHobbies struct {
	InterestsHobbiesID int
	AlumniId           int
	Title              string
}

type Events struct {
	EventID     int
	Title       string
	Description *string
	DateOfEvent time.Time
	EventType   string
	ModeOfEvent string
	Location    string
}
type AlumniAttending struct {
	AttendID int
	EventID  int
	AlumniId int
}
