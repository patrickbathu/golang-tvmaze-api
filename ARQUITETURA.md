# 🏗️ Arquitetura da TVMaze API

## 📊 Diagrama de Fluxo

```
┌─────────────┐
│   Cliente   │ (Browser, App, cURL)
└──────┬──────┘
       │ HTTP Request
       │
       ▼
┌─────────────────────────────────────────────────────┐
│              TVMaze API (Go Server)                  │
│  ┌──────────────────────────────────────────────┐  │
│  │         Logging Middleware                    │  │
│  └──────────────┬───────────────────────────────┘  │
│                 │                                    │
│  ┌──────────────▼───────────────────────────────┐  │
│  │          Router (http.ServeMux)              │  │
│  │  • GET /                                      │  │
│  │  • GET /schedule?country=XX                   │  │
│  │  • GET /search?q=NOME                         │  │
│  │  • GET /show?id=ID                            │  │
│  └──────────────┬───────────────────────────────┘  │
│                 │                                    │
│  ┌──────────────▼───────────────────────────────┐  │
│  │           Handlers                            │  │
│  │  • homeHandler()                              │  │
│  │  • scheduleHandler()                          │  │
│  │  • searchHandler()                            │  │
│  │  • showHandler()                              │  │
│  └──────────────┬───────────────────────────────┘  │
│                 │                                    │
│  ┌──────────────▼───────────────────────────────┐  │
│  │      Business Logic Functions                 │  │
│  │  • getTodaySchedule()                         │  │
│  │  • searchShow()                               │  │
│  │  • getShowByID()                              │  │
│  └──────────────┬───────────────────────────────┘  │
│                 │                                    │
│  ┌──────────────▼───────────────────────────────┐  │
│  │         HTTP Client (httpClient)              │  │
│  │  • Timeout: 15s                               │  │
│  │  • User-Agent header                          │  │
│  └──────────────┬───────────────────────────────┘  │
└─────────────────┼────────────────────────────────────┘
                  │ HTTPS Request
                  │
                  ▼
          ┌──────────────┐
          │  TVMaze API  │ (api.tvmaze.com)
          │  (External)  │
          └──────────────┘
```

## 🔄 Fluxo de uma Requisição

### Exemplo: GET /schedule?country=BR

```
1. Cliente → API Go
   GET /schedule?country=BR

2. Logging Middleware
   📨 Log: "GET /schedule"

3. Router
   Direciona para scheduleHandler()

4. scheduleHandler()
   • Valida parâmetro country
   • Chama getTodaySchedule("BR")

5. getTodaySchedule()
   • Formata data atual (2026-01-03)
   • Monta URL: https://api.tvmaze.com/schedule?country=BR&date=2026-01-03
   • Faz requisição HTTP

6. API Externa (TVMaze)
   • Processa requisição
   • Retorna JSON com programação

7. getTodaySchedule()
   • Lê resposta
   • Faz parse do JSON
   • Retorna []Schedule

8. scheduleHandler()
   • Monta Response{success, data, count}
   • Serializa para JSON
   • Retorna ao cliente

9. Logging Middleware
   ✅ Log: "GET /schedule - 523ms"

10. Cliente
    Recebe JSON com programação
```

## 🧱 Componentes

### 1️⃣ Structs (Modelos de Dados)

```go
Response        → Resposta padrão da API
Schedule        → Item da programação
Show            → Informações do show
Episode         → Episódio do show
Network         → Rede de TV
Country         → País
Image           → URLs de imagens
```

### 2️⃣ Handlers (Controladores HTTP)

```go
homeHandler()      → GET /
scheduleHandler()  → GET /schedule
searchHandler()    → GET /search
showHandler()      → GET /show
```

### 3️⃣ Business Logic (Lógica de Negócio)

```go
getTodaySchedule() → Busca programação
searchShow()       → Busca shows
getShowByID()      → Busca detalhes
```

### 4️⃣ Middleware

```go
loggingMiddleware() → Logging de requisições
```

### 5️⃣ HTTP Client

```go
httpClient → Cliente global com timeout
```

## 🔐 Camadas da Aplicação

```
┌─────────────────────────────────┐
│   Presentation Layer             │ ← Handlers (HTTP)
├─────────────────────────────────┤
│   Business Logic Layer           │ ← Functions (getTodaySchedule)
├─────────────────────────────────┤
│   Integration Layer              │ ← HTTP Client (TVMaze API)
└─────────────────────────────────┘
```

## 🎯 Padrões de Design Utilizados

### 1. Handler Pattern
```go
func handler(w http.ResponseWriter, r *http.Request) {
    // Processar requisição
    // Retornar resposta
}
```

### 2. Middleware Pattern
```go
func middleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Fazer algo antes
        next(w, r)
        // Fazer algo depois
    }
}
```

### 3. Repository Pattern (simplificado)
```go
// Funções que abstraem acesso a dados externos
getTodaySchedule()
searchShow()
getShowByID()
```

## 🚀 Ciclo de Vida do Servidor

```
1. main()
   ↓
2. Configurar rotas (mux)
   ↓
3. Criar servidor HTTP
   ↓
4. Iniciar goroutine do servidor
   ↓
5. Aguardar sinal de interrupção
   ↓
6. Graceful shutdown (30s timeout)
   ↓
7. Servidor encerra
```

## 🔄 Concorrência

```
┌──────────────────────────────────┐
│  Main Goroutine                   │
│  • Inicializa servidor            │
│  • Aguarda sinais (Ctrl+C)        │
└──────────────────────────────────┘
         │
         │ Spawns
         │
         ▼
┌──────────────────────────────────┐
│  Server Goroutine                 │
│  • ListenAndServe()               │
│  • Aceita conexões                │
└──────────────────────────────────┘
         │
         │ Spawns (automático)
         │
         ▼
┌──────────────────────────────────┐
│  Handler Goroutines               │
│  • Uma por requisição             │
│  • Processa em paralelo           │
└──────────────────────────────────┘
```

## 📦 Dependências

```
Standard Library (stdlib):
├── context      → Graceful shutdown
├── encoding/json → JSON parsing
├── fmt          → Formatação
├── io           → I/O operations
├── log          → Logging
├── net/http     → HTTP server/client
├── os           → Variáveis de ambiente
├── os/signal    → Sinais do sistema
├── syscall      → System calls
└── time         → Datas e timeouts
```

## 🔒 Segurança

```
✅ Timeouts configurados (evita DoS)
✅ Defer para fechar recursos
✅ Error handling robusto
✅ CORS habilitado (controlado)
✅ User-Agent configurado
✅ Graceful shutdown
✅ Status codes apropriados
```

## 📊 Performance

```
Características:
• Baixa latência (dependente da TVMaze API)
• Stateless (fácil de escalar horizontalmente)
• Connection reuse (HTTP client)
• Timeouts configurados
• Sem bloqueios desnecessários
```

## 🧪 Testing Strategy

```
Unit Tests:
├── Handler Tests (HTTP)
├── Function Tests (Business Logic)
├── Middleware Tests
└── Integration Tests (com API real)

Coverage: ~85%
```

## 📈 Escalabilidade

### Vertical (mais recursos)
```
• Aumentar CPU/RAM
• Ajustar timeouts
• Connection pooling
```

### Horizontal (mais instâncias)
```
• Load balancer (Nginx, HAProxy)
• Múltiplas instâncias da API
• Session-less (stateless)
```

### Cache
```
• Redis para respostas
• TTL baseado em dados
• Cache-Control headers
```

## 🎯 Melhorias Futuras

```
1. Cache Layer (Redis)
   Client → API Go → Cache → TVMaze API

2. Database Layer (PostgreSQL)
   Para favoritos, histórico, etc.

3. Message Queue (RabbitMQ)
   Para processamento assíncrono

4. Service Mesh (Istio)
   Para múltiplos microserviços
```

## 🏁 Conclusão

A arquitetura é:
- ✅ **Simples** - Fácil de entender
- ✅ **Robusta** - Error handling adequado
- ✅ **Escalável** - Stateless design
- ✅ **Testável** - Boa separação de concerns
- ✅ **Maintível** - Código organizado
- ✅ **Performática** - Timeouts e HTTP client otimizado

---

**Padrão RESTful + Clean Architecture**
