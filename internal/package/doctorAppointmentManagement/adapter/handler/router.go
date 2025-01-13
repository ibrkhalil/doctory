package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ibrkhalil/doctory/internal/package/doctorAppointmentManagement/adapter/repository"
	"github.com/ibrkhalil/doctory/internal/package/doctorAppointmentManagement/core/utils"
)

func RegisterRoutes(router *gin.Engine) {
	doctorGroup := router.Group("/doctor")
	{
		doctorGroup.GET("/appointments", listAvailabilitySlots)
		doctorGroup.POST("/appointments", createAvailabilitySlot)
	}
}

func listAvailabilitySlots(ctx *gin.Context) {
	service := repository.NewDoctorAvailabilitySlotController()
	slots, err := service.ListAvailabilitySlots()
	if err != nil {
		errors.New("An error happened when listing appointments ")
	} else {
		ctx.JSON(http.StatusOK, slots)
	}
}

func createAvailabilitySlot(ctx *gin.Context) {
	service := repository.NewDoctorAvailabilitySlotController()
	req, err := utils.CreateAvailabilityFromRequest(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
	} else {
		err := service.AddAvailabilitySlot(req)
		if err != nil {
			ctx.JSON(http.StatusCreated, gin.H{"message": "Availability already taken!"})
		} else {
			ctx.JSON(http.StatusCreated, gin.H{"message": "Created availability time!"})
		}

	}
}
