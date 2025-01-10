package app

import (
	"github.com/gin-gonic/gin"
	"github.com/ibrkhalil/doctory/internal/package/appointmentBooking"
	"github.com/ibrkhalil/doctory/internal/package/doctorAvailability"
)

func Main() {
	r := gin.Default()
	appointmentBooking.Start(r)
	doctorAvailability.Start(r)
	r.Run(":8080")
}
