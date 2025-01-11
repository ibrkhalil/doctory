package appointmentBooking

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	appointmentGroup := router.Group("/appointments")
	{
		appointmentGroup.POST("/", createAppointment)
		appointmentGroup.GET("/:id", listAppointments)
	}
}

func createAppointment(ctx *gin.Context) {
	appointmentToSave, err := CreateAppointmentFromRequest(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, appointmentToSave)
	}
	CreateAppointment(appointmentToSave)
	ctx.JSON(http.StatusCreated, appointmentToSave)
}

func listAppointments(ctx *gin.Context) {
	appointmentList := ListAppointments()
	ctx.JSON(http.StatusCreated, appointmentList)
}
