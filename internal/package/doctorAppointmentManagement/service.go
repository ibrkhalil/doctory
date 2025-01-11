package doctorAppointmentManagement

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ibrkhalil/doctory/internal/schema"
)

func Start(ginEngineInstance *gin.Engine) {
	InitModule(ginEngineInstance)
}

func validateAppointmentFromRequest(appointmentSlot schema.DoctorAvailabilitySlot) bool {
	if appointmentSlot.Cost == 0 || len(appointmentSlot.DoctorID) == 0 ||
		len(appointmentSlot.DoctorName) == 0 || len(appointmentSlot.ID) == 0 ||
		appointmentSlot.Time.IsZero() {
		return false
	}
	return true
}

func CreateAvailabilityFromRequest(ctx *gin.Context) (schema.DoctorAvailabilitySlot, error) {
	bodyAsByteArray, _ := ioutil.ReadAll(ctx.Request.Body)
	var appointmentSlot schema.DoctorAvailabilitySlot
	json.Unmarshal(bodyAsByteArray, &appointmentSlot)

	appointmentSlot.ID = uuid.NewString()
	// Assumes an appointment is one hour
	appointmentSlot.ToTime = appointmentSlot.Time.Add(time.Hour)

	isValidAppointment := validateAppointmentFromRequest(appointmentSlot)
	if isValidAppointment {
		return appointmentSlot, nil
	}
	return schema.DoctorAvailabilitySlot{}, errors.New("An error occured when parsing request")
}
