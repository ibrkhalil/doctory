package repository

import (
	"log"
	"testing"
	"time"

	"github.com/ibrkhalil/doctory/internal/db"
	"github.com/ibrkhalil/doctory/internal/package/appointmentBooking/infrastructure"
	"github.com/ibrkhalil/doctory/internal/schema"
	"github.com/stretchr/testify/assert"
)

func TestViewUpcomingAppointments(t *testing.T) {
	db.GetInstance()
	defer db.Clear()

	patient1Name := "Mohamed Henedy"
	patient2Name := "Michael Scofield"
	doctorName := "Dr. Gregory House"

	availabilitySlot1 := schema.DoctorAvailabilitySlot{
		ID:         "1",
		DoctorID:   "1",
		Time:       time.Now().Add(time.Hour),
		DoctorName: doctorName,
		IsReserved: false,
		Cost:       5,
		ToTime:     time.Now().Add(2 * time.Hour),
	}
	availabilitySlot2 := schema.DoctorAvailabilitySlot{
		ID:         "2",
		DoctorID:   "1",
		Time:       time.Now().Add(2 * time.Hour),
		DoctorName: doctorName,
		IsReserved: false,
		Cost:       5,
		ToTime:     time.Now().Add(3 * time.Hour),
	}
	service := NewDoctorAvailabilitySlotController()
	errorAddingAvailabilitySlot1 := service.AddAvailabilitySlot(availabilitySlot1)
	errorAddingAvailabilitySlot2 := service.AddAvailabilitySlot(availabilitySlot2)

	if errorAddingAvailabilitySlot1 != nil {
		log.Print("Error adding availability slot 1: ", errorAddingAvailabilitySlot1)
	}

	if errorAddingAvailabilitySlot2 != nil {
		log.Print("Error adding availability slot 2: ", errorAddingAvailabilitySlot2)
	}

	appointment1 := schema.AppointmentSlot{ID: "1", PatientName: patient1Name}
	appointment2 := schema.AppointmentSlot{ID: "2", PatientName: patient2Name}
	_, errorCreatingAppointment1 := infrastructure.CreateAppointment(&appointment1)
	_, errorCreatingAppointment2 := infrastructure.CreateAppointment(&appointment2)

	if errorCreatingAppointment1 != nil {
		log.Print("Error creating appointment1!")
	}

	if errorCreatingAppointment2 != nil {
		log.Print("Error creating appointment2!")
	}

	result, err := service.ViewUpcomingAppointments()
	assert.NoError(t, err)
	assert.Equal(t, result[0].PatientName, patient1Name)
	assert.Equal(t, result[1].PatientName, patient2Name)
	assert.Len(t, result, 2)
}
