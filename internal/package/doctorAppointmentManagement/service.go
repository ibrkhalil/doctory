package doctorAppointmentManagement

import "github.com/gin-gonic/gin"

func Start(ginEngineInstance *gin.Engine) {
	InitModule(ginEngineInstance)
}
