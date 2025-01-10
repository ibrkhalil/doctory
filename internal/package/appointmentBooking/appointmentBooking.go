package appointmentBooking

import "github.com/gin-gonic/gin"

func InitModule(ginInstanceEngine *gin.Engine) {
	RegisterRoutes(ginInstanceEngine)
}

var AppointmentBookingDB = NewInMemoryDB()
