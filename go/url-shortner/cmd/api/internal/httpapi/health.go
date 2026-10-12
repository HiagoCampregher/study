package httpapi

import (
	"net/http"
)

func HealthLive(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, _ = w.Write([]byte(`{"status": "OK"}`))
}
