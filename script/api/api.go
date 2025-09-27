package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"statement-validator/model"
)

func SendTransaction(transaction model.TransactionEvent) error {
	jsonData, err := json.Marshal(transaction)
	if err != nil {
		return err
	}

	resp, err := http.Post("http://localhost:8080/events", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status code: %d", resp.StatusCode)
	}

	return nil
}

func TestStatementQuery(userID, accountType, currencyType string) error {
	fmt.Printf("🧪 Testando consulta de extrato para user=%s, account=%s, currency=%s...\n",
		userID, accountType, currencyType)

	url := fmt.Sprintf("http://localhost:8080/statement/%s/%s/%s/30d", userID, accountType, currencyType)
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("erro na consulta de extrato: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status code inesperado: %d", resp.StatusCode)
	}

	fmt.Printf("✅ Consulta de extrato realizada com sucesso\n")
	return nil
}
