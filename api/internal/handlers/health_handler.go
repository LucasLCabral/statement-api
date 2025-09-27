package handlers

import (
	"net/http"
	"time"

	"github.com/LucasLCabral/statement-api/internal/helpers"
)

// Health check handler
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	helper.SendSuccess(w, map[string]string{
		"status":  "healthy",
		"time":    time.Now().Format(time.RFC3339),
		"service": "statement-api",
	}, "API is running")
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	helper.SendSuccess(w, map[string]string{
		"status":  "healthy",
		"time":    time.Now().Format(time.RFC3339),
		"service": "statement-api",
	}, "Avenue statement API challenge let'ssssss GO(lang) 🤪")
}
