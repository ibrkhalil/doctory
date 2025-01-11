package appointmentBooking

import (
	"time"

	"github.com/google/uuid"
	"github.com/ibrkhalil/doctory/internal/db"
	"github.com/ibrkhalil/doctory/internal/package/confirmAppointment"
	"github.com/ibrkhalil/doctory/internal/schema"
)

func reserveAvailabilitySlotAndNotify(db *db.SingletonDB, availabilitySlot schema.DoctorAvailabilitySlot, appointment *schema.AppointmentSlot) {
	availabilitySlot.IsReserved = true
	appointment.StartingTime = availabilitySlot.Time.UTC()
	appointment.ReservedAt = time.Now().UTC()
	confirmAppointment.NotifyDoctorOfAppointmentBooking(*appointment)
	confirmAppointment.NotifyPatientOfAppointmentBooking(*appointment)
	db.SetAppointmentSlots(appointment.ID, *appointment)
	db.SetDoctorAvailabilitySlot(availabilitySlot.ID, availabilitySlot)

}

func CreateAppointment(appointment schema.AppointmentSlot) bool {
	db := db.GetInstance()
	appointment.ID = uuid.NewString()
	doctorAvailabilitySlots := db.GetAllDoctorAvailabilitySlots()
	var reservedFlag bool
	for _, availabilitySlot := range doctorAvailabilitySlots {
		if !availabilitySlot.IsReserved && !reservedFlag {
			reserveAvailabilitySlotAndNotify(db, availabilitySlot, &appointment)
			reservedFlag = true
		}

	}
	return reservedFlag
}

func ListAppointments() []schema.AppointmentSlot {
	appointments := db.GetInstance().GetAllAppointmentSlots()
	return appointments
}
