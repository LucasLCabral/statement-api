package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/LucasLCabral/statement-api/internal/models"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// Inicializar database
func initDB() {
	dsn := "host=localhost user=statement_user password=statement_pass dbname=statement_db port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Falha ao conectar com database:", err)
	}

	// Auto-migrate
	err = db.AutoMigrate(&models.TransactionEvent{})
	if err != nil {
		log.Fatal("Falha na migration:", err)
	}

	fmt.Println("✅ Database conectado e migrado com sucesso")
}

func main() {
	// Inicializar DB
	initDB()

	// Router
	r := mux.NewRouter()

	// Middleware
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}
			next.ServeHTTP(w, r)
		})
	})

	// Rotas básicas
	r.HandleFunc("/health", healthHandler).Methods("GET")
	r.HandleFunc("/", rootHandler).Methods("GET")

	fmt.Println("🚀 Statement API iniciando na porta 8080")
	fmt.Println("📍 Endpoints disponíveis:")
	fmt.Println("   GET  /health - Health check")
	fmt.Println("   GET  / - Root endpoint")

	log.Fatal(http.ListenAndServe(":8080", r))
}

