package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ibrkhalil/doctory/internal/package/appointmentBooking/usecases"
)

func RegisterRoutes(router *gin.Engine) {
	service := usecases.NewAppointmentBookingController()

	appointmentGroup := router.Group("/appointments")
	{
		appointmentGroup.POST("/", service.CreateAppointment)
		appointmentGroup.GET("/", service.ListAppointments)
	}
}

func InitModule(ginEngineInstance *gin.Engine) {
	RegisterRoutes(ginEngineInstance)
}
