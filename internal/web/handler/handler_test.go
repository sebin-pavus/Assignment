package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sebin-pavus/Assignment/internal/model"
	"github.com/sebin-pavus/Assignment/internal/web/handler"
)

// Test that a POST request succeeds returns 200 OK
func TestPostComputeSuccess(t *testing.T) {
	// Create a response recorder
	w := httptest.NewRecorder()

	// Get a new router
	r := gin.Default()

	// Define the route similar to its definition in the routes file
	r.POST("/compute", handler.PostCompute)

	// Create a request to send to the above route
	payload := GetPayload(1, 1)
	var buf bytes.Buffer
	_ = json.NewEncoder(&buf).Encode(payload)
	req, _ := http.NewRequest("POST", "/compute", &buf)
	req.Header.Add("Content-Type", "application/json")

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	// Test that the http status code is 200
	if w.Code != http.StatusOK {
		t.Fail()
	}

}

// Test that a POST request with invalid input give 400 Bad Request
func TestPostComputeInvalidInput(t *testing.T) {

	// Create a response recorder
	w := httptest.NewRecorder()

	// Get a new router
	r := gin.Default()

	// Define the route similar to its definition in the routes file
	r.POST("/compute", handler.PostCompute)

	// Create a request to send to the above route
	var output model.Output
	var buf bytes.Buffer
	_ = json.NewEncoder(&buf).Encode(output)
	req, _ := http.NewRequest("POST", "/compute", &buf)
	req.Header.Add("Content-Type", "application/json")

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fail()
	}
}

// Test that a POST request with 0 as input give 400 Bad Request
func TestPostComputeDivisonByZeroFailure(t *testing.T) {

	// Create a response recorder
	w := httptest.NewRecorder()

	// Get a new router
	r := gin.Default()

	// Define the route similar to its definition in the routes file
	r.POST("/compute", handler.PostCompute)

	// Create a request to send to the above route
	payload := GetPayload(1, 0)
	var buf bytes.Buffer
	_ = json.NewEncoder(&buf).Encode(payload)
	req, _ := http.NewRequest("POST", "/compute", &buf)
	req.Header.Add("Content-Type", "application/json")

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fail()
	}
}

func GetPayload(A float64, B float64) model.Input {
	var payload model.Input
	payload.A = A
	payload.B = B
	return payload
}
