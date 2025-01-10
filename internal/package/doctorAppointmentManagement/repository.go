package doctorAppointmentManagement

import (
	"errors"
	"time"
)

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		availabilities: make(map[string]map[time.Time][]time.Time),
	}
}

func (db *InMemoryDB) AddAvailability(availability Availability) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if _, exists := db.availabilities[availability.DoctorID]; !exists {
		db.availabilities[availability.DoctorID] = make(map[time.Time][]time.Time)
	}

	db.availabilities[availability.DoctorID][availability.Date] = availability.Slots
	return nil
}

func (db *InMemoryDB) GetAvailability(doctorID string, date time.Time) ([]time.Time, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	if dates, exists := db.availabilities[doctorID]; exists {
		if slots, exists := dates[date]; exists {
			return slots, nil
		}
	}
	return nil, errors.New("availability not found")
}

func (db *InMemoryDB) RemoveAvailability(doctorID string, date time.Time) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if dates, exists := db.availabilities[doctorID]; exists {
		if _, exists := dates[date]; exists {
			delete(dates, date)
			return nil
		}
	}
	return errors.New("availability not found")
}
