package usecases

import (
	"errors"
	"slices"
	"time"

	"github.com/LucasLCabral/statement-api/internal/models"
	"github.com/LucasLCabral/statement-api/internal/repositories"
)

type TransactionUsecaseInterface interface {
	CreateEvent(transaction *models.TransactionEvent) error
	GetAllTransactionsByUserIDAndAccountType(userID string, account models.AccountType) ([]models.TransactionEvent, error)
	GetStatement(userID string, account models.AccountType, currency models.CurrencyType, period string) ([]models.TransactionEvent, error)
}

type TransactionUsecase struct {
	transactionRepository repositories.TransactionRepositoryInterface
}

func NewTransactionUsecase(transactionRepository repositories.TransactionRepositoryInterface) TransactionUsecaseInterface {
	return &TransactionUsecase{transactionRepository: transactionRepository}
}

func (u *TransactionUsecase) CreateEvent(transaction *models.TransactionEvent) error {
	if transaction.ID == "" {
		return errors.New("id is required")
	}
	if transaction.UserID == "" {
		return errors.New("userID is required")
	}
	if transaction.Amount <= 0 {
		return errors.New("amount must be positive")
	}

	if !isValidTransactionType(transaction.Type) {
		return errors.New("Invalid transaction type: " + string(transaction.Type))
	}
	if !isValidAccountType(transaction.Account) {
		return errors.New("Invalid account type: " + string(transaction.Account))
	}
	if !isValidCurrencyType(transaction.Currency) {
		return errors.New("Invalid currency type: " + string(transaction.Currency))
	}
	if !isValidDirectionType(transaction.Direction) {
		return errors.New("Invalid direction type: " + string(transaction.Direction))
	}

	now := time.Now()
	if transaction.ProcessedAt.IsZero() {
		transaction.ProcessedAt = now
	}
	if transaction.CreatedAt.IsZero() {
		transaction.CreatedAt = now
	}

	if err := validateBusinessRules(transaction); err != nil {
		return err
	}

	return u.transactionRepository.CreateEvent(transaction)
}

func (u *TransactionUsecase) GetAllTransactionsByUserIDAndAccountType(userID string, account models.AccountType) ([]models.TransactionEvent, error) {
	transactions, err := u.transactionRepository.GetAllTransactionsByUserIDAndAccountType(userID, account)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (u *TransactionUsecase) GetStatement(userID string, account models.AccountType, currency models.CurrencyType, period string) ([]models.TransactionEvent, error) {
	statement, err := u.transactionRepository.GetStatement(userID, account, currency, period)
	if err != nil {
		return nil, err
	}

	return statement, nil
}

func isValidTransactionType(t models.TransactionType) bool {
	validTypes := []models.TransactionType{
		models.TransactionTypePIX,
		models.TransactionTypeTED,
		models.TransactionTypeCAMBIO,
		models.TransactionTypeCARTAO,
		models.TransactionTypeACAO,
		models.TransactionTypeWire,
	}

	return slices.Contains(validTypes, t)
}

func isValidAccountType(a models.AccountType) bool {
	validTypes := []models.AccountType{
		models.AccountTypeBrasileira,
		models.AccountTypeInvestimento,
		models.AccountTypeBanking,
	}

	for _, validType := range validTypes {
		if a == validType {
			return true
		}
	}
	return false
}

func isValidCurrencyType(c models.CurrencyType) bool {
	validTypes := []models.CurrencyType{
		models.CurrencyTypeBRL,
		models.CurrencyTypeUSD,
		models.CurrencyTypeEUR,
	}

	for _, validType := range validTypes {
		if c == validType {
			return true
		}
	}
	return false
}

func isValidDirectionType(d models.DirectionType) bool {
	validTypes := []models.DirectionType{
		models.DirectionTypeDebito,
		models.DirectionTypeCredito,
	}

	for _, validType := range validTypes {
		if d == validType {
			return true
		}
	}
	return false
}

func validateBusinessRules(transaction *models.TransactionEvent) error {
	if transaction.Type == models.TransactionTypePIX && transaction.Currency != models.CurrencyTypeBRL {
		return errors.New("PIX transactions must be in BRL currency")
	}

	if transaction.Account == models.AccountTypeBrasileira && transaction.Currency != models.CurrencyTypeBRL {
		return errors.New("brazilian account only accepts BRL currency")
	}

	if transaction.Type == models.TransactionTypeTED && transaction.Currency != models.CurrencyTypeBRL {
		return errors.New("TED transactions must be in BRL currency")
	}

	return nil
}
