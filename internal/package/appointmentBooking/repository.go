package appointmentBooking

import (
	"errors"

	"github.com/ibrkhalil/doctory/internal/schema"
)

type appointmentSlotsDB struct {
	*schema.AppointmentBookingDB
}

func NewInMemoryAppointmentSlotsDB() *appointmentSlotsDB {
	return &appointmentSlotsDB{}
}

func (db *appointmentSlotsDB) CreateAppointment(appointment schema.AppointmentSlot) error {
	db.Mutex.Lock()
	defer db.Mutex.Unlock()

	if _, exists := db.AppointmentSlots[appointment.ID]; exists {
		return errors.New("appointment already exists")
	}

	db.AppointmentSlots[appointment.ID] = appointment
	return nil
}

func (db *appointmentSlotsDB) GetAppointment(id string) (schema.AppointmentSlot, error) {
	db.Mutex.Lock()
	defer db.Mutex.Unlock()

	appointment, exists := db.AppointmentSlots[id]
	if !exists {
		return schema.AppointmentSlot{}, errors.New("appointment not found")
	}

	return appointment, nil
}

func (db *appointmentSlotsDB) UpdateAppointment(id string, updated schema.AppointmentSlot) error {
	db.Mutex.Lock()
	defer db.Mutex.Unlock()

	if _, exists := db.AppointmentSlots[id]; !exists {
		return errors.New("appointment not found")
	}

	db.AppointmentSlots[id] = updated
	return nil
}

func (db *appointmentSlotsDB) DeleteAppointment(id string) error {
	db.Mutex.Lock()
	defer db.Mutex.Unlock()

	if _, exists := db.AppointmentSlots[id]; !exists {
		return errors.New("appointment not found")
	}

	delete(db.AppointmentSlots, id)
	return nil
}

func (db *appointmentSlotsDB) ListAppointments() []schema.AppointmentSlot {
	db.Mutex.Lock()
	defer db.Mutex.Unlock()

	appointments := make([]schema.AppointmentSlot, 0, len(db.AppointmentSlots))
	for _, appointment := range db.AppointmentSlots {
		appointments = append(appointments, appointment)
	}

	return appointments
}
