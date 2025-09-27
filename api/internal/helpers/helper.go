package helper

import (
	"encoding/json"
	"net/http"

	"github.com/LucasLCabral/statement-api/internal/models"
)

// Helpers de resposta
func sendJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func SendError(w http.ResponseWriter, statusCode int, errorMsg string) {
	sendJSON(w, statusCode, models.APIResponse{
		Success: false,
		Error:   errorMsg,
	})
}

func SendSuccess(w http.ResponseWriter, data interface{}, message string) {
	sendJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Data:    data,
		Message: message,
	})
}