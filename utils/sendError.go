package utils

import (
	"encoding/json"
	"net/http"
)

func SendError(w http.ResponseWriter, status int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	errorResponse := map[string]string{"error": err.Error()}
	json.NewEncoder(w).Encode(errorResponse)
}
