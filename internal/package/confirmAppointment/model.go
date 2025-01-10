package confirmAppointment

import "sync"

type AppointmentConfirmation struct {
	ID     string
	Status string
}

type InMemoryDB struct {
	appointments map[string]AppointmentConfirmation
	mu           sync.Mutex
}
