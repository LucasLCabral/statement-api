package model

import "time"

// Tipos conforme especificação do desafio
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
	ID          string              `json:"id" db:"id"`
	UserID      string              `json:"user_id" db:"user_id"`
	Account     AccountType         `json:"account" db:"account_type"`
	Currency    CurrencyType        `json:"currency" db:"currency_type"`
	Type        TransactionType     `json:"type" db:"transaction_type"`
	Direction   DirectionType       `json:"direction" db:"direction_type"`
	Amount      float64             `json:"amount" db:"amount"`
	Balance     float64             `json:"balance" db:"balance"`
	Metadata    TransactionMetadata `json:"metadata" db:"metadata"`
	ProcessedAt time.Time           `json:"processed_at" db:"processed_at"`
	CreatedAt   time.Time           `json:"created_at" db:"created_at"`
}

type TransactionScenario struct {
	Account    AccountType
	Currency   CurrencyType
	Types      []TransactionType
	Directions []DirectionType
}
