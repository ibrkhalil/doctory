package ports

import (
	"time"

	"github.com/ibrkhalil/doctory/internal/schema"
)

type DoctorAvailabilitySlotStorage interface {
	AddAvailabilitySlot(availability schema.DoctorAvailabilitySlot) error
	GetAvailabilityAtTime(date time.Time) (bool, error)
	ListAvailabilitySlots() ([]schema.DoctorAvailabilitySlot, error)
	ViewUpcomingAppointments() ([]schema.AppointmentSlot, error)
	CancelAppointmentById(ID string) bool
	ConfirmAppointmentById(ID string) bool
}
