package schema

import (
	"time"
)

type AppointmentSlot struct {
	ID           string    `json:"id"`
	SlotId       string    `json:"slotId"`
	PatientID    string    `json:"patientId"`
	PatientName  string    `json:"patientName"`
	ReservedAt   time.Time `json:"reservedAt"`
	StartingTime time.Time `json:"startingTime"`
}

type DoctorAvailabilitySlot struct {
	ID         string    `json:"id"`
	Time       time.Time `json:"time"`
	DoctorID   string    `json:"doctorId"`
	DoctorName string    `json:"doctorName"`
	IsReserved bool      `json:"isReserved"`
	Cost       float32   `json:"cost"`
	ToTime     time.Time `json:"toTime"`
}
