package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/adon988/go_api_example/api/controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Mock struct for MemberController
type MockMemberController struct{}

// GetMmeberInfo mock implementation
func (m *MockMemberController) GetMmeberInfo(c *gin.Context) {
	// Create a mock member info response
	mockData := controllers.MemberinfoResponse{
		ID:        "123456",
		Name:      "test",
		Birthday:  "2021-01-01",
		Email:     "example@example.com",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	mockResponse := controllers.GetMemberResonse{
		Code: 0,
		Data: mockData,
		Msg:  "success",
	}
	// Convert the mock response to JSON
	mockJSON, _ := json.Marshal(mockResponse)
	// Set the response body with the mock JSON
	c.String(http.StatusOK, string(mockJSON))
}

func TestGetMemberInfo(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a new Gin router
	router := gin.New()

	// Create a new instance of the MockMemberController
	controller := &MockMemberController{}

	// Define a route for the GetMmeberInfo handler
	router.GET("/member", controller.GetMmeberInfo)

	// Create a new HTTP request to the /member endpoint
	req, _ := http.NewRequest("GET", "/member", nil)

	// Create a new HTTP response recorder
	res := httptest.NewRecorder()

	// Serve the HTTP request
	router.ServeHTTP(res, req)

	// Assert that the response status code is 200 OK
	assert.Equal(t, http.StatusOK, res.Code)

	// Assert that the response body contains the expected data
	expectedData := controllers.MemberinfoResponse{
		ID:        "123456",
		Name:      "test",
		Birthday:  "2021-01-01",
		Email:     "example@example.com",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	expectedResponse := controllers.GetMemberResonse{
		Code: 0,
		Data: expectedData,
		Msg:  "success",
	}
	expectedJSON, _ := json.Marshal(expectedResponse)
	assert.JSONEq(t, string(expectedJSON), res.Body.String())
}
