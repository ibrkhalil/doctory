package repository

import (
	"errors"
	"time"

	"github.com/ibrkhalil/doctory/internal/db"
	"github.com/ibrkhalil/doctory/internal/package/doctorAppointmentManagement/core/ports"
	"github.com/ibrkhalil/doctory/internal/schema"
)

type DoctorAvailabilitySlotController struct {
	service *ports.DoctorAvailabilitySlotStorage
}

func NewDoctorAvailabilitySlotController() *DoctorAvailabilitySlotController {
	return &DoctorAvailabilitySlotController{}
}

func withinTimeSpan(start, end, check time.Time) bool {
	if start.Before(end) {
		return !check.Before(start) && !check.After(end)
	}
	if start.Equal(end) {
		return check.Equal(start)
	}
	return !start.After(check) || !end.Before(check)
}

func availabilityTimeConflict(availability schema.DoctorAvailabilitySlot) bool {
	availabilitySlots := db.GetInstance().GetAllDoctorAvailabilitySlots()
	if len(availabilitySlots) == 0 {
		return false
	}
	for _, availabilitySlot := range availabilitySlots {
		if withinTimeSpan(availabilitySlot.Time, availabilitySlot.ToTime, availability.Time) {
			return true
		}
	}
	return false
}

func (c *DoctorAvailabilitySlotController) AddAvailabilitySlot(availability schema.DoctorAvailabilitySlot) error {
	db := db.GetInstance()
	db.GetDoctorAvailabilitySlotByKey(availability.ID)
	if !availabilityTimeConflict(availability) {
		db.SetDoctorAvailabilitySlot(availability.ID, availability)
		return nil
	} else {
		return errors.New("availability slot time is not available")
	}
}

func (c *DoctorAvailabilitySlotController) GetAvailabilityAtTime(date time.Time) (bool, error) {
	db := db.GetInstance()
	availabilitySlots := db.GetAllDoctorAvailabilitySlots()
	for _, availabilitySlot := range availabilitySlots {
		if date.Unix() > availabilitySlot.Time.Unix() &&
			date.Unix() < availabilitySlot.ToTime.Unix() {
			// Within range
			return availabilitySlot.IsReserved, nil

		}
	}
	return false, errors.New("no available slots at the time")
}

func (c *DoctorAvailabilitySlotController) ListAvailabilitySlots() ([]schema.DoctorAvailabilitySlot, error) {
	db := db.GetInstance()
	availabilitySlots := db.GetAllDoctorAvailabilitySlots()
	if len(availabilitySlots) > 0 {
		return availabilitySlots, nil
	} else {
		return nil, errors.New("No availabilities exist for the doctor!")
	}

}

func (c *DoctorAvailabilitySlotController) ViewUpcomingAppointments() ([]schema.AppointmentSlot, error) {
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

func (c *DoctorAvailabilitySlotController) CancelAppointmentById(ID string) bool {
	db := db.GetInstance()
	return db.CancelAppointmentById(ID)
}

func (c *DoctorAvailabilitySlotController) ConfirmAppointmentById(ID string) bool {
	db := db.GetInstance()
	return db.ConfirmAppointmentById(ID)
}
