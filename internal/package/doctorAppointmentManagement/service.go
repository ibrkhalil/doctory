package doctorAppointmentManagement

import "github.com/gin-gonic/gin"

func Start(ginInstanceEngine *gin.Engine) {
	InitModule(ginInstanceEngine)
}
