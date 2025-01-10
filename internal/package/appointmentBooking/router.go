package appointmentBooking

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	appointmentGroup := router.Group("/appointments")
	{
		appointmentGroup.POST("/", createAppointment)
		appointmentGroup.GET("/:id", getAppointment)
		appointmentGroup.PUT("/:id", updateAppointment)
		appointmentGroup.DELETE("/:id", deleteAppointment)
	}
}

func createAppointment(c *gin.Context) {
	appointmentToSave := CreateAppointmentFromRequest(c)
	AppointmentBookingDB.CreateAppointment(appointmentToSave)
	c.JSON(http.StatusCreated, appointmentToSave)
}

func getAppointment(c *gin.Context) {
	appointmentToReturn := CreateAppointmentFromRequest(c)
	AppointmentBookingDB.GetAppointment(appointmentToReturn.ID)
	c.JSON(http.StatusCreated, appointmentToReturn)
}

func updateAppointment(c *gin.Context) {
	updatedAppointment := CreateAppointmentFromRequest(c)
	AppointmentBookingDB.UpdateAppointment(updatedAppointment.ID, updatedAppointment)
	c.JSON(http.StatusCreated, updatedAppointment)
}

func deleteAppointment(c *gin.Context) {
	deletedAppointment := CreateAppointmentFromRequest(c)
	AppointmentBookingDB.DeleteAppointment(deletedAppointment.ID)
	c.JSON(http.StatusCreated, deletedAppointment)
}
