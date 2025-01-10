package confirmAppointment

import (
	"log"
	"time"
)

func NotifyDoctorOfConfirmation(appointmentTime time.Time) {
	log.Print("Availability confirmed at time " + appointmentTime.Format(time.RFC822Z))
}

func NotifyPatientOfConfirmation(appointmentTime time.Time) {
	log.Print("Doctor confirmed appointment at time " + appointmentTime.Format(time.RFC822Z))
}

func ConfirmAppointmentById(appointmentTime time.Time) {
	NotifyDoctorOfConfirmation(appointmentTime)
	NotifyPatientOfConfirmation(appointmentTime)
}
