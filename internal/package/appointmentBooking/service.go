package appointmentBooking

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/ibrkhalil/doctory/internal/schema"
)

func Start(ginEngineInstance *gin.Engine) {
	InitModule(ginEngineInstance)
}

func validateAppointmentFromRequest(appointmentSlot schema.AppointmentSlot) bool {
	if appointmentSlot.ReservedAt.IsZero() || len(appointmentSlot.PatientID) == 0 ||
		len(appointmentSlot.PatientName) == 0 || len(appointmentSlot.ID) == 0 ||
		len(appointmentSlot.SlotId) == 0 {
		return false
	}
	return true
}

func CreateAppointmentFromRequest(ctx *gin.Context) (schema.AppointmentSlot, error) {
	bodyAsByteArray, _ := ioutil.ReadAll(ctx.Request.Body)
	var appointmentSlot schema.AppointmentSlot
	json.Unmarshal(bodyAsByteArray, &appointmentSlot)
	isValidRequest := validateAppointmentFromRequest(appointmentSlot)
	if isValidRequest {
		return appointmentSlot, nil
	} else {
		return schema.AppointmentSlot{}, errors.New("Invalid Request")
	}
}
