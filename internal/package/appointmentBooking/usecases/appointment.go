package usecases

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ibrkhalil/doctory/internal/package/appointmentBooking"
	"github.com/ibrkhalil/doctory/internal/package/appointmentBooking/infrastructure"
)

type AppointmentBookingController struct {
	service *infrastructure.AppointmeentStorage
}

func NewTodoController() *AppointmentBookingController {
	return &AppointmentBookingController{}
}

func (c *AppointmentBookingController) CreateAppointment(ctx *gin.Context) {
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

func (c *AppointmentBookingController) ListAppointments(ctx *gin.Context) {
	appointmentList := infrastructure.ListAppointments()
	ctx.JSON(http.StatusOK, appointmentList)
}
