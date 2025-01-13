package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	appointmentBookingService "github.com/ibrkhalil/doctory/internal/package/appointmentBooking/interfaces/api"
	doctorAppointmentManagementService "github.com/ibrkhalil/doctory/internal/package/doctorAppointmentManagement/core/service"
)

func Main() {
	r := gin.Default()
	appointmentBookingService.InitModule(r)
	doctorAppointmentManagementService.InitModule(r)

	// Add health endpoint
	r.GET("health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	})

	r.Run(":8080")
}
