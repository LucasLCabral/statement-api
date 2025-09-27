package handlers

import (
	"encoding/json"
	"net/http"
	"net/url"

	helper "github.com/LucasLCabral/statement-api/internal/helpers"
	"github.com/LucasLCabral/statement-api/internal/models"
	"github.com/LucasLCabral/statement-api/internal/usecases"
	"github.com/gorilla/mux"
)

type TransactionHandler struct {
	usecase usecases.TransactionUsecaseInterface
}

func NewTransactionHandler(usecase usecases.TransactionUsecaseInterface) *TransactionHandler {
	return &TransactionHandler{usecase: usecase}
}

// POST /events
func (h *TransactionHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helper.SendError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var request models.TransactionEventRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		helper.SendError(w, http.StatusBadRequest, "Invalid JSON: "+err.Error())
		return
	}

	// Converter para TransactionEvent usando o método ToTransactionEvent
	event := request.ToTransactionEvent()

	if err := h.usecase.CreateEvent(&event); err != nil {
		helper.SendError(w, http.StatusInternalServerError, "Failed to create event: "+err.Error())
		return
	}

	helper.SendSuccess(w, event, "Event created successfully")
}

// GET /statement/{userID}/{accountType}/{currencyType}/{period}
func (h *TransactionHandler) GetStatement(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.SendError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	vars := mux.Vars(r)
	userID := vars["userID"]
	accountTypeEncoded := vars["accountType"]
	currencyTypeEncoded := vars["currencyType"]
	period := vars["period"]

	if userID == "" {
		helper.SendError(w, http.StatusBadRequest, "userID is required")
		return
	}

	// URL decode para espaços
	accountType, err := url.QueryUnescape(accountTypeEncoded)
	if err != nil {
		helper.SendError(w, http.StatusBadRequest, "Invalid accountType format")
		return
	}

	currencyType, err := url.QueryUnescape(currencyTypeEncoded)
	if err != nil {
		helper.SendError(w, http.StatusBadRequest, "Invalid currencyType format")
		return
	}

	if period == "" {
		helper.SendError(w, http.StatusBadRequest, "period is required")
		return
	}

	account := models.AccountType(accountType)
	currency := models.CurrencyType(currencyType)

	transactions, err := h.usecase.GetStatement(userID, account, currency, period)
	if err != nil {
		helper.SendError(w, http.StatusInternalServerError, "Failed to get statement: "+err.Error())
		return
	}

	helper.SendSuccess(w, transactions, "Statement retrieved successfully")
}

// GET /transactions/{userID}/{accountType}
func (h *TransactionHandler) GetAllTransactionsByUserIDAndAccountType(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.SendError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	vars := mux.Vars(r)
	userID := vars["userID"]
	accountTypeEncoded := vars["accountType"]

	if userID == "" {
		helper.SendError(w, http.StatusBadRequest, "userID is required")
		return
	}

	accountType, err := url.QueryUnescape(accountTypeEncoded)
	if err != nil {
		helper.SendError(w, http.StatusBadRequest, "Invalid accountType format")
		return
	}

	account := models.AccountType(accountType)

	transactions, err := h.usecase.GetAllTransactionsByUserIDAndAccountType(userID, account)
	if err != nil {
		helper.SendError(w, http.StatusInternalServerError, "Failed to get transactions: "+err.Error())
		return
	}

	helper.SendSuccess(w, transactions, "Transactions retrieved successfully")
}

// GET /events/types
func (h *TransactionHandler) GetEventTypes(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helper.SendError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	types := map[string]interface{}{
		"transaction_types": []models.TransactionType{
			models.TransactionTypePIX,
			models.TransactionTypeTED,
			models.TransactionTypeCAMBIO,
			models.TransactionTypeCARTAO,
			models.TransactionTypeACAO,
			models.TransactionTypeWire,
		},
		"account_types": []models.AccountType{
			models.AccountTypeBrasileira,
			models.AccountTypeInvestimento,
			models.AccountTypeBanking,
		},
		"currency_types": []models.CurrencyType{
			models.CurrencyTypeBRL,
			models.CurrencyTypeUSD,
			models.CurrencyTypeEUR,
		},
		"direction_types": []models.DirectionType{
			models.DirectionTypeDebito,
			models.DirectionTypeCredito,
		},
	}

	helper.SendSuccess(w, types, "Event types retrieved successfully")
}
