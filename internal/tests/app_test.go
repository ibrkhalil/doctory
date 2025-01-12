package apptests

import (
	"log"
	"testing"
	"time"

	"github.com/ibrkhalil/doctory/internal/package/appointmentBooking"
	"github.com/ibrkhalil/doctory/internal/package/doctorAppointmentManagement"
	"github.com/ibrkhalil/doctory/internal/schema"
)

func TestCreateAppointment(t *testing.T) {
	appointment := schema.AppointmentSlot{ID: "1", PatientName: "Mohamed Henedy"}
	result, err := appointmentBooking.CreateAppointment(&appointment)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result {
		log.Print("Successful booking")
	} else {
		log.Print("Failure when booking appointment")
	}
}

func TestListAppointments(t *testing.T) {
	availabilitySlot1 := schema.DoctorAvailabilitySlot{
		ID:         "1",
		Time:       time.Now().Add(time.Hour),
		DoctorID:   "1",
		DoctorName: "MAHOOOOODZ",
		IsReserved: false,
		Cost:       5,
	}
	availabilitySlot2 := schema.DoctorAvailabilitySlot{
		ID:         "1",
		Time:       time.Now().Add(2 * time.Hour),
		DoctorID:   "1",
		DoctorName: "MAHOOOOODZ",
		IsReserved: false,
		Cost:       5,
	}
	errorAddingAvailabilitySlot1 := doctorAppointmentManagement.AddAvailabilitySlot(availabilitySlot1)
	errorAddingAvailabilitySlot2 := doctorAppointmentManagement.AddAvailabilitySlot(availabilitySlot2)

	if errorAddingAvailabilitySlot1 != nil {
		log.Print("Eror adding availability slot 1")
	}

	if errorAddingAvailabilitySlot2 != nil {
		log.Print("Eror adding availability slot 1")
	}

	appointment := schema.AppointmentSlot{ID: "1", PatientName: "Mohamed Henedy"}
	appointment2 := schema.AppointmentSlot{ID: "1", PatientName: "Michael Scofield"}
	_, errorCreatingAppointment1 := appointmentBooking.CreateAppointment(&appointment)
	_, errorCreatingAppointment2 := appointmentBooking.CreateAppointment(&appointment2)

	if errorCreatingAppointment1 != nil {
		log.Print("Error creating appointment1!")
	}

	if errorCreatingAppointment2 != nil {
		log.Print("Error creating appointment2!")
	}

	result := appointmentBooking.ListAppointments()
	if len(result) == 2 {
		log.Print("Successfully booked and listed 2 appointments!")
	} else {
		log.Print("Failure when booking appointment")
	}
}
