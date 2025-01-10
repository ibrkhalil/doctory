package schema

import (
	"sync"
	"time"
)

type AppointmentSlot struct {
	ID          string    `json:"id"`
	SlotId      string    `json:"slotId"`
	PatientID   string    `json:"patientId"`
	PatientName string    `json:"patientName"`
	ReservedAt  time.Time `json:"reservedAt"`
}

type AppointmentBookingDB struct {
	Mutex            sync.Mutex
	AppointmentSlots map[string]AppointmentSlot `json:"appointments"`
}

type DoctorAvailability struct {
	ID         string    `json:"id"`
	DoctorID   string    `json:"doctorId"`
	DoctorName string    `json:"doctorName"`
	IsReserved bool      `json:"isReserved"`
	Cost       float32   `json:"cost"`
	Time       time.Time `json:"time"`
	ToTime     time.Time `json:"toTime"`
}

type DoctorAvailabilityInMemoryDB struct {
	DoctorAvailibities []DoctorAvailability `json:"doctorAvailibities"`
	Mutex              sync.RWMutex
}
