package doctorAppointmentManagement

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	doctorGroup := router.Group("/doctor")
	{
		doctorGroup.GET("/appointments", listAvailabilitySlots)
		doctorGroup.POST("/appointments", createAvailabilitySlot)
	}
}

func listAvailabilitySlots(ctx *gin.Context) {
	slots, err := ListAvailabilitySlots()
	if err != nil {
		errors.New("An error happened when listing appointments ")
	} else {
		ctx.JSON(http.StatusOK, slots)
	}
}

func createAvailabilitySlot(ctx *gin.Context) {
	req, err := CreateAvailabilityFromRequest(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
	} else {
		err := AddAvailabilitySlot(req)
		if err != nil {
			ctx.JSON(http.StatusCreated, gin.H{"message": "Availability already taken!"})
		} else {
			ctx.JSON(http.StatusCreated, gin.H{"message": "Created availability time!"})
		}

	}
}
