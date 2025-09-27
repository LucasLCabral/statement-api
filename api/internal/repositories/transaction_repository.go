package repositories

import (
	"time"

	"github.com/LucasLCabral/statement-api/internal/models"
	"gorm.io/gorm"
)

type TransactionRepositoryInterface interface {
	CreateEvent(event *models.TransactionEvent) error
	GetAllTransactionsByUserIDAndAccountType(userID string, account models.AccountType) ([]models.TransactionEvent, error)
	GetStatement(userID string, account models.AccountType, currency models.CurrencyType, period string) ([]models.TransactionEvent, error)
}

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepositoryInterface {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) CreateEvent(event *models.TransactionEvent) error {
	return r.db.Create(event).Error
}

func (r *TransactionRepository) GetAllTransactionsByUserIDAndAccountType(userID string, account models.AccountType) ([]models.TransactionEvent, error) {
	var transactions []models.TransactionEvent
	err := r.db.Where("user_id = ? AND account = ?", userID, string(account)).
		Order("created_at DESC").
		Find(&transactions).Error
	return transactions, err
}

func (r *TransactionRepository) GetStatement(userID string, account models.AccountType, currency models.CurrencyType, period string) ([]models.TransactionEvent, error) {
	var transactions []models.TransactionEvent

	startDate := parsePeriodToTime(period)

	err := r.db.Where("user_id = ? AND account = ? AND currency = ? AND created_at >= ?",
		userID,
		string(account),
		string(currency),
		startDate).
		Order("created_at DESC").
		Find(&transactions).Error

	return transactions, err
}

func parsePeriodToTime(period string) time.Time {
	now := time.Now()
	switch period {
	case "7d":
		return now.AddDate(0, 0, -7)
	case "30d":
		return now.AddDate(0, 0, -30)
	case "90d":
		return now.AddDate(0, 0, -90)
	case "1y":
		return now.AddDate(-1, 0, 0)
	default:
		return now.AddDate(0, 0, -30)
	}
}
