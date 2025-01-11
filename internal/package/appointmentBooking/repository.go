package appointmentBooking

import (
	"errors"

	"github.com/ibrkhalil/doctory/internal/db"
	"github.com/ibrkhalil/doctory/internal/package/confirmAppointment"
	"github.com/ibrkhalil/doctory/internal/schema"
)

func reserveAvailabilitySlotAndNotify(availabilitySlot schema.DoctorAvailabilitySlot, appointment schema.AppointmentSlot) {
	db := db.GetInstance()
	availabilitySlot.IsReserved = true
	confirmAppointment.NotifyDoctorOfAppointmentBooking(appointment)
	confirmAppointment.NotifyPatientOfAppointmentBooking(appointment)
	appointment.StartingTime = availabilitySlot.Time
	db.SetAppointmentSlots(appointment.ID, appointment)
	db.SetDoctorAvailabilitySlot(availabilitySlot.ID, availabilitySlot)

}

func CreateAppointment(appointment schema.AppointmentSlot) error {
	db := db.GetInstance()
	_, alreadyExists := db.GetAppointmentSlotByKey(appointment.ID)
	if alreadyExists {
		return errors.New("Appointment already alreadyExists")
	} else {
		doctorAvailabilitySlots := db.GetAllDoctorAvailabilitySlots()
		for _, availabilitySlot := range doctorAvailabilitySlots {
			if !availabilitySlot.IsReserved {
				reserveAvailabilitySlotAndNotify(availabilitySlot, appointment)
			}
		}
		return nil
	}
}

func ListAppointments() []schema.AppointmentSlot {
	appointments := db.GetInstance().GetAllAppointmentSlots()
	return appointments
}
