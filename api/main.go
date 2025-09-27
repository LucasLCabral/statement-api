package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/LucasLCabral/statement-api/internal/handlers"
	"github.com/LucasLCabral/statement-api/internal/models"
	"github.com/LucasLCabral/statement-api/internal/repositories"
	"github.com/LucasLCabral/statement-api/internal/usecases"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var repo repositories.TransactionRepositoryInterface
var usecase usecases.TransactionUsecaseInterface
var handler handlers.TransactionHandler

func initDB() {
	dsn := "host=localhost user=statement_user password=statement_pass dbname=statement_db port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = db.AutoMigrate(&models.TransactionEvent{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	fmt.Println("✅ Database connected and migrated successfully")
}

func main() {
	initDB()
	repo = repositories.NewTransactionRepository(db)
	usecase = usecases.NewTransactionUsecase(repo)
	handler = *handlers.NewTransactionHandler(usecase)

	r := mux.NewRouter()

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

	r.HandleFunc("/health", handlers.HealthHandler).Methods("GET")

	r.HandleFunc("/events", handler.CreateEvent).Methods("POST")
	r.HandleFunc("/statement/{userID}/{accountType}/{currencyType}/{period}", handler.GetStatement).Methods("GET")
	r.HandleFunc("/transactions/{userID}/{accountType}", handler.GetAllTransactionsByUserIDAndAccountType).Methods("GET")
	r.HandleFunc("/events/types", handler.GetEventTypes).Methods("GET")

	fmt.Println("💰💰💰 Statement API running on port 8080")
	fmt.Println("📍 Endpoints:")
	fmt.Println("   GET  /health - Health check")
	fmt.Println("   POST /events - Create event")
	fmt.Println("   GET  /statement/{userID}/{accountType}/{currencyType}/{period} - Get statement")
	fmt.Println("   GET  /transactions/{userID}/{accountType} - Get all transactions by user ID and account type")

	log.Fatal(http.ListenAndServe(":8080", r))
}
