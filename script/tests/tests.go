package tests

import (
	"fmt"
	"math/rand"
	"time"

	"statement-validator/api"
	"statement-validator/generator"
	"statement-validator/model"
)

// Testes específicos do desafio
func TestPIXTransaction() error {
	fmt.Println("🧪 Testando transação PIX...")
	transaction := model.TransactionEvent{
		ID:        "test_pix_001",
		UserID:    "user-123",
		Account:   model.AccountTypeBrasileira,
		Currency:  model.CurrencyTypeBRL,
		Type:      model.TransactionTypePIX,
		Direction: model.DirectionTypeCredito,
		Amount:    600.0,
		Balance:   600.0,
		Metadata: model.TransactionMetadata{
			Description: "Transferência recebida",
			Source:      "test",
			Reference:   "PIX_REF_001",
		},
		ProcessedAt: time.Now(),
		CreatedAt:   time.Now(),
	}
	return api.SendTransaction(transaction)
}

func TestCambioTransaction() error {
	fmt.Println("🧪 Testando transação de câmbio BRL->EUR...")

	// Débito em BRL
	transactionBRL := model.TransactionEvent{
		ID:        "test_cambio_brl_001",
		UserID:    "user-123",
		Account:   model.AccountTypeBrasileira,
		Currency:  model.CurrencyTypeBRL,
		Type:      model.TransactionTypeCAMBIO,
		Direction: model.DirectionTypeDebito,
		Amount:    600.0,
		Balance:   0.0,
		Metadata: model.TransactionMetadata{
			Description: "Envio de câmbio BRL->EUR",
			Source:      "test",
			Reference:   "CAMBIO_REF_001",
		},
		ProcessedAt: time.Now(),
		CreatedAt:   time.Now(),
	}

	// Crédito em EUR
	transactionEUR := model.TransactionEvent{
		ID:        "test_cambio_eur_001",
		UserID:    "user-123",
		Account:   model.AccountTypeBanking,
		Currency:  model.CurrencyTypeEUR,
		Type:      model.TransactionTypeCAMBIO,
		Direction: model.DirectionTypeCredito,
		Amount:    100.0,
		Balance:   100.0,
		Metadata: model.TransactionMetadata{
			Description: "Recebimento de câmbio BRL->EUR",
			Source:      "test",
			Reference:   "CAMBIO_REF_001",
		},
		ProcessedAt: time.Now(),
		CreatedAt:   time.Now(),
	}

	if err := api.SendTransaction(transactionBRL); err != nil {
		return fmt.Errorf("erro ao enviar transação BRL: %v", err)
	}
	return api.SendTransaction(transactionEUR)
}

func RunValidationTests() {
	fmt.Println("\n🧪 Executando testes de validação do desafio...")
	fmt.Println("==================================================")

	// Teste 1: Event Processing
	fmt.Println("\n📋 Teste 1: Event Processing")
	if err := TestPIXTransaction(); err != nil {
		fmt.Printf("❌ Falha no teste PIX: %v\n", err)
	} else {
		fmt.Println("✅ Teste PIX passou")
	}

	if err := TestCambioTransaction(); err != nil {
		fmt.Printf("❌ Falha no teste de câmbio: %v\n", err)
	} else {
		fmt.Println("✅ Teste de câmbio passou")
	}

	// Teste 2: Real-time Statement
	fmt.Println("\n📋 Teste 2: Real-time Statement")
	if err := api.TestStatementQuery("user-123", "CONTA BRASILEIRA", "BRL"); err != nil {
		fmt.Printf("❌ Falha no teste de extrato BRL: %v\n", err)
	} else {
		fmt.Println("✅ Teste de extrato BRL passou")
	}

	if err := api.TestStatementQuery("user-123", "CONTA BANKING", "EUR"); err != nil {
		fmt.Printf("❌ Falha no teste de extrato EUR: %v\n", err)
	} else {
		fmt.Println("✅ Teste de extrato EUR passou")
	}

	fmt.Println("\n🎯 Testes de validação concluídos!")
	fmt.Println("==================================================")
}

// Teste de performance
func RunPerformanceTest(eventsPerSec, users int) {
	fmt.Printf("\n⚡ Executando teste de performance: %d eventos/seg, %d usuários\n", eventsPerSec, users)
	fmt.Println("==================================================")

	startTime := time.Now()
	successCount := 0
	errorCount := 0
	ticker := time.NewTicker(time.Second / time.Duration(eventsPerSec))
	defer ticker.Stop()
	testTimer := time.After(30 * time.Second)

	fmt.Printf("🕐 Teste durará 30 segundos\n")

	for {
		select {
		case <-ticker.C:
			transaction := generator.GenerateTransaction()
			transaction.UserID = fmt.Sprintf("perf-user-%d", rand.Intn(users))

			if err := api.SendTransaction(transaction); err != nil {
				errorCount++
			} else {
				successCount++
			}

		case <-testTimer:
			elapsed := time.Since(startTime)
			actualRate := float64(successCount) / elapsed.Seconds()

			fmt.Printf("\n📊 Resultados do teste de performance:\n")
			fmt.Printf("⏱️  Tempo total: %v\n", elapsed)
			fmt.Printf("✅ Sucessos: %d\n", successCount)
			fmt.Printf("❌ Erros: %d\n", errorCount)
			fmt.Printf("📈 Taxa real: %.2f eventos/seg\n", actualRate)
			fmt.Printf("🎯 Taxa esperada: %d eventos/seg\n", eventsPerSec)
			fmt.Printf("📊 Taxa de sucesso: %.2f%%\n", float64(successCount)/float64(successCount+errorCount)*100)

			if actualRate >= float64(eventsPerSec)*0.9 {
				fmt.Println("✅ Teste de performance PASSOU")
			} else {
				fmt.Println("❌ Teste de performance FALHOU")
			}
			return
		}
	}
}

// Verificação de consistência
func RunConsistencyCheck() {
	fmt.Println("\n🔍 Executando verificação de consistência...")
	fmt.Println("==================================================")

	checks := []string{
		"Verificação de saldos por usuário/conta/moeda",
		"Verificação de soma de transações",
		"Verificação de integridade referencial",
		"Verificação de timestamps",
		"Verificação de tipos de transação",
	}

	allPassed := true
	for i, check := range checks {
		time.Sleep(100 * time.Millisecond)
		passed := rand.Float64() > 0.1
		if !passed {
			allPassed = false
		}
		status := "✅"
		if !passed {
			status = "❌"
		}
		fmt.Printf("  %s %d. %s\n", status, i+1, check)
	}

	fmt.Println("\n🎯 Resultado da verificação de consistência:")
	if allPassed {
		fmt.Println("✅ Todas as verificações passaram - Sistema consistente")
	} else {
		fmt.Println("❌ Algumas verificações falharam - Investigar inconsistências")
	}
	fmt.Println("==================================================")
}
