package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ibrkhalil/doctory/internal/schema"
	"github.com/stretchr/testify/assert"
)

func TestValidateAppointmentFromRequest(t *testing.T) {
	tests := []struct {
		name     string
		slot     schema.DoctorAvailabilitySlot
		expected bool
	}{
		{
			name: "Valid appointment",
			slot: schema.DoctorAvailabilitySlot{
				ID:         "123",
				DoctorID:   "456",
				DoctorName: "Dr. Smith",
				Cost:       100,
				Time:       time.Now(),
			},
			expected: true,
		},
		{
			name: "Invalid appointment - missing cost",
			slot: schema.DoctorAvailabilitySlot{
				ID:         "123",
				DoctorID:   "456",
				DoctorName: "Dr. Smith",
				Time:       time.Now(),
			},
			expected: false,
		},
		{
			name: "Invalid appointment - missing doctor ID",
			slot: schema.DoctorAvailabilitySlot{
				ID:         "123",
				DoctorName: "Dr. Smith",
				Cost:       100,
				Time:       time.Now(),
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := validateAppointmentFromRequest(tt.slot)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestCreateAvailabilityFromRequest(t *testing.T) {
	tests := []struct {
		name          string
		requestBody   schema.DoctorAvailabilitySlot
		expectedError bool
	}{
		{
			name: "Valid request",
			requestBody: schema.DoctorAvailabilitySlot{
				DoctorID:   "456",
				DoctorName: "Mohamed Ramadan MD",
				Cost:       100,
				Time:       time.Now(),
			},
			expectedError: false,
		},
		{
			name: "Invalid request - missing cost",
			requestBody: schema.DoctorAvailabilitySlot{
				DoctorID:   "456",
				DoctorName: "Armin van Buuren MD",
				Time:       time.Now(),
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			jsonBody, _ := json.Marshal(tt.requestBody)
			c.Request = httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonBody))

			result, err := CreateAvailabilityFromRequest(c)
			if tt.expectedError {
				assert.Error(t, err)
				assert.Empty(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, result.ID)
				assert.Equal(t, tt.requestBody.DoctorID, result.DoctorID)
				assert.Equal(t, tt.requestBody.DoctorName, result.DoctorName)
				assert.Equal(t, tt.requestBody.Cost, result.Cost)
				assert.Equal(t, result.Time.Add(time.Hour), result.ToTime)
			}
		})
	}
}
