package appointmentBooking

import (
	"time"

	"github.com/google/uuid"
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
	appointment.ReservedAt = time.Now()
	db.SetAppointmentSlots(appointment.ID, appointment)
	db.SetDoctorAvailabilitySlot(availabilitySlot.ID, availabilitySlot)

}

func CreateAppointment(appointment schema.AppointmentSlot) error {
	db := db.GetInstance()
	appointment.ID = uuid.NewString()
	doctorAvailabilitySlots := db.GetAllDoctorAvailabilitySlots()
	for _, availabilitySlot := range doctorAvailabilitySlots {
		if !availabilitySlot.IsReserved {
			reserveAvailabilitySlotAndNotify(availabilitySlot, appointment)
		}
	}
	return nil
}

func ListAppointments() []schema.AppointmentSlot {
	appointments := db.GetInstance().GetAllAppointmentSlots()
	return appointments
}
