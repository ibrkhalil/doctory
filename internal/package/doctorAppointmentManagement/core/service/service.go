package service

import (
	"github.com/gin-gonic/gin"
	"github.com/ibrkhalil/doctory/internal/package/doctorAppointmentManagement/adapter/handler"
)

func InitModule(ginEngineInstance *gin.Engine) {
	handler.RegisterRoutes(ginEngineInstance)
}
