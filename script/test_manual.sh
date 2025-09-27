#!/bin/bash

# Script para testes manuais com curl
# Baseado nos casos de teste do desafio

echo "🧪 Executando testes manuais com curl..."
echo "=================================================="

# Verificar se o servidor está rodando
echo "🔍 Verificando servidor..."
if ! curl -s http://localhost:8080/health > /dev/null 2>&1; then
    echo "❌ Servidor não está rodando. Inicie o servidor primeiro."
    exit 1
fi
echo "✅ Servidor está rodando"

echo ""
echo "📋 Teste 1: Event Processing"
echo "----------------------------"

# Teste PIX
echo "🧪 Enviando transação PIX..."
curl -X POST http://localhost:8080/events \
  -H "Content-Type: application/json" \
  -d '{
    "id": "test_pix_manual_001",
    "user_id": "user-123",
    "account": "CONTA BRASILEIRA",
    "currency": "BRL",
    "type": "PIX",
    "direction": "CREDITO",
    "amount": 600.0,
    "balance": 600.0,
    "metadata": {
      "description": "Transferência recebida",
      "source": "manual-test",
      "reference": "PIX_MANUAL_001"
    },
    "processed_at": "'$(date -u +%Y-%m-%dT%H:%M:%SZ)'",
    "created_at": "'$(date -u +%Y-%m-%dT%H:%M:%SZ)'"
  }'

echo ""
echo "🧪 Enviando transação de câmbio BRL->EUR (débito)..."
curl -X POST http://localhost:8080/events \
  -H "Content-Type: application/json" \
  -d '{
    "id": "test_cambio_brl_manual_001",
    "user_id": "user-123",
    "account": "CONTA BRASILEIRA",
    "currency": "BRL",
    "type": "CAMBIO",
    "direction": "DEBITO",
    "amount": 600.0,
    "balance": 0.0,
    "metadata": {
      "description": "Envio de câmbio BRL->EUR",
      "source": "manual-test",
      "reference": "CAMBIO_MANUAL_001"
    },
    "processed_at": "'$(date -u +%Y-%m-%dT%H:%M:%SZ)'",
    "created_at": "'$(date -u +%Y-%m-%dT%H:%M:%SZ)'"
  }'

echo ""
echo "🧪 Enviando transação de câmbio BRL->EUR (crédito)..."
curl -X POST http://localhost:8080/events \
  -H "Content-Type: application/json" \
  -d '{
    "id": "test_cambio_eur_manual_001",
    "user_id": "user-123",
    "account": "CONTA BANKING",
    "currency": "EUR",
    "type": "CAMBIO",
    "direction": "CREDITO",
    "amount": 100.0,
    "balance": 100.0,
    "metadata": {
      "description": "Recebimento de câmbio BRL->EUR",
      "source": "manual-test",
      "reference": "CAMBIO_MANUAL_001"
    },
    "processed_at": "'$(date -u +%Y-%m-%dT%H:%M:%SZ)'",
    "created_at": "'$(date -u +%Y-%m-%dT%H:%M:%SZ)'"
  }'

echo ""
echo "📋 Teste 2: Real-time Statement"
echo "-------------------------------"

echo "🧪 Consultando extrato BRL..."
curl -X GET "http://localhost:8080/statement/user-123/CONTA%20BRASILEIRA/BRL/30d"

echo ""
echo "🧪 Consultando extrato EUR..."
curl -X GET "http://localhost:8080/statement/user-123/CONTA%20BANKING/EUR/30d"

echo ""
echo "📋 Teste 3: Consultas adicionais"
echo "--------------------------------"

echo "🧪 Consultando histórico de transações..."
curl -X GET "http://localhost:8080/transactions/user-123"

echo ""
echo "🧪 Consultando tipos de eventos..."
curl -X GET "http://localhost:8080/events/types"

echo ""
echo "🧪 Health check..."
curl -X GET "http://localhost:8080/health"

echo ""
echo "🎯 Testes manuais concluídos!"
echo "=================================================="
