package doctorAvailabilitySlot

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	doctorGroup := router.Group("/doctor")
	{
		doctorGroup.GET("/appointments", listAppointmentSlots)
		doctorGroup.POST("/appointments", createAppointmentSlot)
	}
}

func listAppointmentSlots(ctx *gin.Context) {
	slots, err := AvailabilityDB.ListAppointmentSlots()
	if err != nil {
		errors.New("An error happened when listing appointments ")
	}
	ctx.JSON(http.StatusOK, slots)
}

func createAppointmentSlot(ctx *gin.Context) {
	req, err := CreateAvailabilityFromRequest(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
	} else {
		AvailabilityDB.AddAvailabilitySlot(req)
		ctx.JSON(http.StatusCreated, gin.H{"message": "Appointment created"})
	}
}
