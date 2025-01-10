package confirmAppointment

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func Start(ginInstanceEngine *gin.Engine) {
	InitModule(ginInstanceEngine)
}

func CreateAppointmentFromRequest(c *gin.Context) Appointment {
	bodyAsByteArray, _ := ioutil.ReadAll(c.Request.Body)
	var jsonMap Appointment
	json.Unmarshal(bodyAsByteArray, &jsonMap)
	return jsonMap
}
