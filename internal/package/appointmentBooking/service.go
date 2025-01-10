package appointmentBooking

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/ibrkhalil/doctory/internal/schema"
)

func Start(ginEngineInstance *gin.Engine) {
	InitModule(ginEngineInstance)
}

func CreateAppointmentFromRequest(ctx *gin.Context) schema.AppointmentSlot {
	bodyAsByteArray, _ := ioutil.ReadAll(ctx.Request.Body)
	var appointmentSlot schema.AppointmentSlot
	json.Unmarshal(bodyAsByteArray, &appointmentSlot)
	return appointmentSlot
}
