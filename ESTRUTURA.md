# 🏗️ Estrutura do Projeto

Este documento explica a nova estrutura profissional do projeto Go.

## 📁 Estrutura de Diretórios

```
goLang/
├── cmd/                          # Aplicações principais
│   └── api/
│       └── main.go              # Entry point da aplicação
│
├── internal/                     # Código privado da aplicação
│   ├── models/                  # Modelos de dados
│   │   ├── tvmaze.go           # Structs TVMaze (Show, Episode, Schedule, etc)
│   │   ├── github.go           # Structs GitHub (GitHubUser)
│   │   └── response.go         # Struct de resposta padrão da API
│   │
│   ├── clients/                 # Clientes HTTP para APIs externas
│   │   ├── tvmaze.go           # Cliente TVMaze API
│   │   └── github.go           # Cliente GitHub API
│   │
│   ├── services/                # Lógica de negócio
│   │   ├── tvmaze.go           # Serviço TVMaze (filtros, validações, etc)
│   │   └── github.go           # Serviço GitHub
│   │
│   ├── handlers/                # Handlers HTTP (camada de apresentação)
│   │   ├── tvmaze.go           # Handlers TVMaze
│   │   ├── github.go           # Handlers GitHub
│   │   └── docs.go             # Handler de documentação
│   │
│   ├── middleware/              # Middlewares HTTP
│   │   └── middleware.go       # Logging, CORS, etc
│   │
│   └── router/                  # Configuração de rotas
│       └── router.go           # Setup de todas as rotas
│
├── pkg/                         # Código reutilizável (bibliotecas públicas)
│   └── utils/
│       └── strings.go          # Utilidades para strings
│
├── examples/                    # Exemplos de código
│   └── primeiroGoLang.go       # Primeiro exemplo Go
│
├── docs/                        # Documentação adicional
│   ├── ARQUITETURA.md
│   ├── DEPLOY_GUIDE.md
│   └── TVMAZE_RESUMO.md
│
├── go.mod                       # Dependências do módulo Go
├── Makefile                     # Comandos make
├── Dockerfile                   # Configuração Docker
├── docker-compose.yml           # Orquestração Docker
├── README.md                    # Documentação principal
└── LICENSE                      # Licença MIT
```

## 🎯 Camadas da Aplicação

### 1. **cmd/** - Aplicações
- Entry point da aplicação
- Inicializa dependências
- Configura servidor HTTP
- Implementa graceful shutdown

### 2. **internal/models/** - Modelos de Dados
- Structs que representam os dados
- Tags JSON para serialização
- Sem lógica de negócio

### 3. **internal/clients/** - Clientes HTTP
- Comunicação com APIs externas
- Requisições HTTP (GET, POST, etc)
- Parsing de respostas
- Tratamento de erros de rede

### 4. **internal/services/** - Lógica de Negócio
- Orquestra chamadas aos clientes
- Validações
- Transformações de dados
- Filtros e buscas
- Regras de negócio

### 5. **internal/handlers/** - Handlers HTTP
- Recebe requisições HTTP
- Valida parâmetros
- Chama serviços
- Formata respostas JSON
- Trata erros HTTP

### 6. **internal/middleware/** - Middlewares
- Logging de requisições
- CORS
- Autenticação (futuro)
- Rate limiting (futuro)

### 7. **internal/router/** - Roteamento
- Registro de todas as rotas
- Associa rotas aos handlers
- Aplica middlewares

### 8. **pkg/utils/** - Utilitários
- Funções auxiliares reutilizáveis
- Sem dependências internas
- Pode ser usado por outras aplicações

## 🔄 Fluxo de Requisição

```
Cliente HTTP
    ↓
Router (registra rotas)
    ↓
Middleware (logging, CORS)
    ↓
Handler (valida parâmetros)
    ↓
Service (lógica de negócio)
    ↓
Client (chamada à API externa)
    ↓
Service (processa resposta)
    ↓
Handler (formata JSON)
    ↓
Cliente HTTP (resposta)
```

## 🏃 Como Executar

### Executar com Go
```bash
go run cmd/api/main.go
```

### Compilar e executar
```bash
make build
./bin/api-server
```

### Com Docker
```bash
docker build -t tvmaze-api .
docker run -p 8080:8080 tvmaze-api
```

### Com Make
```bash
make run          # Executar localmente
make build        # Compilar
make test         # Testes
make docker-build # Build Docker
```

## 📚 Princípios Aplicados

### 1. **Separation of Concerns**
Cada camada tem uma responsabilidade específica

### 2. **Dependency Injection**
Dependências injetadas via construtores

### 3. **Interface Segregation**
Cada pacote expõe apenas o necessário

### 4. **Single Responsibility**
Cada arquivo/struct tem uma responsabilidade

### 5. **Clean Architecture**
- Camadas bem definidas
- Dependências apontam para dentro
- Fácil de testar
- Fácil de manter

## 🧪 Testes

### Estrutura de Testes
```
internal/
  services/
    tvmaze_test.go
    github_test.go
  handlers/
    tvmaze_test.go
    github_test.go
  clients/
    tvmaze_test.go
    github_test.go
```

### Executar testes
```bash
make test                # Todos os testes
make test-coverage       # Com cobertura
go test ./internal/...   # Apenas internal
```

## 🚀 Benefícios da Nova Estrutura

✅ **Manutenibilidade**: Código organizado e fácil de encontrar
✅ **Testabilidade**: Cada camada pode ser testada isoladamente
✅ **Escalabilidade**: Fácil adicionar novos endpoints/features
✅ **Reusabilidade**: Código pode ser reutilizado
✅ **Padrão de Mercado**: Estrutura reconhecida pela comunidade Go
✅ **Onboarding**: Novos desenvolvedores entendem rapidamente
✅ **Produção-Ready**: Pronto para ambientes profissionais

## 📖 Convenções Go

- `internal/`: Código privado, não pode ser importado por outros projetos
- `pkg/`: Código público, pode ser importado
- `cmd/`: Aplicações executáveis
- Nomes de pacotes em minúsculas
- Construtores começam com `New`
- Interfaces terminam com `er` (ex: `Reader`, `Writer`)

## 🔧 Próximos Passos

1. ✅ Refatoração completa da estrutura
2. 🔲 Adicionar testes unitários para todas as camadas
3. 🔲 Implementar interfaces para facilitar mocks
4. 🔲 Adicionar configuração via arquivo/env vars
5. 🔲 Implementar cache (Redis)
6. 🔲 Adicionar métricas e monitoring
7. 🔲 CI/CD pipeline
8. 🔲 OpenAPI/Swagger spec
