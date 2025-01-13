package infrastructure

import (
	"time"

	"github.com/google/uuid"
	"github.com/ibrkhalil/doctory/internal/db"
	"github.com/ibrkhalil/doctory/internal/package/confirmAppointment"
	"github.com/ibrkhalil/doctory/internal/schema"
)

type AppointmeentStorage interface {
	CreateAppointment(appointment schema.AppointmentSlot) (bool, error)
	ListAppointments() []schema.AppointmentSlot
}

func reserveAvailabilitySlotAndNotify(db *db.SingletonDB, availabilitySlot schema.DoctorAvailabilitySlot, appointment *schema.AppointmentSlot) error {
	availabilitySlot.IsReserved = true
	appointment.StartingTime = availabilitySlot.Time
	availabilitySlot.ToTime = appointment.StartingTime.Add(time.Hour)
	appointment.ReservedAt = time.Now()

	errorNotifiyingDoctor := confirmAppointment.NotifyDoctorOfAppointmentBooking(*appointment)
	errorNotifiyingPatient := confirmAppointment.NotifyPatientOfAppointmentBooking(*appointment, availabilitySlot.DoctorName)

	if errorNotifiyingDoctor != nil {
		return errorNotifiyingDoctor
	}

	if errorNotifiyingPatient != nil {
		return errorNotifiyingPatient
	}

	db.SetAppointmentSlots(appointment.ID, *appointment)
	db.SetDoctorAvailabilitySlot(availabilitySlot.ID, availabilitySlot)

	return nil
}

func CreateAppointment(appointment *schema.AppointmentSlot) (bool, error) {
	db := db.GetInstance()
	appointment.ID = uuid.NewString()
	doctorAvailabilitySlots := db.GetAllDoctorAvailabilitySlots()
	reservedFlag := false
	for _, availabilitySlot := range doctorAvailabilitySlots {
		if !availabilitySlot.IsReserved && !reservedFlag {
			reservationErrorStatus := reserveAvailabilitySlotAndNotify(db, availabilitySlot, appointment)
			if reservationErrorStatus != nil {
				return false, reservationErrorStatus
			}
			reservedFlag = true
		}

	}
	return reservedFlag, nil
}

func ListAppointments() []schema.AppointmentSlot {
	appointments := db.GetInstance().GetAllAppointmentSlots()
	return appointments
}
