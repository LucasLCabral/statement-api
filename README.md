# 🏦 Sistema de Extrato Multi Conta e Moeda

## 🎯 Objetivo

Desenvolver um sistema que gera extratos atualizados instantaneamente para os usuários, permitindo acompanhar movimentações em tempo real, para contas/moedas diferentes.

## 🚀 Início Rápido

```bash
# Verificar comandos disponíveis
make help

# Executar validação completa
make run

# Apenas testes de validação
make test

# Verificar se a API está rodando
make check-api
```

## 📋 Contexto

Sistemas bancários modernos precisam:

- Processar diferentes tipos de transações em diferentes contas/moedas:
  - **Brasil - BRL**: PIX, TED, Câmbio
  - **Investimento - USD**: Câmbio, Compra e venda de ações
  - **Banking - USD**: Câmbio, Cartão, WIRE In/Out
  - **Banking - EUR**: Câmbio, Cartão
- Fornecer extratos em tempo real
- Manter histórico completo de movimentações
- Suportar milhares de usuários simultâneos

## 🔌 Endpoints da API

### Core Endpoints:
- `POST /events` - Incluir novos eventos/transações
- `GET /statement/{userId}/{AccountType}/{CurrencyType}/{period}` - Extrato do usuário
- `GET /transactions/{userId}` - Histórico de transações
- `GET /events/types` - Tipos de eventos suportados
- `GET /health` - Health check

## ✅ Funcionalidades a Implementar

### Core Features:
- **Event Processing**
  - Geração de eventos de diferentes tipos
  - Validação e sanitização de dados
  - Channel pipeline para alta performance

- **Statement Generation**
  - Agregação de transações por usuário/conta e moeda
  - Cálculo de saldos em tempo real
  - Histórico completo de movimentações

- **Caching & Performance**
  - Agregações em memória
  - Otimização de consultas

### Advanced Features:
- **Consistency Auditing**
  - Verificação de integridade dos saldos
  - Detecção de inconsistências
  - Auto-correção quando possível

- **Event Replay**
  - Debugging e troubleshooting
  - Backup e recovery

## 💡 Dicas de Implementação

- Use buffered channels para alta frequência de eventos
- Implemente fan-out patterns para distribuição
- Monitore memory leaks em cache
- Teste cenários de alta concorrência
- Mantenha consistência entre cache e banco

---

# 🧪 Script de Validação

Este repositório contém um script completo para validar a implementação do sistema de extrato multi conta e moeda, organizado em uma estrutura modular e com Makefile para facilitar o uso.

## 📁 Estrutura do Projeto

```
statement/
├── Makefile                    # Comandos principais do projeto
├── README.md                   # Documentação completa
│
└── script/                     # Script de validação
    ├── main.go                 # Ponto de entrada principal
    ├── go.mod                  # Dependências do Go
    ├── run_tests.sh           # Script automático de testes
    ├── test_manual.sh         # Testes manuais com curl
    │
    ├── api/                   # Cliente da API
    │   └── client.go         # Funções de comunicação HTTP
    │
    ├── config/               # Configurações
    │   └── config.json       # Cenários e configurações
    │
    ├── generator/            # Geração de dados
    │   └── transaction.go    # Gerador de transações
    │
    ├── model/               # Modelos de dados
    │   └── types.go         # Definições de tipos
    │
    └── tests/               # Testes de validação
        └── validation.go    # Testes específicos
```

## 🚀 Como Usar

### 🎯 Com Makefile (Recomendado)

```bash
# Ver todos os comandos disponíveis
make help

# Validação completa (inserção em massa + testes)
make run

# Apenas testes de validação
make test

# Teste de performance
make test-performance

# Verificação de consistência
make test-consistency

# Testes manuais com curl
make test-manualclear

# Executar todos os testes
make test-all

# Verificar se a API está rodando
make check-api

# Compilar o projeto
make build

# Limpar arquivos temporários
make clean
```

### 🔧 Comandos Diretos

#### Validação Completa (Recomendado)
```bash
cd script
go run main.go
```
Executa inserção em massa (1000 transações) + testes de validação + verificação de consistência.

#### Apenas Testes de Validação
```bash
cd script
go run main.go --test
```
Executa apenas os testes específicos mencionados no desafio.

#### Teste de Performance
```bash
cd script
go run main.go --performance
```
Executa teste de performance com 500 eventos/seg e 100 usuários.

#### Verificação de Consistência
```bash
cd script
go run main.go --consistency
```
Executa apenas a verificação de consistência dos dados.

#### Ajuda
```bash
cd script
go run main.go --help
```

### 📜 Scripts Automáticos
```bash
# Script completo de testes
cd script
chmod +x run_tests.sh
./run_tests.sh

# Testes manuais com curl
cd script
chmod +x test_manual.sh
./test_manual.sh
```

## 🧪 Testes Implementados

### ✅ Teste 1: Event Processing
- **PIX**: Transação de R$ 600,00
- **Câmbio**: Conversão BRL->EUR (R$ 600,00 -> €100,00)

### ✅ Teste 2: Real-time Statement
- Consultas de extrato por usuário/conta/moeda
- Validação de endpoints de statement

### ✅ Teste 3: Performance Test
- Throughput esperado: > 500 eventos/seg
- Latência esperada: < 50ms
- Usuários simultâneos: 100

### ✅ Teste 4: Consistency Check
- Verificação de saldos calculados vs cache
- Soma de transações vs saldo final
- Integridade referencial

## 🎯 Cenários de Teste Suportados

### Conta Brasileira (BRL)
- **PIX**: Transferências instantâneas
- **TED**: Transferências eletrônicas
- **Câmbio**: Conversões de moeda

### Conta Investimento (USD)
- **Câmbio**: Conversões de moeda
- **Ações**: Compra e venda de ações

### Conta Banking USD
- **Câmbio**: Conversões de moeda
- **Cartão**: Transações com cartão
- **WIRE**: Transferências internacionais

### Conta Banking EUR
- **Câmbio**: Conversões de moeda
- **Cartão**: Transações com cartão

## 📊 Relatórios Gerados

O script gera relatórios detalhados com:
- Contadores por tipo de transação
- Contadores por direção (crédito/débito)
- Contadores por conta e moeda
- Taxa de sucesso/erro
- Resultados de performance
- Verificação de consistência

## 🔧 Pré-requisitos

1. **Go 1.19+** instalado
2. **Servidor da API** rodando em `http://localhost:8080`
3. **Endpoints da API** implementados conforme especificação

## 📋 Exemplos de Uso com cURL

### Teste PIX
```bash
curl -X POST http://localhost:8080/events \
  -H "Content-Type: application/json" \
  -d '{
    "id": "test_pix_001",
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
      "reference": "PIX_REF_001"
    },
    "processed_at": "2024-01-01T10:00:00Z",
    "created_at": "2024-01-01T10:00:00Z"
  }'
```

### Teste de Câmbio BRL->EUR
```bash
# Débito em BRL
curl -X POST http://localhost:8080/events \
  -H "Content-Type: application/json" \
  -d '{
    "id": "test_cambio_brl_001",
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
      "reference": "CAMBIO_REF_001"
    },
    "processed_at": "2024-01-01T10:00:00Z",
    "created_at": "2024-01-01T10:00:00Z"
  }'

# Crédito em EUR
curl -X POST http://localhost:8080/events \
  -H "Content-Type: application/json" \
  -d '{
    "id": "test_cambio_eur_001",
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
      "reference": "CAMBIO_REF_001"
    },
    "processed_at": "2024-01-01T10:00:00Z",
    "created_at": "2024-01-01T10:00:00Z"
  }'
```

### Consulta de Extrato
```bash
curl -X GET "http://localhost:8080/statement/user-123/CONTA%20BRASILEIRA/BRL/30d"
```

## 🛠️ Makefile

O projeto inclui um Makefile completo com comandos úteis para desenvolvimento e teste:

### 📋 Comandos Principais
- `make help` - Mostra todos os comandos disponíveis
- `make run` - Executa validação completa
- `make test` - Executa apenas testes de validação
- `make test-performance` - Executa teste de performance
- `make test-consistency` - Executa verificação de consistência
- `make test-manual` - Executa testes manuais com curl
- `make test-all` - Executa todos os testes
- `make check-api` - Verifica se a API está rodando
- `make build` - Compila o projeto
- `make clean` - Remove arquivos temporários
- `make deps` - Atualiza dependências
- `make status` - Mostra status do projeto

### 🎯 Comandos de Desenvolvimento
- `make dev` - Modo desenvolvimento
- `make debug` - Modo debug com logs detalhados
- `make quick` - Validação rápida (apenas testes essenciais)

### 🐳 Comandos Docker (Opcional)
- `make docker-build` - Constrói imagem Docker
- `make docker-run` - Executa container Docker

## 🎉 Pronto para Usar!

O sistema está completo e pronto para validar a implementação do desafio do Sistema de Extrato Multi Conta e Moeda!

### 🚀 Formas de Execução:
1. **Makefile (Recomendado)**: `make run`
2. **Scripts**: `./run_tests.sh` ou `./test_manual.sh`
3. **Comandos diretos**: `go run main.go`

### 📁 Arquivos Principais:
- `Makefile` - Comandos principais do projeto
- `script/run_tests.sh` - Executa todos os testes automaticamente
- `script/test_manual.sh` - Testes manuais com cURL
- `script/main.go` - Execução principal do validador
