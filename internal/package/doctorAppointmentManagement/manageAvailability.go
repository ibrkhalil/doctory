package doctorAppointmentManagement

import "github.com/gin-gonic/gin"

func InitModule() {
	NewInMemoryDB()
	r := gin.Default()
	RegisterRoutes(r)
}
