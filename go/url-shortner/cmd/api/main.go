package main

import (
	"fmt"
	"net/http"
	"url-shortener/cmd/api/internal/httpapi"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health/live", httpapi.HealthLive)

	fmt.Println("Starting server on :8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
