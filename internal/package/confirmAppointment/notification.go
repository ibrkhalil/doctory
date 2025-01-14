package confirmAppointment

import (
	"log"

	"github.com/ibrkhalil/doctory/internal/schema"
)

func NotifyDoctorOfAppointmentBooking(appointmnetSlot schema.AppointmentSlot) error {
	log.Print("Patient " + appointmnetSlot.PatientName + " Has successfully booked an appointment at " + appointmnetSlot.ReservedAt.Format("2006-01-02 15:04:05") + " for an appointment at time: " + appointmnetSlot.StartingTime.Format("2006-01-02 15:04:05"))
	return nil
}

func NotifyPatientOfAppointmentBooking(appointmnetSlot schema.AppointmentSlot, doctorName string) error {
	log.Print("Dear Mr(s) " + appointmnetSlot.PatientName + " Appointment successfully booked at " + appointmnetSlot.StartingTime.Format("2006-01-02 15:04:05") + " with doctor " + doctorName)
	return nil
}
