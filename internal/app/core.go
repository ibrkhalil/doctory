package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	appointmentBookingApi "github.com/ibrkhalil/doctory/internal/package/appointmentBooking/interfaces/api"
	"github.com/ibrkhalil/doctory/internal/package/doctorAppointmentManagement"
)

func Main() {
	r := gin.Default()
	appointmentBookingApi.InitModule(r)
	doctorAppointmentManagement.Start(r)

	// Add health endpoint
	r.GET("health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	})

	r.Run(":8080")
}
