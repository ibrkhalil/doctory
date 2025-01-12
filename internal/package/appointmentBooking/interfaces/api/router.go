package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ibrkhalil/doctory/internal/package/appointmentBooking"
	"github.com/ibrkhalil/doctory/internal/package/appointmentBooking/infrastructure"
)

func RegisterRoutes(router *gin.Engine) {
	appointmentGroup := router.Group("/appointments")
	{
		appointmentGroup.POST("/", createAppointment)
		appointmentGroup.GET("/", listAppointments)
	}
}

func createAppointment(ctx *gin.Context) {
	appointmentToSave, err := appointmentBooking.CreateAppointmentFromRequest(ctx)
	var response string
	if err != nil {
		response = "No available appointments found!"
		ctx.JSON(http.StatusConflict, gin.H{"message": response})
		return
	}

	response = "Succesfully created an appointment at " + appointmentToSave.StartingTime.Format(time.RFC822Z)
	ctx.JSON(http.StatusCreated, gin.H{"message": response})
}

func listAppointments(ctx *gin.Context) {
	appointmentList := infrastructure.ListAppointments()
	ctx.JSON(http.StatusOK, appointmentList)
}

func InitModule(ginEngineInstance *gin.Engine) {
	RegisterRoutes(ginEngineInstance)
}
