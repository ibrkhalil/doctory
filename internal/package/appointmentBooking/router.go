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
	c.JSON(http.StatusCreated, gin.H{"message": "Appointment created"})
}

func getAppointment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Appointment details"})
}

func updateAppointment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Appointment updated"})
}

func deleteAppointment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Appointment deleted"})
}
