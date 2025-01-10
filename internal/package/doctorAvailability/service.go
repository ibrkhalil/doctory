package doctorAvailability

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/ibrkhalil/doctory/internal/schema"
)

func Start(ginEngineInstance *gin.Engine) {
	InitModule(ginEngineInstance)
}

func CreateAvailabilityFromRequest(ctx *gin.Context) schema.DoctorAvailability {
	bodyAsByteArray, _ := ioutil.ReadAll(ctx.Request.Body)
	var appointmentSlot schema.DoctorAvailability
	json.Unmarshal(bodyAsByteArray, &appointmentSlot)
	return appointmentSlot
}
