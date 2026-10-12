package httpapi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthLive(t *testing.T) {
	request := httptest.NewRequest(
		http.MethodGet,
		"/health/live",
		nil,
	)

	response := httptest.NewRecorder()

	HealthLive(response, request)

	result := response.Result()
	defer result.Body.Close()

	if result.StatusCode != http.StatusOK {
		t.Errorf(
			"expected status %d, got %d",
			http.StatusOK,
			result.StatusCode,
		)
	}

	if result.Header.Get("Content-Type") != "application/json" {
		t.Errorf(
			"expected Content-Type application/json, got %s",
			result.Header.Get("Content-Type"),
		)
	}

	var body struct {
		Status string `json:"status"`
	}

	err := json.NewDecoder(result.Body).Decode(&body)
	if err != nil {
		t.Fatalf("could not decode response body: %v", err)
	}

	if body.Status != "OK" {
		t.Errorf("expected status OK, got %s", body.Status)
	}
}
