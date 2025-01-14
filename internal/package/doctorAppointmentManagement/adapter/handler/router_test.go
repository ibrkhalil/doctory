package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/ibrkhalil/doctory/internal/db"
	"github.com/stretchr/testify/assert"
)

func TestListAvailabilitySlots(t *testing.T) {
	db.GetInstance()
	defer db.Clear()

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	RegisterRoutes(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/doctor/appointments", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateAvailabilitySlot(t *testing.T) {
	db.GetInstance()
	defer db.Clear()

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	RegisterRoutes(router)

	// Bad request
	badReqBody := `{"doctor_id": "1", "time": "2023-10-10T10:00:00Z", "cost": 100}`
	req, _ := http.NewRequest(http.MethodPost, "/doctor/appointments", strings.NewReader(badReqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Good request
	goodReqBody := `{"doctorId": "1", "doctorName": "Ahmed3", "cost": 100, "time": "2025-01-13T11:20:11.564832+02:00", "isReserved": "false"}`
	goodReq, _ := http.NewRequest(http.MethodPost, "/doctor/appointments", strings.NewReader(goodReqBody))
	goodReq.Header.Set("Content-Type", "application/json")

	goodW := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, goodW.Code)
}
