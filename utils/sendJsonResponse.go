package utils

import (
	"encoding/json"
	"net/http"
)

// wrapper for json.NewEncoder().Encode()
func JSON(w http.ResponseWriter, response interface{}) {
	json.NewEncoder(w).Encode(response)
}
