package Helpers

import (
	"encoding/json"
	"net/http"
)

// RespondWithJSON sends a JSON response with the specified status code
func RespondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
