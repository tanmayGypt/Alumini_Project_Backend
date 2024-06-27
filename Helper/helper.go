package helper

import (
	"context"
	"fmt"
	"log"
	models "my-go-backend/Models"
	"my-go-backend/prisma/db"
)

func CreateAlumni(ctx context.Context, client *db.PrismaClient, alumni models.AlumniProfile) {
	createdAlumni, err := client.alumni_profile.CreateOne(
		db.alumni_profile.first_name.Set(alumni.FirstName),
		db.alumni_profile.last_name.Set(alumni.LastName),
		db.alumni_profile.email.Set(alumni.Email),
		db.alumni_profile.mobile_no.Set(alumni.MobileNo),
		db.alumni_profile.branch.Set(alumni.Branch),
		db.alumni_profile.batch_year.Set(alumni.Batch),
		db.alumni_profile.enrollment_no.Set(alumni.EnrollmentNo),
		db.alumini_profile.tenth.Set(alumni.Percentage_In_10th),
		db.alumini_profile.xllth.Set(alumni.Percentage_In_12th),
		db.alumini_profile.degree.Set(alumni.HighestQualification),
		db.alumini_profile.github_profile.Set(alumni.GithubProfile),
  		db.alumini_profile.leet_code_profile.Set(alumni.LeetCodeProfile),
  		db.alumini_profile.linked_in_profile.Set(alumni.LinkedInProfile),
  		db.alumini_profile.codeforce_profile.Set(alumni.CodeForceProfile),
  		db.alumini_profile.code_chef_profile.Set(alumni.CodeChefProfile),
	).Exec(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(createdAlumni)
}

func CreateEvent(ctx context.Context, client *db.PrismaClient, event models.Events) {
	createdEvent, err := client.Events.CreateOne(
		db.events.title.Set(event.Title),
		db.events.description.Set(*event.Description),
		db.events.event_date_time.Set(event.DateOfEvent),
		db.events.mode_of_event.Set(event.ModeOfEvent),
		db.events.location.Set(event.Location),
		db.events.event_type.Set(event.EventType),
	).Exec(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(createdEvent)

}
