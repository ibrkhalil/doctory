package doctorAppointmentManagement

import (
	"errors"
	"time"

	"github.com/ibrkhalil/doctory/internal/db"
	"github.com/ibrkhalil/doctory/internal/schema"
)

func inTimeSpan(start, end, check time.Time) bool {
	if start.Before(end) {
		return !check.Before(start) && !check.After(end)
	}
	if start.Equal(end) {
		return check.Equal(start)
	}
	return !start.After(check) || !end.Before(check)
}

func availabilityTimeExists(availability schema.DoctorAvailabilitySlot) bool {
	availabilitySlots := db.GetInstance().GetAllDoctorAvailabilitySlots()
	for _, availabilitySlot := range availabilitySlots {
		if inTimeSpan(availabilitySlot.Time, availabilitySlot.ToTime, availability.Time) {
			return true
		}
	}
	return false
}

func AddAvailabilitySlot(availability schema.DoctorAvailabilitySlot) error {
	db := db.GetInstance()
	_, idAalreadyExists := db.GetDoctorAvailabilitySlotByKey(availability.ID)
	if !idAalreadyExists {
		if !availabilityTimeExists(availability) {
			db.SetDoctorAvailabilitySlot(availability.ID, availability)
		} else {
			return errors.New("Availability slot already taken!")
		}
	} else {
		return errors.New("Availability slot already taken!")
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

func ListAvailabilitySlots() ([]schema.DoctorAvailabilitySlot, error) {
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
