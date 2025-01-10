package app

import (
	"github.com/ibrkhalil/doctory/internal/package/appointmentBooking"
	"github.com/ibrkhalil/doctory/internal/package/confirmAppointment"
	"github.com/ibrkhalil/doctory/internal/package/doctorAppointmentManagement"
	"github.com/ibrkhalil/doctory/internal/package/doctorAvailability"
)

func Main() {
	appointmentBooking.Start()
	doctorAppointmentManagement.Start()
	confirmAppointment.Start()
	doctorAvailability.Start()
}
