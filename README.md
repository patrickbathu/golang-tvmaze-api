# 📺 TVMaze API - Go Microservice

API REST em Go que consulta a programação de TV em tempo real usando a [TVMaze API](https://www.tvmaze.com/api).

## 🚀 Como Executar

### Método 1: Executar diretamente
```bash
go run tvmaze-api.go
```

### Método 2: Compilar e executar
```bash
go build -o tvmaze-server tvmaze-api.go
./tvmaze-server
```

### Método 3: Com Docker
```bash
docker-compose up
```

## 🔌 Endpoints

### 1. Informações da API
```bash
curl http://localhost:8080/
```

### 2. Programação de hoje (EUA)
```bash
curl "http://localhost:8080/schedule?country=US"
```

### 3. Programação do Brasil
```bash
curl "http://localhost:8080/schedule?country=BR"
```

### 4. Buscar show
```bash
curl "http://localhost:8080/search?q=friends"
```

### 5. Detalhes de um show
```bash
curl "http://localhost:8080/show?id=431"
```

## 🧪 Testes

```bash
# Executar testes
go test -v tvmaze-api.go tvmaze-api_test.go

# Com coverage
go test -cover tvmaze-api.go tvmaze-api_test.go
```

## 📦 Estrutura do Projeto

```
.
├── tvmaze-api.go           # 🎯 API principal
├── tvmaze-api_test.go      # 🧪 Testes unitários
├── api.go                  # 📦 API GitHub (outro exemplo)
├── api_test.go             # 🧪 Testes API GitHub
├── go.mod                  # 📦 Dependências
├── Dockerfile              # 🐳 Container
├── docker-compose.yml      # 🎼 Orquestração
└── examples/
    └── primeiroGoLang.go   # 👋 Hello World
```

## �� Deploy

Ver documentação completa em:
- `DEPLOY_TVMAZE.md` - Guia completo de deploy
- `TVMAZE_RESUMO.md` - Resumo do projeto
- `ARQUITETURA.md` - Arquitetura da aplicação

### Deploy Rápido (Railway)
```bash
# Instalar CLI
brew install railway

# Deploy
railway login
railway init
railway up
```

## 📚 Documentação Completa

- **DEPLOY_TVMAZE.md** - Guia de deploy em várias plataformas
- **README_TVMAZE.md** - Documentação detalhada da API
- **TVMAZE_RESUMO.md** - Resumo completo do projeto
- **ARQUITETURA.md** - Diagrama e explicação da arquitetura

## 🎯 Conceitos Go Implementados

- ✅ Structs e JSON tags
- ✅ HTTP Server e Client
- ✅ Error handling
- ✅ Context e graceful shutdown
- ✅ Middleware
- ✅ Testing
- ✅ Goroutines
- ✅ Channels

## 📝 Licença

MIT

---

Criado como projeto de aprendizado de Go Lang 🐹
