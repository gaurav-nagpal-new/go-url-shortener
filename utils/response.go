package utils

import (
	"encoding/json"
	"net/http"
)

func SendResponse(val any, statusCode int, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(&val)
}
