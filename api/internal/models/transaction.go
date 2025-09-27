package models

import (
	"time"

	"gorm.io/gorm"
)

type TransactionType string

const (
	TransactionTypePIX    TransactionType = "PIX"
	TransactionTypeTED    TransactionType = "TED"
	TransactionTypeCAMBIO TransactionType = "CAMBIO"
	TransactionTypeCARTAO TransactionType = "TRANSACAO DE CARTAO"
	TransactionTypeACAO   TransactionType = "ACAO"
	TransactionTypeWire   TransactionType = "WIRE"
)

type AccountType string

const (
	AccountTypeBrasileira   AccountType = "CONTA BRASILEIRA"
	AccountTypeInvestimento AccountType = "CONTA INVESTIMENTO"
	AccountTypeBanking      AccountType = "CONTA BANKING"
)

type CurrencyType string

const (
	CurrencyTypeBRL CurrencyType = "BRL"
	CurrencyTypeUSD CurrencyType = "USD"
	CurrencyTypeEUR CurrencyType = "EUR"
)

type DirectionType string

const (
	DirectionTypeDebito  DirectionType = "DEBITO"
	DirectionTypeCredito DirectionType = "CREDITO"
)

type TransactionMetadata struct {
	Description string `json:"description"`
	Source      string `json:"source,omitempty"`
	Reference   string `json:"reference,omitempty"`
}

type TransactionEvent struct {
	ID        string          `json:"id" gorm:"primaryKey;type:varchar(255)"`
	UserID    string          `json:"user_id" gorm:"index;not null;type:varchar(255)"`
	Account   AccountType     `json:"account" gorm:"index;not null;type:varchar(100)"`
	Currency  CurrencyType    `json:"currency" gorm:"index;not null;type:varchar(10)"`
	Type      TransactionType `json:"type" gorm:"not null;type:varchar(50)"`
	Direction DirectionType   `json:"direction" gorm:"not null;type:varchar(20)"`
	Amount    float64         `json:"amount" gorm:"not null"`
	Balance   float64         `json:"balance" gorm:"not null"`

	Description string `json:"description" gorm:"type:text"`
	Source      string `json:"source" gorm:"type:varchar(100)"`
	Reference   string `json:"reference" gorm:"type:varchar(255)"`

	ProcessedAt time.Time      `json:"processed_at" gorm:"not null"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

func (TransactionEvent) TableName() string {
	return "transaction_events"
}

func (t *TransactionEvent) BeforeCreate(tx *gorm.DB) error {
	now := time.Now()
	if t.ProcessedAt.IsZero() {
		t.ProcessedAt = now
	}
	if t.CreatedAt.IsZero() {
		t.CreatedAt = now
	}
	return nil
}

func (t *TransactionEvent) GetMetadata() TransactionMetadata {
	return TransactionMetadata{
		Description: t.Description,
		Source:      t.Source,
		Reference:   t.Reference,
	}
}

func (t *TransactionEvent) SetMetadata(metadata TransactionMetadata) {
	t.Description = metadata.Description
	t.Source = metadata.Source
	t.Reference = metadata.Reference
}

type TransactionEventRequest struct {
	ID          string              `json:"id"`
	UserID      string              `json:"user_id"`
	Account     AccountType         `json:"account"`
	Currency    CurrencyType        `json:"currency"`
	Type        TransactionType     `json:"type"`
	Direction   DirectionType       `json:"direction"`
	Amount      float64             `json:"amount"`
	Balance     float64             `json:"balance"`
	Metadata    TransactionMetadata `json:"metadata"`
	ProcessedAt *time.Time          `json:"processed_at,omitempty"`
	CreatedAt   *time.Time          `json:"created_at,omitempty"`
}

func (req *TransactionEventRequest) ToTransactionEvent() TransactionEvent {
	now := time.Now()

	event := TransactionEvent{
		ID:          req.ID,
		UserID:      req.UserID,
		Account:     req.Account,
		Currency:    req.Currency,
		Type:        req.Type,
		Direction:   req.Direction,
		Amount:      req.Amount,
		Balance:     req.Balance,
		Description: req.Metadata.Description,
		Source:      req.Metadata.Source,
		Reference:   req.Metadata.Reference,
		ProcessedAt: now,
		CreatedAt:   now,
	}

	if req.ProcessedAt != nil && !req.ProcessedAt.IsZero() {
		event.ProcessedAt = *req.ProcessedAt
	}
	if req.CreatedAt != nil && !req.CreatedAt.IsZero() {
		event.CreatedAt = *req.CreatedAt
	}

	return event
}

type TransactionEventResponse struct {
	ID          string              `json:"id"`
	UserID      string              `json:"user_id"`
	Account     AccountType         `json:"account"`
	Currency    CurrencyType        `json:"currency"`
	Type        TransactionType     `json:"type"`
	Direction   DirectionType       `json:"direction"`
	Amount      float64             `json:"amount"`
	Balance     float64             `json:"balance"`
	Metadata    TransactionMetadata `json:"metadata"`
	ProcessedAt time.Time           `json:"processed_at"`
	CreatedAt   time.Time           `json:"created_at"`
}

type StatementResponse struct {
	ID           string                     `json:"id"`
	Account      AccountType                `json:"account"`
	Currency     CurrencyType               `json:"currency"`
	Type         TransactionType            `json:"type"`
	Period       string                     `json:"period"`
	Transactions []TransactionEventResponse `json:"transactions"`
}

func (t *TransactionEvent) ToResponse() TransactionEventResponse {
	return TransactionEventResponse{
		ID:          t.ID,
		UserID:      t.UserID,
		Account:     t.Account,
		Currency:    t.Currency,
		Type:        t.Type,
		Direction:   t.Direction,
		Amount:      t.Amount,
		Balance:     t.Balance,
		Metadata:    t.GetMetadata(),
		ProcessedAt: t.ProcessedAt,
		CreatedAt:   t.CreatedAt,
	}
}
