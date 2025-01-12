package schema

import (
	"time"
)

type AppointmentState int

const (
	DEFAULT_APPOINTMENT_STATE AppointmentState = iota
	CONFIRMED_APPOINTMENT_STATE
	CANCELLED_APPOINTMENT_STATE
)

type AppointmentSlot struct {
	ID           string           `json:"id"`
	SlotId       int              `json:"slotId"`
	PatientID    string           `json:"patientId"`
	PatientName  string           `json:"patientName"`
	ReservedAt   time.Time        `json:"reservedAt"`
	StartingTime time.Time        `json:"startingTime"`
	State        AppointmentState `json:"state"`
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
