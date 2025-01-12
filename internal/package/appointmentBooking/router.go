package appointmentBooking

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	appointmentGroup := router.Group("/appointments")
	{
		appointmentGroup.POST("/", createAppointment)
		appointmentGroup.GET("/", listAppointments)
	}
}

func createAppointment(ctx *gin.Context) {
	appointmentToSave, err := CreateAppointmentFromRequest(ctx)
	var response string
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "No available appointments found!"})
		return
	}

	if err != nil {
		log.Print("Failed creating an appointment")
	}

	response = "Succesfully created an appointment at " + appointmentToSave.StartingTime.Format(time.RFC822Z)
	ctx.JSON(http.StatusCreated, gin.H{"message": response})
}

func listAppointments(ctx *gin.Context) {
	appointmentList := ListAppointments()
	ctx.JSON(http.StatusOK, appointmentList)
}
