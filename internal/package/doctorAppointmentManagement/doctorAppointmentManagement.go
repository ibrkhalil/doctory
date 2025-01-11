package doctorAppointmentManagement

import "github.com/gin-gonic/gin"

func InitModule(ginEngineInstance *gin.Engine) {
	RegisterRoutes(ginEngineInstance)
}
