package confirmAppointment

import (
	"log"
	"time"

	"github.com/ibrkhalil/doctory/internal/schema"
)

func NotifyDoctorOfConfirmation(availability schema.DoctorAvailability) {
	log.Print("Availability confirmed at time " + availability.Time.Format(time.RFC822Z))
}

func NotifyPatientOfConfirmation(availability schema.DoctorAvailability) {
	log.Print("Doctor confirmed appointment at time " + availability.Time.Format(time.RFC822Z))
}
