package doctorAppointmentManagement

import (
	"errors"
	"time"

	"github.com/ibrkhalil/doctory/internal/db"
	"github.com/ibrkhalil/doctory/internal/schema"
)

func AddAvailabilitySlot(availability schema.DoctorAvailabilitySlot) error {
	db := db.GetInstance()
	_, alreadyExists := db.GetDoctorAvailabilitySlotByKey(availability.ID)
	if !alreadyExists {
		db.SetDoctorAvailabilitySlot(availability.ID, availability)
	} else {
		errors.New("Availability slot already taken!")
	}
	return nil
}

func GetAvailabilityAtTime(date time.Time) (bool, error) {
	db := db.GetInstance()
	availabilitySlots := db.GetAllDoctorAvailabilitySlots()
	for _, availabilitySlot := range availabilitySlots {
		if date.Unix() > availabilitySlot.Time.Unix() &&
			date.Unix() < availabilitySlot.ToTime.Unix() {
			// Within range
			return availabilitySlot.IsReserved, nil

		}
	}
	return false, errors.New("No available slots at the time")
}

func ListAppointmentSlots() ([]schema.DoctorAvailabilitySlot, error) {
	db := db.GetInstance()
	availabilitySlots := db.GetAllDoctorAvailabilitySlots()
	if len(availabilitySlots) > 0 {
		return availabilitySlots, nil
	} else {
		return nil, errors.New("No availabilities exist for the doctor!")
	}

}

func ViewUpcomingAppointments() ([]schema.AppointmentSlot, error) {
	db := db.GetInstance()
	appointments := db.GetAllAppointmentSlots()
	now := time.Now()
	var futureAvailabilities []schema.AppointmentSlot
	for _, appointment := range appointments {
		if appointment.StartingTime.After(now) && !appointment.ReservedAt.IsZero() {
			futureAvailabilities = append(futureAvailabilities, appointment)
		}
	}
	return futureAvailabilities, nil
}

func CancelAppointmentById(ID string) bool {
	db := db.GetInstance()
	return db.CancelAppointmentById(ID)
}

func ConfirmAppointmentById(ID string) bool {
	db := db.GetInstance()
	return db.ConfirmAppointmentById(ID)
}
