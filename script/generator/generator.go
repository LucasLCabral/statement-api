package generator

import (
	"fmt"
	"math/rand"
	"time"

	"statement-validator/model"
)

var Scenarios = map[string]model.TransactionScenario{
	"brasileira": {
		Account:    model.AccountTypeBrasileira,
		Currency:   model.CurrencyTypeBRL,
		Types:      []model.TransactionType{model.TransactionTypePIX, model.TransactionTypeTED, model.TransactionTypeCAMBIO},
		Directions: []model.DirectionType{model.DirectionTypeCredito, model.DirectionTypeDebito},
	},
	"investimento": {
		Account:    model.AccountTypeInvestimento,
		Currency:   model.CurrencyTypeUSD,
		Types:      []model.TransactionType{model.TransactionTypeCAMBIO, model.TransactionTypeACAO},
		Directions: []model.DirectionType{model.DirectionTypeCredito, model.DirectionTypeDebito},
	},
	"banking_usd": {
		Account:    model.AccountTypeBanking,
		Currency:   model.CurrencyTypeUSD,
		Types:      []model.TransactionType{model.TransactionTypeCAMBIO, model.TransactionTypeCARTAO, model.TransactionTypeWire},
		Directions: []model.DirectionType{model.DirectionTypeCredito, model.DirectionTypeDebito},
	},
	"banking_eur": {
		Account:    model.AccountTypeBanking,
		Currency:   model.CurrencyTypeEUR,
		Types:      []model.TransactionType{model.TransactionTypeCAMBIO, model.TransactionTypeCARTAO, model.TransactionTypeWire},
		Directions: []model.DirectionType{model.DirectionTypeCredito, model.DirectionTypeDebito},
	},
}

var UserIDs = []string{
	"user-001", "user-002", "user-003", "user-004", "user-005",
	"user-006", "user-007", "user-008", "user-009", "user-010",
	"user-011", "user-012", "user-013", "user-014", "user-015",
}

func GenerateRandomAmount(transactionType model.TransactionType, direction model.DirectionType) float64 {
	baseAmounts := map[model.TransactionType]map[model.DirectionType][2]float64{
		model.TransactionTypePIX: {
			model.DirectionTypeCredito: {10.0, 5000.0},
			model.DirectionTypeDebito:  {5.0, 3000.0},
		},
		model.TransactionTypeTED: {
			model.DirectionTypeCredito: {100.0, 50000.0},
			model.DirectionTypeDebito:  {200.0, 100000.0},
		},
		model.TransactionTypeCAMBIO: {
			model.DirectionTypeCredito: {50.0, 10000.0},
			model.DirectionTypeDebito:  {100.0, 15000.0},
		},
		model.TransactionTypeCARTAO: {
			model.DirectionTypeCredito: {5.0, 2000.0},
			model.DirectionTypeDebito:  {10.0, 5000.0},
		},
		model.TransactionTypeACAO: {
			model.DirectionTypeCredito: {100.0, 10000.0},
			model.DirectionTypeDebito:  {150.0, 15000.0},
		},
		model.TransactionTypeWire: {
			model.DirectionTypeCredito: {500.0, 100000.0},
			model.DirectionTypeDebito:  {1000.0, 200000.0},
		},
	}

	if amounts, exists := baseAmounts[transactionType]; exists {
		if directionAmounts, exists := amounts[direction]; exists {
			return directionAmounts[0] + rand.Float64()*(directionAmounts[1]-directionAmounts[0])
		}
	}
	return 10.0 + rand.Float64()*1000.0
}

func GenerateTransaction() model.TransactionEvent {
	scenarioNames := make([]string, 0, len(Scenarios))
	for name := range Scenarios {
		scenarioNames = append(scenarioNames, name)
	}
	scenario := Scenarios[scenarioNames[rand.Intn(len(scenarioNames))]]

	transactionType := scenario.Types[rand.Intn(len(scenario.Types))]
	direction := scenario.Directions[rand.Intn(len(scenario.Directions))]
	amount := GenerateRandomAmount(transactionType, direction)
	userID := UserIDs[rand.Intn(len(UserIDs))]
	now := time.Now()

	return model.TransactionEvent{
		ID:          fmt.Sprintf("txn_%d_%s", time.Now().UnixNano(), userID),
		UserID:      userID,
		Account:     scenario.Account,
		Currency:    scenario.Currency,
		Type:        transactionType,
		Direction:   direction,
		Amount:      amount,
		Balance:     amount + rand.Float64()*10000,
		ProcessedAt: now.Add(time.Duration(rand.Intn(100)) * time.Millisecond),
		CreatedAt:   now,
	}
}
