package utils

import (
	"encoding/json"
	"net/http"
)

func Message(w http.ResponseWriter, message string) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"message": message})
}
