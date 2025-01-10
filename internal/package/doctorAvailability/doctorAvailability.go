package doctorAvailability

import "github.com/gin-gonic/gin"

func InitModule(ginEngineInstance *gin.Engine) {
	RegisterRoutes(ginEngineInstance)
}

var AvailabilityDB = NewInMemoryAvailabilityDB()
