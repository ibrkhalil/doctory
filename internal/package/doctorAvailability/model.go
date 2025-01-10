package doctorAvailability

import "sync"

type DoctorAvailability struct {
	ID        string
	DoctorID  string
	StartTime string
	EndTime   string
}

type InMemoryDB struct {
	data map[string]DoctorAvailability
	mu   sync.RWMutex
}
