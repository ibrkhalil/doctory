package doctorAvailability

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	doctorGroup := router.Group("/doctor")
	{
		doctorGroup.GET("/appointments", getAppointments)
		doctorGroup.POST("/appointments", createAppointment)
		doctorGroup.PUT("/appointments/:id", updateAppointment)
		doctorGroup.DELETE("/appointments/:id", deleteAppointment)
	}
}

func getAppointments(ctx *gin.Context) {
	AvailabilityDB.ListAvailability()
	ctx.JSON(http.StatusOK, gin.H{"message": "Get all appointments"})
}

func createAppointment(ctx *gin.Context) {
	req := CreateAvailabilityFromRequest(ctx)
	AvailabilityDB.AddAvailability(req)
	ctx.JSON(http.StatusCreated, gin.H{"message": "Appointment created"})
}

func updateAppointment(ctx *gin.Context) {
	req := CreateAvailabilityFromRequest(ctx)
	AvailabilityDB.UpdateAvailability(req)
	ctx.JSON(http.StatusOK, gin.H{"message": "Appointment updated"})
}

func deleteAppointment(ctx *gin.Context) {
	req := CreateAvailabilityFromRequest(ctx)
	AvailabilityDB.DeleteAvailability(req.ID)
	ctx.JSON(http.StatusOK, gin.H{"message": "Appointment deleted"})
}
