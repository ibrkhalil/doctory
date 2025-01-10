package doctorAvailability

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	doctorAvailabilityGroup := router.Group("/doctorAvailability")
	{
		doctorAvailabilityGroup.GET("/", getAllDoctorAvailabilities)
		doctorAvailabilityGroup.GET("/:id", getDoctorAvailabilityByID)
		doctorAvailabilityGroup.POST("/", createDoctorAvailability)
		doctorAvailabilityGroup.PUT("/:id", updateDoctorAvailability)
		doctorAvailabilityGroup.DELETE("/:id", deleteDoctorAvailability)
	}
}

func getAllDoctorAvailabilities(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get all doctor availabilities"})
}

func getDoctorAvailabilityByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Get doctor availability by ID", "id": id})
}

func createDoctorAvailability(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "Create new doctor availability"})
}

func updateDoctorAvailability(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Update doctor availability", "id": id})
}

func deleteDoctorAvailability(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Delete doctor availability", "id": id})
}
