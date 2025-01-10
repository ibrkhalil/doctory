package confirmAppointment

import "sync"

type Appointment struct {
	ID     string
	Status string
}

type InMemoryDB struct {
	appointments map[string]*Appointment
	mu           sync.Mutex
}
