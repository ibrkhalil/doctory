package appointmentBooking

import (
	"sync"
	"time"
)

type Appointment struct {
	ID        string
	PatientID string
	DoctorID  string
	StartTime time.Time
	EndTime   time.Time
}

type appointmentBookingDB struct {
	atomicMux    sync.Mutex
	appointments map[string]Appointment
}
