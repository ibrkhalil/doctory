package appointmentBooking

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	appointmentGroup := router.Group("/appointments")
	{
		appointmentGroup.POST("/", createAppointment)
		appointmentGroup.GET("/", listAppointments)
	}
}

func createAppointment(ctx *gin.Context) {
	appointmentToSave, err := CreateAppointmentFromRequest(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, appointmentToSave)
	}
	didCreateAppointment := CreateAppointment(appointmentToSave)
	var message string
	if didCreateAppointment {
		message = "Succesfully created an appointment"
	} else {
		message = "No available appointments found"
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": message})
}

func listAppointments(ctx *gin.Context) {
	appointmentList := ListAppointments()
	ctx.JSON(http.StatusOK, appointmentList)
}
