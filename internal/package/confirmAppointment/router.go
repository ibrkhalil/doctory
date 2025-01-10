package confirmAppointment

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	appointmentGroup := router.Group("/appointments")
	{
		appointmentGroup.POST("/", addAppointment)
		appointmentGroup.GET("/:id", ConfirmAppointment)
	}
}

func addAppointment(c *gin.Context) {
	appointmentToConfirm := CreateAppointmentFromRequest(c)
	ConfirmAppointmentDB.ConfirmAppointment(appointmentToConfirm.ID)
	c.JSON(http.StatusCreated, appointmentToConfirm)
}

func ConfirmAppointment(c *gin.Context) {
	appointmentToAdd := CreateAppointmentFromRequest(c)
	ConfirmAppointmentDB.AddAppointment(appointmentToAdd)
	c.JSON(http.StatusCreated, appointmentToAdd)
}
