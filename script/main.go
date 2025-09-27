package main

import (
	"fmt"
	"os"
	"time"

	"statement-validator/api"
	"statement-validator/generator"
	"statement-validator/tests"
)

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "--test":
			tests.RunValidationTests()
			return
		case "--performance":
			eventsPerSec := 500
			users := 100
			if len(os.Args) > 2 {
				eventsPerSec = 500
			}
			if len(os.Args) > 3 {
				users = 100
			}
			tests.RunPerformanceTest(eventsPerSec, users)
			return
		case "--consistency":
			tests.RunConsistencyCheck()
			return
		case "--help":
			fmt.Println("Uso: go run main.go [opção]")
			fmt.Println("Opções:")
			fmt.Println("  --test        Executa apenas os testes de validação")
			fmt.Println("  --performance Executa teste de performance")
			fmt.Println("  --consistency Executa verificação de consistência")
			fmt.Println("  --help        Mostra esta ajuda")
			fmt.Println("  (sem opção)   Executa inserção em massa com testes")
			return
		}
	}

	fmt.Println("🚀 Iniciando inserção de transações para validação do desafio...")
	fmt.Println("📊 Cenários disponíveis:")
	for name, scenario := range generator.Scenarios {
		fmt.Printf("  - %s: Account=%s, Currency=%s, Types=%v\n",
			name, scenario.Account, scenario.Currency, scenario.Types)
	}
	fmt.Println()

	// Executar testes de validação antes da inserção em massa
	fmt.Println("🧪 Executando testes de validação...")
	tests.RunValidationTests()

	// Executar inserção em massa
	runBulkInsert()

	// Executar verificação de consistência no final
	fmt.Println("\n🔍 Executando verificação de consistência final...")
	tests.RunConsistencyCheck()

	fmt.Println("\n🎉 Validação do desafio concluída!")
	fmt.Println("==================================================")
}

func runBulkInsert() {
	// Configurações
	totalTransactions := 1000
	batchSize := 10
	delayBetweenBatches := 100 * time.Millisecond

	fmt.Printf("📈 Gerando %d transações em lotes de %d...\n", totalTransactions, batchSize)
	fmt.Printf("⏱️  Delay entre lotes: %v\n\n", delayBetweenBatches)

	successCount := 0
	errorCount := 0

	// Contadores por tipo
	typeCounts := make(map[string]int)
	directionCounts := make(map[string]int)
	accountCounts := make(map[string]int)
	currencyCounts := make(map[string]int)

	for i := 0; i < totalTransactions; i += batchSize {
		batchEnd := i + batchSize
		if batchEnd > totalTransactions {
			batchEnd = totalTransactions
		}

		fmt.Printf("🔄 Processando lote %d-%d...\n", i+1, batchEnd)

		for j := i; j < batchEnd; j++ {
			transaction := generator.GenerateTransaction()

			if err := api.SendTransaction(transaction); err != nil {
				fmt.Printf("❌ Erro na transação %d: %v\n", j+1, err)
				errorCount++
			} else {
				successCount++
				typeCounts[string(transaction.Type)]++
				directionCounts[string(transaction.Direction)]++
				accountCounts[string(transaction.Account)]++
				currencyCounts[string(transaction.Currency)]++

				if (j+1)%50 == 0 {
					fmt.Printf("✅ %d transações enviadas com sucesso\n", j+1)
				}
			}
		}

		if batchEnd < totalTransactions {
			time.Sleep(delayBetweenBatches)
		}
	}

	// Relatório de resultados
	fmt.Printf("\n🎯 Resumo da execução:\n")
	fmt.Printf("✅ Sucessos: %d\n", successCount)
	fmt.Printf("❌ Erros: %d\n", errorCount)
	fmt.Printf("📊 Taxa de sucesso: %.2f%%\n", float64(successCount)/float64(totalTransactions)*100)

	fmt.Printf("\n📈 Transações por tipo:\n")
	for transactionType, count := range typeCounts {
		fmt.Printf("  %s: %d transações\n", transactionType, count)
	}

	fmt.Printf("\n🔄 Transações por direção:\n")
	for direction, count := range directionCounts {
		fmt.Printf("  %s: %d transações\n", direction, count)
	}

	fmt.Printf("\n🏦 Transações por conta:\n")
	for account, count := range accountCounts {
		fmt.Printf("  %s: %d transações\n", account, count)
	}

	fmt.Printf("\n💱 Transações por moeda:\n")
	for currency, count := range currencyCounts {
		fmt.Printf("  %s: %d transações\n", currency, count)
	}
}
