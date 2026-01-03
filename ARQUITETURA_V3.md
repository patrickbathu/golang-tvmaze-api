# 🏛️ Arquitetura da Aplicação - Versão 3.0

## 📊 Diagrama da Arquitetura em Camadas

```
┌─────────────────────────────────────────────────────────────────┐
│                         CLIENTE HTTP                             │
│                    (Browser, curl, Postman)                      │
└──────────────────────────┬──────────────────────────────────────┘
                           │
                           ▼
┌─────────────────────────────────────────────────────────────────┐
│                    ROUTER (router.go)                            │
│                  Registra todas as rotas                         │
│     GET /schedule, /search, /show, /genre, /now, /docs...       │
└──────────────────────────┬──────────────────────────────────────┘
                           │
                           ▼
┌─────────────────────────────────────────────────────────────────┐
│                 MIDDLEWARE (middleware.go)                       │
│                                                                  │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐         │
│  │   Logging    │  │     CORS     │  │    Auth      │         │
│  │   📝 Logs    │  │  🌍 Headers  │  │  (futuro)    │         │
│  └──────────────┘  └──────────────┘  └──────────────┘         │
└──────────────────────────┬──────────────────────────────────────┘
                           │
                           ▼
┌─────────────────────────────────────────────────────────────────┐
│                  HANDLERS (handlers/)                            │
│              Camada de Apresentação HTTP                         │
│                                                                  │
│  ┌─────────────────────┐      ┌─────────────────────┐          │
│  │  TVMazeHandler      │      │  GitHubHandler      │          │
│  │  • Home()           │      │  • GetUser()        │          │
│  │  • Schedule()       │      │  • Home()           │          │
│  │  • Search()         │      └─────────────────────┘          │
│  │  • ShowDetails()    │                                        │
│  │  • Genre()          │      ┌─────────────────────┐          │
│  │  • NowPlaying()     │      │   DocsHandler       │          │
│  └─────────────────────┘      │  • Docs()           │          │
│                                └─────────────────────┘          │
│                                                                  │
│  Responsabilidades:                                              │
│  ✓ Validar parâmetros da requisição                            │
│  ✓ Chamar services                                              │
│  ✓ Formatar resposta JSON                                       │
│  ✓ Definir status codes HTTP                                    │
└──────────────────────────┬──────────────────────────────────────┘
                           │
                           ▼
┌─────────────────────────────────────────────────────────────────┐
│                  SERVICES (services/)                            │
│                 Lógica de Negócio                               │
│                                                                  │
│  ┌─────────────────────┐      ┌─────────────────────┐          │
│  │  TVMazeService      │      │  GitHubService      │          │
│  │  • GetTodaySchedule │      │  • GetUser()        │          │
│  │  • SearchShows()    │      └─────────────────────┘          │
│  │  • GetShowByID()    │                                        │
│  │  • GetScheduleBy    │                                        │
│  │    Genre()          │                                        │
│  │  • GetNowPlaying()  │                                        │
│  └─────────────────────┘                                        │
│                                                                  │
│  Responsabilidades:                                              │
│  ✓ Orquestrar chamadas aos clients                             │
│  ✓ Validações de negócio                                        │
│  ✓ Filtros e transformações                                     │
│  ✓ Cálculos e processamento                                     │
└──────────────────────────┬──────────────────────────────────────┘
                           │
                           ▼
┌─────────────────────────────────────────────────────────────────┐
│                  CLIENTS (clients/)                              │
│            Comunicação com APIs Externas                         │
│                                                                  │
│  ┌─────────────────────┐      ┌─────────────────────┐          │
│  │  TVMazeClient       │      │  GitHubClient       │          │
│  │  • GetSchedule()    │      │  • GetUser()        │          │
│  │  • SearchShows()    │      │                     │          │
│  │  • GetShowByID()    │      │  HTTP Client        │          │
│  │                     │      │  Timeout: 10s       │          │
│  │  HTTP Client        │      └─────────────────────┘          │
│  │  Timeout: 15s       │                                        │
│  └─────────────────────┘                                        │
│                                                                  │
│  Responsabilidades:                                              │
│  ✓ Requisições HTTP                                             │
│  ✓ Parsing de respostas                                         │
│  ✓ Tratamento de erros de rede                                  │
│  ✓ Timeouts e retries                                           │
└──────────────────────────┬──────────────────────────────────────┘
                           │
                           ▼
┌─────────────────────────────────────────────────────────────────┐
│                     APIs EXTERNAS                                │
│                                                                  │
│  ┌─────────────────────┐      ┌─────────────────────┐          │
│  │   TVMaze API        │      │   GitHub API        │          │
│  │  api.tvmaze.com     │      │  api.github.com     │          │
│  └─────────────────────┘      └─────────────────────┘          │
└─────────────────────────────────────────────────────────────────┘


         ┌─────────────────────────────────────────┐
         │        MODELS (models/)                 │
         │     Usado por todas as camadas          │
         │                                         │
         │  • Show, Episode, Schedule              │
         │  • GitHubUser                           │
         │  • Response                             │
         │  • Network, Country, Image              │
         └─────────────────────────────────────────┘
```

## 🔄 Fluxo de uma Requisição Típica

### Exemplo: `GET /schedule?country=BR`

```
1. Cliente faz requisição HTTP
   GET http://localhost:8080/schedule?country=BR

2. Router identifica a rota
   /schedule → TVMazeHandler.Schedule

3. Middleware processa
   • Logging: registra "GET /schedule"
   • CORS: adiciona headers

4. Handler processa
   TVMazeHandler.Schedule()
   • Extrai parâmetro: country = "BR"
   • Valida entrada
   • Chama service

5. Service executa lógica de negócio
   TVMazeService.GetTodaySchedule("BR")
   • Calcula data de hoje: "2026-01-03"
   • Chama client

6. Client faz requisição à API externa
   TVMazeClient.GetSchedule("BR", "2026-01-03")
   • HTTP GET https://api.tvmaze.com/schedule?country=BR&date=2026-01-03
   • Parse JSON
   • Retorna []models.Schedule

7. Service retorna para Handler
   • Pode aplicar filtros adicionais
   • Retorna dados processados

8. Handler formata resposta
   • Cria Response{Success: true, Data: schedule}
   • Serializa para JSON
   • Define status 200

9. Middleware finaliza
   • Logging: "GET /schedule - 234ms"

10. Cliente recebe resposta JSON
    {
      "success": true,
      "data": [...],
      "count": 50
    }
```

## 🧩 Componentes Detalhados

### 1. **cmd/api/main.go**
- Entry point da aplicação
- Inicializa todas as dependências
- Configura servidor HTTP
- Implementa graceful shutdown

### 2. **internal/models/**
Structs de dados puros, sem lógica

**tvmaze.go:**
- Show, Episode, Schedule
- Network, Country, Image

**github.go:**
- GitHubUser

**response.go:**
- Response (formato padrão de resposta)

### 3. **internal/clients/**
Comunicação com APIs externas

**TVMazeClient:**
```go
type TVMazeClient struct {
    httpClient *http.Client
    baseURL    string
}
```

**GitHubClient:**
```go
type GitHubClient struct {
    httpClient *http.Client
    baseURL    string
}
```

### 4. **internal/services/**
Lógica de negócio

**TVMazeService:**
- Filtra por gênero
- Calcula "now playing"
- Valida entradas

**GitHubService:**
- Valida username
- Pode adicionar cache (futuro)

### 5. **internal/handlers/**
Interface HTTP

**TVMazeHandler:**
- Extrai query params
- Valida HTTP
- Formata JSON

**GitHubHandler:**
- Similar ao TVMaze

### 6. **internal/middleware/**
Cross-cutting concerns

- Logging
- CORS
- Auth (futuro)
- Metrics (futuro)

### 7. **internal/router/**
Configuração de rotas

```go
mux.HandleFunc("/schedule", middleware.Logging(handler.Schedule))
```

## 🎯 Princípios de Design

### 1. **Separation of Concerns**
Cada camada tem responsabilidade única

### 2. **Dependency Injection**
Dependências injetadas via construtores
```go
client := clients.NewTVMazeClient()
service := services.NewTVMazeService(client)
handler := handlers.NewTVMazeHandler(service)
```

### 3. **Testabilidade**
Cada camada testável isoladamente

### 4. **Manutenibilidade**
Código organizado e documentado

### 5. **Escalabilidade**
Fácil adicionar novos recursos

## 🔮 Evoluções Futuras

### 1. Adicionar Cache
```
Service → Cache → Client
```

### 2. Adicionar Database
```
Service → Repository → Database
```

### 3. Adicionar Queue
```
Handler → Queue → Worker
```

### 4. Adicionar gRPC
```
gRPC Handler → Service → Client
```

## 📚 Referências

- [Standard Go Project Layout](https://github.com/golang-standards/project-layout)
- [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [Effective Go](https://golang.org/doc/effective_go)

---

**Última atualização:** Janeiro 2026  
**Versão:** 3.0.0
