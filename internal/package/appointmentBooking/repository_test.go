package appointmentBooking

import (
	"log"
	"testing"
	"time"

	"github.com/ibrkhalil/doctory/internal/package/doctorAppointmentManagement"
	"github.com/ibrkhalil/doctory/internal/schema"
	"github.com/stretchr/testify/assert"
)

func TestCreateAppointment(t *testing.T) {
	doctorName := "MAHOOOOODZ"
	patientName := "Mohamed Henedy"
	appointment := schema.AppointmentSlot{ID: "1", PatientName: patientName}
	result, err := CreateAppointment(&appointment)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// There should be an error as there's no availability times available.
	assert.NotEqual(t, result, nil)

	availabilitySlot1 := schema.DoctorAvailabilitySlot{
		ID:         "1",
		Time:       time.Now().Add(time.Hour),
		DoctorID:   "1",
		DoctorName: doctorName,
		IsReserved: false,
		Cost:       5,
	}

	// Create availability slot
	doctorAppointmentManagement.AddAvailabilitySlot(availabilitySlot1)

	appointment2 := schema.AppointmentSlot{ID: "1", PatientName: patientName}
	result2, err := CreateAppointment(&appointment2)

	// Succeeded
	assert.Equal(t, result2, true)

}

func TestListAppointments(t *testing.T) {
	patient1Name := "Mohamed Henedy"
	patient2Name := "Michael Scofield"
	doctorName := "MAHOOOOODZ"

	availabilitySlot1 := schema.DoctorAvailabilitySlot{
		ID:         "1",
		DoctorID:   "1",
		Time:       time.Now().Add(time.Hour),
		DoctorName: doctorName,
		IsReserved: false,
		Cost:       5,
	}
	availabilitySlot2 := schema.DoctorAvailabilitySlot{
		ID:         "2",
		DoctorID:   "1",
		Time:       time.Now().Add(2 * time.Hour),
		DoctorName: doctorName,
		IsReserved: false,
		Cost:       5,
	}
	errorAddingAvailabilitySlot1 := doctorAppointmentManagement.AddAvailabilitySlot(availabilitySlot1)
	errorAddingAvailabilitySlot2 := doctorAppointmentManagement.AddAvailabilitySlot(availabilitySlot2)

	if errorAddingAvailabilitySlot1 != nil {
		log.Print("Eror adding availability slot 1")
	}

	if errorAddingAvailabilitySlot2 != nil {
		log.Print("Eror adding availability slot 2")
	}

	appointment := schema.AppointmentSlot{ID: "1", PatientName: patient1Name}
	appointment2 := schema.AppointmentSlot{ID: "2", PatientName: patient2Name}
	_, errorCreatingAppointment1 := CreateAppointment(&appointment)
	_, errorCreatingAppointment2 := CreateAppointment(&appointment2)

	if errorCreatingAppointment1 != nil {
		log.Print("Error creating appointment1!")
	}

	if errorCreatingAppointment2 != nil {
		log.Print("Error creating appointment2!")
	}

	result := ListAppointments()
	time.Sleep(300)
	assert.Equal(t, result[0].PatientName, patient1Name)
	assert.Equal(t, result[1].PatientName, patient2Name)
	assert.Len(t, result, 2)
}
