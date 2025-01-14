package usecases

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	controller := NewAppointmentBookingController()

	router.POST("/appointments", controller.CreateAppointment)
	router.GET("/appointments", controller.ListAppointments)

	return router
}

func TestCreateAppointment(t *testing.T) {
	router := setupTestRouter()

	t.Run("successful appointment creation", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/appointments", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)

		var response map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Contains(t, response["message"], "Succesfully created an appointment at")
	})
}

func TestListAppointments(t *testing.T) {
	router := setupTestRouter()

	t.Run("list appointments", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/appointments", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response []interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
	})
}
func TestNewAppointmentBookingController(t *testing.T) {
	controller := NewAppointmentBookingController()
	assert.NotNil(t, controller)
	assert.IsType(t, &AppointmentBookingController{}, controller)
}
