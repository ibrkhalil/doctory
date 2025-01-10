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

func createAppointment(ctx *gin.Context) {
	appointmentToSave := CreateAppointmentFromRequest(ctx)
	AppointmentSlotsDB.CreateAppointment(appointmentToSave)
	ctx.JSON(http.StatusCreated, appointmentToSave)
}

func getAppointment(ctx *gin.Context) {
	appointmentToReturn := CreateAppointmentFromRequest(ctx)
	AppointmentSlotsDB.GetAppointment(appointmentToReturn.ID)
	ctx.JSON(http.StatusCreated, appointmentToReturn)
}

func updateAppointment(ctx *gin.Context) {
	updatedAppointment := CreateAppointmentFromRequest(ctx)
	AppointmentSlotsDB.UpdateAppointment(updatedAppointment.ID, updatedAppointment)
	ctx.JSON(http.StatusCreated, updatedAppointment)
}

func deleteAppointment(ctx *gin.Context) {
	deletedAppointment := CreateAppointmentFromRequest(ctx)
	AppointmentSlotsDB.DeleteAppointment(deletedAppointment.ID)
	ctx.JSON(http.StatusCreated, deletedAppointment)
}
