package handlers

import (
	"net/http"
	"time"

	"github.com/LucasLCabral/statement-api/internal/helper"
)

// Health check handler
func healthHandler(w http.ResponseWriter, r *http.Request) {
	helper.sendSuccess(w, map[string]string{
		"status":  "healthy",
		"time":    time.Now().Format(time.RFC3339),
		"service": "statement-api",
	}, "API is running")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	helper.sendSuccess(w, map[string]string{
		"status":  "healthy",
		"time":    time.Now().Format(time.RFC3339),
		"service": "statement-api",
	}, "Avenue statement API challenge let'ssssss GO(lang) 🤪")
}
