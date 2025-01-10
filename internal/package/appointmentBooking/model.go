package appointmentBooking

import (
	"sync"
	"time"
)

type Appointment struct {
	ID          string
	PatientID   string
	PatientName string
	ReservedAt  time.Time
}

type appointmentBookingDB struct {
	atomicMux    sync.Mutex
	appointments map[string]Appointment
}
