package doctorAppointmentManagement

import (
	"sync"
	"time"
)

type Availability struct {
	DoctorID string
	Date     time.Time
	Slots    []time.Time
}

type InMemoryDB struct {
	mu             sync.Mutex
	availabilities map[string]map[time.Time][]time.Time
}
