package infrastructure

import (
	"log"
	"testing"
	"time"

	"github.com/ibrkhalil/doctory/internal/db"
	"github.com/ibrkhalil/doctory/internal/package/doctorAppointmentManagement/adapter/repository"
	"github.com/ibrkhalil/doctory/internal/schema"
	"github.com/stretchr/testify/assert"
)

func TestCreateAppointment(t *testing.T) {
	db.GetInstance()
	defer db.Clear()

	doctorName := "Dr. Gregory House"
	patientName := "Jason Statham"
	appointment := schema.AppointmentSlot{ID: "1", PatientName: patientName}
	result, err := CreateAppointment(&appointment)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	assert.NotEqual(t, result, nil)

	availabilitySlot1 := schema.DoctorAvailabilitySlot{
		ID:         "1",
		Time:       time.Now().Add(time.Hour),
		DoctorID:   "1",
		DoctorName: doctorName,
		IsReserved: false,
		Cost:       5,
	}

	service := repository.NewDoctorAvailabilitySlotController()
	service.AddAvailabilitySlot(availabilitySlot1)

	appointment2 := schema.AppointmentSlot{ID: "1", PatientName: patientName}
	result2, err := CreateAppointment(&appointment2)

	assert.Equal(t, result2, true)

}

func TestListAppointments(t *testing.T) {
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
	service := repository.NewDoctorAvailabilitySlotController()
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
	_, errorCreatingAppointment1 := CreateAppointment(&appointment1)
	_, errorCreatingAppointment2 := CreateAppointment(&appointment2)

	if errorCreatingAppointment1 != nil {
		log.Print("Error creating appointment1!")
	}

	if errorCreatingAppointment2 != nil {
		log.Print("Error creating appointment2!")
	}

	result := ListAppointments()
	assert.Equal(t, result[0].PatientName, patient1Name)
	assert.Equal(t, result[1].PatientName, patient2Name)
	assert.Len(t, result, 2)
}

func TestConfirmAppoinment(t *testing.T) {
	db.GetInstance()
	defer db.Clear()

	doctorName := "Dr. Gregory House"
	patientName := "Jason Statham"
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
	service := repository.NewDoctorAvailabilitySlotController()
	service.AddAvailabilitySlot(availabilitySlot1)

	appointment2 := schema.AppointmentSlot{ID: "1", PatientName: patientName}
	_, errCreatingAppointment := CreateAppointment(&appointment2)

	assert.NoError(t, errCreatingAppointment)

	status := service.ConfirmAppointmentById(appointment2.ID)

	appointments := db.GetInstance().GetAllAppointmentSlots()
	assert.Equal(t, appointments[0].State, schema.CONFIRMED_APPOINTMENT_STATE)

	assert.True(t, status)
}

func TestCancelAppoinment(t *testing.T) {
	db.GetInstance()
	defer db.Clear()

	doctorName := "Dr. Gregory House"
	patientName := "Jason Statham"
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
	service := repository.NewDoctorAvailabilitySlotController()
	service.AddAvailabilitySlot(availabilitySlot1)

	appointment2 := schema.AppointmentSlot{ID: "1", PatientName: patientName}
	_, errCreatingAppointment := CreateAppointment(&appointment2)

	assert.NoError(t, errCreatingAppointment)

	status := service.CancelAppointmentById(appointment2.ID)
	appointments := db.GetInstance().GetAllAppointmentSlots()
	assert.Equal(t, appointments[0].State, schema.CANCELLED_APPOINTMENT_STATE)
	assert.True(t, status)
}
