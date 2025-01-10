package doctorAppointmentManagement

import (
	"errors"
	"time"

	"github.com/ibrkhalil/doctory/internal/package/confirmAppointment"
	"github.com/ibrkhalil/doctory/internal/schema"
)

type availabilityDB struct {
	*schema.DoctorAvailabilityInMemoryDB
}

func NewInMemoryAvailabilityDB() *schema.DoctorAvailabilityInMemoryDB {
	return &schema.DoctorAvailabilityInMemoryDB{
		DoctorAvailibities: []schema.DoctorAvailability{},
	}
}

func (db availabilityDB) AddAvailability(availability schema.DoctorAvailability) error {
	db.Mutex.Lock()
	defer db.Mutex.Unlock()
	availability.ToTime = availability.Time.Add(time.Hour)
	db.DoctorAvailibities = append(db.DoctorAvailibities, availability)
	return nil
}

func (db availabilityDB) GetAvailabilityAtTime(date time.Time) (bool, error) {
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

func (db *availabilityDB) RemoveAvailability(ID string) {
	db.Mutex.Lock()
	defer db.Mutex.Unlock()

	var filteredAvailabilities []schema.DoctorAvailability
	for _, v := range db.DoctorAvailibities {
		if v.ID != ID {
			filteredAvailabilities = append(filteredAvailabilities, v)
		}
	}
	db.DoctorAvailibities = filteredAvailabilities
}

func (db availabilityDB) ListAvailability() ([]schema.DoctorAvailability, error) {
	db.Mutex.Lock()
	defer db.Mutex.Unlock()

	if len(db.DoctorAvailibities) > 0 {
		return db.DoctorAvailibities, nil
	} else {
		return nil, errors.New("No availabilities exist for the doctor!")
	}
}

func (db availabilityDB) ViewUpcomingAppointments() ([]schema.DoctorAvailability, error) {
	db.Mutex.Lock()
	defer db.Mutex.Unlock()
	now := time.Now()
	var futureAvailabilities []schema.DoctorAvailability
	for _, v := range db.DoctorAvailibities {
		if v.Time.After(now) {
			futureAvailabilities = append(futureAvailabilities, v)
		}
	}
	return futureAvailabilities, nil
}

func (db availabilityDB) CancelAppointmentAtTime(availabilityTime time.Time) ([]schema.DoctorAvailability, error) {
	db.Mutex.Lock()
	defer db.Mutex.Unlock()
	var filteredAvailabilities []schema.DoctorAvailability
	for _, v := range db.DoctorAvailibities {
		if v.Time != availabilityTime {
			filteredAvailabilities = append(filteredAvailabilities, v)
		}
	}
	return filteredAvailabilities, nil
}

func (db availabilityDB) CancelAppointmentById(ID string) ([]schema.DoctorAvailability, error) {
	db.Mutex.Lock()
	defer db.Mutex.Unlock()
	var filteredAvailabilities []schema.DoctorAvailability
	for _, v := range db.DoctorAvailibities {
		if v.ID != ID {
			filteredAvailabilities = append(filteredAvailabilities, v)
		}
	}
	return filteredAvailabilities, nil
}

func (db availabilityDB) ConfirmAppointmentById(ID string) ([]schema.DoctorAvailability, error) {
	db.Mutex.Lock()
	defer db.Mutex.Unlock()
	var targetedAvailabilityToConfirm schema.DoctorAvailability
	var filteredAvailabilities []schema.DoctorAvailability
	for _, v := range db.DoctorAvailibities {
		if v.ID != ID {
			filteredAvailabilities = append(filteredAvailabilities, v)
		} else {
			targetedAvailabilityToConfirm = v
		}
	}
	confirmAppointment.NotifyDoctorOfConfirmation(targetedAvailabilityToConfirm)
	confirmAppointment.NotifyPatientOfConfirmation(targetedAvailabilityToConfirm)

	db.DoctorAvailibities = filteredAvailabilities
	return filteredAvailabilities, nil
}
