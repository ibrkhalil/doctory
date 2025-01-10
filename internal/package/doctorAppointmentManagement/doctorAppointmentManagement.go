package doctorAppointmentManagement

import "github.com/gin-gonic/gin"

func InitModule(ginInstanceEngine *gin.Engine) {
	RegisterRoutes(ginInstanceEngine)
}

var DoctorAppointmentManagementDB = NewInMemoryDB()
