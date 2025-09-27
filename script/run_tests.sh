#!/bin/bash

# Script para executar os testes de validação do desafio
# Sistema de Extrato Multi Conta e Moeda

echo "🚀 Iniciando validação do desafio..."
echo "=================================================="

# Verificar se o servidor da API está rodando
echo "🔍 Verificando se o servidor da API está rodando..."
if curl -s http://localhost:8080/health > /dev/null 2>&1; then
    echo "✅ Servidor da API está rodando"
else
    echo "❌ Servidor da API não está rodando em http://localhost:8080"
    echo "   Por favor, inicie o servidor antes de executar os testes"
    exit 1
fi

echo ""
echo "🧪 Executando testes de validação..."
echo "=================================================="

# Teste 1: Apenas testes de validação
echo "📋 Teste 1: Executando testes de validação específicos"
go run main.go --test

echo ""
echo "⚡ Teste 2: Executando teste de performance"
go run main.go --performance

echo ""
echo "🔍 Teste 3: Executando verificação de consistência"
go run main.go --consistency

echo ""
echo "🎯 Teste 4: Executando validação completa (inserção em massa + testes)"
go run main.go

echo ""
echo "🎉 Todos os testes foram executados!"
echo "=================================================="
echo "📊 Resumo dos testes executados:"
echo "  ✅ Testes de validação específicos"
echo "  ✅ Teste de performance"
echo "  ✅ Verificação de consistência"
echo "  ✅ Validação completa com inserção em massa"
echo ""
echo "💡 Para executar testes individuais, use:"
echo "  go run main.go --test        # Apenas validação"
echo "  go run main.go --performance # Apenas performance"
echo "  go run main.go --consistency # Apenas consistência"
echo "  go run main.go               # Validação completa"
