package confirmAppointment

import (
	"errors"
)

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		appointments: make(map[string]*Appointment),
	}
}

func (db *InMemoryDB) ConfirmAppointment(id string) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	appointment, exists := db.appointments[id]
	if !exists {
		return errors.New("appointment not found")
	}

	if appointment.Status == "confirmed" {
		return errors.New("appointment already confirmed")
	}

	appointment.Status = "confirmed"
	return nil
}

func (db *InMemoryDB) AddAppointment(appointment *Appointment) {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.appointments[appointment.ID] = appointment
}
