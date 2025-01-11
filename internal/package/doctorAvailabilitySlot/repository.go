package doctorAvailabilitySlot

import (
	"errors"
	"time"

	"github.com/ibrkhalil/doctory/internal/schema"
)

type InMemoryAvailabilityDB struct {
	*schema.DoctorAvailabilitySlotInMemoryDB
}

func NewInMemoryAvailabilityDB() *InMemoryAvailabilityDB {
	return &InMemoryAvailabilityDB{
		DoctorAvailabilitySlotInMemoryDB: &schema.DoctorAvailabilitySlotInMemoryDB{
			DoctorAvailibities: []schema.DoctorAvailabilitySlot{},
		},
	}
}

func (db InMemoryAvailabilityDB) AddAvailabilitySlot(availability schema.DoctorAvailabilitySlot) error {
	db.Mutex.Lock()
	defer db.Mutex.Unlock()
	db.DoctorAvailibities = append(db.DoctorAvailibities, availability)
	return nil
}

func (db InMemoryAvailabilityDB) UpdateAvailability(availability schema.DoctorAvailabilitySlot) error {
	db.Mutex.Lock()
	defer db.Mutex.Unlock()
	var res []schema.DoctorAvailabilitySlot
	for _, v := range db.DoctorAvailibities {
		if v.ID == availability.ID {
			res = append(res, availability)
		} else {
			res = append(res, v)
		}
	}
	return nil
}

func (db InMemoryAvailabilityDB) GetAvailabilityAtTime(date time.Time) (bool, error) {
	db.Mutex.Lock()
	defer db.Mutex.Unlock()
	for _, v := range db.DoctorAvailibities {
		if date.Unix() > v.Time.Unix() &&
			date.Unix() < v.ToTime.Unix() {
			// Within range
			return v.IsReserved, nil

		}
	}
	return false, errors.New("No available slots at the time")
}

func (db *InMemoryAvailabilityDB) DeleteAvailability(ID string) {
	db.Mutex.Lock()
	defer db.Mutex.Unlock()

	var filteredAvailabilities []schema.DoctorAvailabilitySlot
	for _, v := range db.DoctorAvailibities {
		if v.ID != ID {
			filteredAvailabilities = append(filteredAvailabilities, v)
		}
	}
	db.DoctorAvailibities = filteredAvailabilities
}

func (db InMemoryAvailabilityDB) ListAppointmentSlots() ([]schema.DoctorAvailabilitySlot, error) {
	db.Mutex.Lock()
	defer db.Mutex.Unlock()

	if len(db.DoctorAvailibities) > 0 {
		return db.DoctorAvailibities, nil
	} else {
		return nil, errors.New("No availabilities exist for the doctor!")
	}

}

func (db InMemoryAvailabilityDB) ViewUpcomingAppointments() ([]schema.DoctorAvailabilitySlot, error) {
	db.Mutex.Lock()
	defer db.Mutex.Unlock()
	now := time.Now()
	var futureAvailabilities []schema.DoctorAvailabilitySlot
	for _, v := range db.DoctorAvailibities {
		if v.Time.After(now) {
			futureAvailabilities = append(futureAvailabilities, v)
		}
	}
	return futureAvailabilities, nil
}

func (db InMemoryAvailabilityDB) CancelAppointmentAtTime(availabilityTime time.Time) ([]schema.DoctorAvailabilitySlot, error) {
	db.Mutex.Lock()
	defer db.Mutex.Unlock()
	var filteredAvailabilities []schema.DoctorAvailabilitySlot
	for _, v := range db.DoctorAvailibities {
		if v.Time != availabilityTime {
			filteredAvailabilities = append(filteredAvailabilities, v)
		}
	}
	return filteredAvailabilities, nil
}

func (db InMemoryAvailabilityDB) CancelAppointmentById(ID string) ([]schema.DoctorAvailabilitySlot, error) {
	db.Mutex.Lock()
	defer db.Mutex.Unlock()
	var filteredAvailabilities []schema.DoctorAvailabilitySlot
	for _, v := range db.DoctorAvailibities {
		if v.ID != ID {
			filteredAvailabilities = append(filteredAvailabilities, v)
		}
	}
	return filteredAvailabilities, nil
}
