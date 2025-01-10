package appointmentBooking

import (
	"errors"
)

func NewAppointmentBookingDB() *appointmentBookingDB {
	return &appointmentBookingDB{
		appointments: make(map[string]Appointment),
	}
}

func (db *appointmentBookingDB) CreateAppointment(appointment Appointment) error {
	db.atomicMux.Lock()
	defer db.atomicMux.Unlock()

	if _, exists := db.appointments[appointment.ID]; exists {
		return errors.New("appointment already exists")
	}

	db.appointments[appointment.ID] = appointment
	return nil
}

func (db *appointmentBookingDB) GetAppointment(id string) (Appointment, error) {
	db.atomicMux.Lock()
	defer db.atomicMux.Unlock()

	appointment, exists := db.appointments[id]
	if !exists {
		return Appointment{}, errors.New("appointment not found")
	}

	return appointment, nil
}

func (db *appointmentBookingDB) UpdateAppointment(id string, updated Appointment) error {
	db.atomicMux.Lock()
	defer db.atomicMux.Unlock()

	if _, exists := db.appointments[id]; !exists {
		return errors.New("appointment not found")
	}

	db.appointments[id] = updated
	return nil
}

func (db *appointmentBookingDB) DeleteAppointment(id string) error {
	db.atomicMux.Lock()
	defer db.atomicMux.Unlock()

	if _, exists := db.appointments[id]; !exists {
		return errors.New("appointment not found")
	}

	delete(db.appointments, id)
	return nil
}

func (db *appointmentBookingDB) ListAppointments() []Appointment {
	db.atomicMux.Lock()
	defer db.atomicMux.Unlock()

	appointments := make([]Appointment, 0, len(db.appointments))
	for _, appointment := range db.appointments {
		appointments = append(appointments, appointment)
	}

	return appointments
}
