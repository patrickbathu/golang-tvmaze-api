# 📋 Guia de Migração v2.0 → v3.0

Este documento explica as mudanças entre a versão antiga (v2.0) e a nova arquitetura profissional (v3.0).

## 🔄 Principais Mudanças

### Estrutura de Arquivos

**ANTES (v2.0)**
```
.
├── tvmaze-api.go           # Tudo em um único arquivo
├── api.go                  # API GitHub em outro arquivo
└── go.mod
```

**DEPOIS (v3.0)**
```
.
├── cmd/api/main.go         # Entry point
├── internal/
│   ├── models/             # Modelos separados
│   ├── clients/            # Clientes HTTP
│   ├── services/           # Lógica de negócio
│   ├── handlers/           # Handlers HTTP
│   ├── middleware/         # Middlewares
│   └── router/             # Roteamento
├── pkg/utils/              # Utilitários
└── go.mod
```

## 🚀 Como Executar

**ANTES:**
```bash
go run tvmaze-api.go
```

**DEPOIS:**
```bash
go run cmd/api/main.go
# ou
make run
```

## 📝 Como Compilar

**ANTES:**
```bash
go build -o tvmaze-server tvmaze-api.go
```

**DEPOIS:**
```bash
go build -o bin/api-server cmd/api/main.go
# ou
make build
```

## 🧪 Como Testar

**ANTES:**
```bash
go test -v tvmaze-api.go tvmaze-api_test.go
```

**DEPOIS:**
```bash
go test ./internal/...
# ou
make test
```

## 🏗️ Mudanças na Arquitetura

### 1. Models (Antes: struct no main)
**ANTES:**
```go
// tvmaze-api.go
type Show struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}
```

**DEPOIS:**
```go
// internal/models/tvmaze.go
package models

type Show struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}
```

### 2. Clientes HTTP (Antes: funções soltas)
**ANTES:**
```go
// tvmaze-api.go
func getTodaySchedule(country string) ([]Schedule, error) {
    url := fmt.Sprintf("https://api.tvmaze.com/schedule?country=%s", country)
    resp, err := http.Get(url)
    // ...
}
```

**DEPOIS:**
```go
// internal/clients/tvmaze.go
type TVMazeClient struct {
    httpClient *http.Client
    baseURL    string
}

func (c *TVMazeClient) GetSchedule(country, date string) ([]models.Schedule, error) {
    // ...
}
```

### 3. Serviços (Antes: lógica nos handlers)
**ANTES:**
```go
func scheduleHandler(w http.ResponseWriter, r *http.Request) {
    country := r.URL.Query().Get("country")
    schedule, err := getTodaySchedule(country)
    // lógica de negócio aqui
    json.NewEncoder(w).Encode(schedule)
}
```

**DEPOIS:**
```go
// internal/services/tvmaze.go
func (s *TVMazeService) GetTodaySchedule(country string) ([]models.Schedule, error) {
    today := time.Now().Format("2006-01-02")
    return s.client.GetSchedule(country, today)
}

// internal/handlers/tvmaze.go
func (h *TVMazeHandler) Schedule(w http.ResponseWriter, r *http.Request) {
    country := r.URL.Query().Get("country")
    schedule, err := h.service.GetTodaySchedule(country)
    json.NewEncoder(w).Encode(schedule)
}
```

### 4. Roteamento (Antes: no main)
**ANTES:**
```go
func main() {
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/schedule", scheduleHandler)
    // ...
    http.ListenAndServe(":8080", nil)
}
```

**DEPOIS:**
```go
// internal/router/router.go
func Setup(tvmazeHandler *handlers.TVMazeHandler) *http.ServeMux {
    mux := http.NewServeMux()
    mux.HandleFunc("/", middleware.Logging(tvmazeHandler.Home))
    mux.HandleFunc("/schedule", middleware.Logging(tvmazeHandler.Schedule))
    return mux
}

// cmd/api/main.go
func main() {
    mux := router.Setup(tvmazeHandler, githubHandler)
    server := &http.Server{Addr: ":8080", Handler: mux}
    server.ListenAndServe()
}
```

## 📊 Comparação de Benefícios

| Aspecto | v2.0 | v3.0 |
|---------|------|------|
| **Organização** | Tudo em 1 arquivo | Separado por responsabilidade |
| **Testabilidade** | Difícil testar isoladamente | Fácil criar mocks e testes unitários |
| **Manutenção** | Arquivo grande, difícil navegar | Arquivos pequenos e focados |
| **Escalabilidade** | Difícil adicionar features | Estrutura preparada para crescer |
| **Padrões** | Básico | Segue padrões de mercado |
| **Reusabilidade** | Código acoplado | Componentes reutilizáveis |

## ✅ Checklist de Migração

Se você tem código baseado na v2.0, siga estes passos:

- [ ] Criar estrutura de pastas (cmd, internal, pkg)
- [ ] Mover structs para `internal/models/`
- [ ] Criar clientes em `internal/clients/`
- [ ] Extrair lógica de negócio para `internal/services/`
- [ ] Criar handlers em `internal/handlers/`
- [ ] Implementar middlewares em `internal/middleware/`
- [ ] Configurar rotas em `internal/router/`
- [ ] Criar novo `main.go` em `cmd/api/`
- [ ] Atualizar Dockerfile e Makefile
- [ ] Criar testes para cada camada
- [ ] Atualizar documentação

## 🎯 Próximos Passos

Com a nova estrutura, você pode facilmente:

1. **Adicionar novos endpoints**: Criar handler → service → client
2. **Adicionar cache**: Implementar no service layer
3. **Adicionar autenticação**: Criar middleware
4. **Adicionar banco de dados**: Criar repository layer
5. **Adicionar métricas**: Middleware de metrics
6. **Migrar para gRPC**: Manter services, trocar handlers

## 📚 Documentação Adicional

- [ESTRUTURA.md](ESTRUTURA.md) - Detalhes da arquitetura
- [README.md](README.md) - Guia principal
- [ARQUITETURA.md](ARQUITETURA.md) - Diagramas

## 🆘 Suporte

Se tiver dúvidas sobre a migração:
1. Leia a documentação em `ESTRUTURA.md`
2. Veja exemplos em `internal/`
3. Execute os testes com `make test`
4. Consulte os commits do GitHub

---

**💡 Dica:** A estrutura antiga ainda funciona e está nos arquivos `tvmaze-api.go` e `api.go`. Use-os como referência durante a migração!
