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
	for _, v := range availabilitySlots {
		if date.Unix() > v.Time.Unix() &&
			date.Unix() < v.ToTime.Unix() {
			// Within range
			return v.IsReserved, nil

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

func ViewUpcomingAppointments() ([]schema.DoctorAvailabilitySlot, error) {
	db := db.GetInstance()
	availabilitySlots := db.GetAllDoctorAvailabilitySlots()
	now := time.Now()
	var futureAvailabilities []schema.DoctorAvailabilitySlot
	for _, availabilitySlot := range availabilitySlots {
		if availabilitySlot.Time.After(now) && availabilitySlot.IsReserved {
			futureAvailabilities = append(futureAvailabilities, availabilitySlot)
		}
	}
	return futureAvailabilities, nil
}

func CancelAppointmentAtTime(availabilityTime time.Time) ([]schema.DoctorAvailabilitySlot, error) {
	db := db.GetInstance()
	availabilitySlots := db.GetAllDoctorAvailabilitySlots()
	var filteredAvailabilities []schema.DoctorAvailabilitySlot
	for _, v := range availabilitySlots {
		if v.Time != availabilityTime {
			filteredAvailabilities = append(filteredAvailabilities, v)
		}
	}
	return filteredAvailabilities, nil
}

func CancelAppointmentById(ID string) ([]schema.DoctorAvailabilitySlot, error) {
	db := db.GetInstance()
	availabilitySlots := db.GetAllDoctorAvailabilitySlots()
	var filteredAvailabilities []schema.DoctorAvailabilitySlot
	for _, v := range availabilitySlots {
		if v.ID != ID {
			filteredAvailabilities = append(filteredAvailabilities, v)
		}
	}
	return filteredAvailabilities, nil
}
