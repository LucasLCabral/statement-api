# 🏦 Sistema de Extrato Multi Conta e Moeda - Makefile
# ========================================================

.PHONY: help build run test test-performance test-consistency test-manual clean install deps check-api

# Configurações
BINARY_NAME=statement-validator
SCRIPT_DIR=script
API_URL=http://localhost:8080

# Cores para output
GREEN=\033[0;32m
YELLOW=\033[1;33m
RED=\033[0;31m
NC=\033[0m # No Color

# Help - Comando padrão
help: ## 📋 Mostra esta ajuda
	@echo "$(GREEN)🏦 Sistema de Extrato Multi Conta e Moeda$(NC)"
	@echo "$(YELLOW)=======================================================$(NC)"
	@echo ""
	@echo "$(GREEN)Comandos disponíveis:$(NC)"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  $(YELLOW)%-20s$(NC) %s\n", $$1, $$2}'
	@echo ""
	@echo "$(GREEN)Exemplos:$(NC)"
	@echo "  make build          # Compila o projeto"
	@echo "  make test           # Executa testes de validação"
	@echo "  make run            # Executa validação completa"
	@echo "  make check-api      # Verifica se a API está rodando"

# Build
build: ## 🔨 Compila o projeto
	@echo "$(GREEN)🔨 Compilando o projeto...$(NC)"
	@cd $(SCRIPT_DIR) && go build -o $(BINARY_NAME) main.go
	@echo "$(GREEN)✅ Compilação concluída!$(NC)"

# Dependências
deps: ## 📦 Instala/atualiza dependências
	@echo "$(GREEN)📦 Atualizando dependências...$(NC)"
	@cd $(SCRIPT_DIR) && go mod tidy
	@cd $(SCRIPT_DIR) && go mod download
	@echo "$(GREEN)✅ Dependências atualizadas!$(NC)"

# Verificar API
check-api: ## 🔍 Verifica se a API está rodando
	@echo "$(GREEN)🔍 Verificando se a API está rodando em $(API_URL)...$(NC)"
	@if curl -s $(API_URL)/health > /dev/null 2>&1; then \
		echo "$(GREEN)✅ API está rodando!$(NC)"; \
	else \
		echo "$(RED)❌ API não está rodando em $(API_URL)$(NC)"; \
		echo "$(YELLOW)💡 Inicie a API antes de executar os testes$(NC)"; \
		exit 1; \
	fi

# Execução principal
run: check-api ## 🚀 Executa validação completa (inserção em massa + testes)
	@echo "$(GREEN)🚀 Executando validação completa...$(NC)"
	@cd $(SCRIPT_DIR) && go run main.go

# Testes
test: check-api ## 🧪 Executa apenas testes de validação
	@echo "$(GREEN)🧪 Executando testes de validação...$(NC)"
	@cd $(SCRIPT_DIR) && go run main.go --test

test-performance: check-api ## ⚡ Executa teste de performance
	@echo "$(GREEN)⚡ Executando teste de performance...$(NC)"
	@cd $(SCRIPT_DIR) && go run main.go --performance

test-consistency: check-api ## 🔍 Executa verificação de consistência
	@echo "$(GREEN)🔍 Executando verificação de consistência...$(NC)"
	@cd $(SCRIPT_DIR) && go run main.go --consistency

test-manual: check-api ## 🧪 Executa testes manuais com curl
	@echo "$(GREEN)🧪 Executando testes manuais...$(NC)"
	@cd $(SCRIPT_DIR) && chmod +x test_manual.sh && ./test_manual.sh

test-all: check-api ## 🎯 Executa todos os testes (script automático)
	@echo "$(GREEN)🎯 Executando todos os testes...$(NC)"
	@cd $(SCRIPT_DIR) && chmod +x run_tests.sh && ./run_tests.sh

# Desenvolvimento
dev: ## 🔧 Modo desenvolvimento (watch mode)
	@echo "$(GREEN)🔧 Modo desenvolvimento ativado...$(NC)"
	@echo "$(YELLOW)💡 Use 'make test' para executar testes rapidamente$(NC)"
	@echo "$(YELLOW)💡 Use 'make run' para validação completa$(NC)"

# Limpeza
clean: ## 🧹 Remove arquivos compilados e temporários
	@echo "$(GREEN)🧹 Limpando arquivos temporários...$(NC)"
	@cd $(SCRIPT_DIR) && rm -f $(BINARY_NAME)
	@echo "$(GREEN)✅ Limpeza concluída!$(NC)"

# Instalação
install: build ## 📦 Instala o binário no sistema
	@echo "$(GREEN)📦 Instalando $(BINARY_NAME)...$(NC)"
	@cp $(SCRIPT_DIR)/$(BINARY_NAME) /usr/local/bin/
	@echo "$(GREEN)✅ Instalação concluída!$(NC)"
	@echo "$(YELLOW)💡 Agora você pode usar '$(BINARY_NAME)' de qualquer lugar$(NC)"

# Status
status: ## 📊 Mostra status do projeto
	@echo "$(GREEN)📊 Status do Projeto$(NC)"
	@echo "$(YELLOW)===================$(NC)"
	@echo ""
	@echo "$(GREEN)Estrutura:$(NC)"
	@ls -la $(SCRIPT_DIR)/
	@echo ""
	@echo "$(GREEN)API Status:$(NC)"
	@make check-api || true
	@echo ""
	@echo "$(GREEN)Go Version:$(NC)"
	@go version
	@echo ""
	@echo "$(GREEN)Go Modules:$(NC)"
	@cd $(SCRIPT_DIR) && go list -m all

# Docker (opcional)
docker-build: ## 🐳 Constrói imagem Docker
	@echo "$(GREEN)🐳 Construindo imagem Docker...$(NC)"
	@docker build -t statement-validator .
	@echo "$(GREEN)✅ Imagem Docker construída!$(NC)"

docker-run: ## 🐳 Executa container Docker
	@echo "$(GREEN)🐳 Executando container...$(NC)"
	@docker run --rm -it statement-validator

# Debug
debug: check-api ## 🐛 Modo debug (com logs detalhados)
	@echo "$(GREEN)🐛 Modo debug ativado...$(NC)"
	@cd $(SCRIPT_DIR) && go run main.go --test

# Validação rápida
quick: check-api ## ⚡ Validação rápida (apenas testes essenciais)
	@echo "$(GREEN)⚡ Executando validação rápida...$(NC)"
	@cd $(SCRIPT_DIR) && go run main.go --test

# Comando padrão
.DEFAULT_GOAL := help
