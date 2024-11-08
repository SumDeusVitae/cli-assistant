package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func registerHandler(w http.ResponseWriter, r *http.Request) {
	// Simulate user registration process
	userID := uuid.New().String() // Generates a unique user ID
	apiKey := uuid.New().String() // Generates a unique API Key

	// You would typically save the user info in a database here

	// Send a JSON response back to the CLI tool
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"user_id": "%s", "api_key": "%s"}`, userID, apiKey)
}
