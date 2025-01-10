package appointmentBooking

import "github.com/gin-gonic/gin"

func InitModule(ginEngineInstance *gin.Engine) {
	RegisterRoutes(ginEngineInstance)
}

var AppointmentSlotsDB = NewInMemoryAppointmentSlotsDB()
