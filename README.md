# 📺 TVMaze API - Go Microservice

API REST profissional em Go que consulta a programação de TV em tempo real usando a [TVMaze API](https://www.tvmaze.com/api) e informações de usuários do [GitHub](https://api.github.com).

**✨ Versão 3.0.0** - Arquitetura profissional com separação de camadas!

## 🚀 Como Executar

### Método 1: Executar diretamente
```bash
go run cmd/api/main.go
```

### Método 2: Compilar e executar
```bash
make build
./bin/api-server
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

### 6. Filtrar por gênero
```bash
curl "http://localhost:8080/genre?genre=Drama&country=US"
```

### 7. O que está passando agora
```bash
curl "http://localhost:8080/now?country=US"
```

### 8. Usuário do GitHub
```bash
curl "http://localhost:8080/api/user?username=torvalds"
```

### 9. Documentação Interativa
```
http://localhost:8080/docs
```

## 🧪 Testes

```bash
# Executar testes
make test

# Com coverage
make test-coverage

# Testar APIs
make test-api
```

## 📦 Estrutura do Projeto (Arquitetura Profissional)

```
goLang/
├── cmd/
│   └── api/
│       └── main.go              # 🎯 Entry point
├── internal/
│   ├── models/                  # 📊 Modelos de dados
│   │   ├── tvmaze.go
│   │   ├── github.go
│   │   └── response.go
│   ├── clients/                 # 🌐 Clientes HTTP
│   │   ├── tvmaze.go
│   │   └── github.go
│   ├── services/                # 💼 Lógica de negócio
│   │   ├── tvmaze.go
│   │   ├── tvmaze_test.go
│   │   ├── github.go
│   │   └── github_test.go
│   ├── handlers/                # 🎮 Handlers HTTP
│   │   ├── tvmaze.go
│   │   ├── github.go
│   │   └── docs.go
│   ├── middleware/              # 🔧 Middlewares
│   │   └── middleware.go
│   └── router/                  # 🛣️ Roteamento
│       └── router.go
├── pkg/
│   └── utils/                   # 🔨 Utilitários
│       └── strings.go
├── go.mod
├── Makefile
├── Dockerfile
└── README.md
```

### 🏗️ Arquitetura em Camadas

**Cliente HTTP** → **Router** → **Middleware** → **Handler** → **Service** → **Client** → **API Externa**

Veja [ESTRUTURA.md](ESTRUTURA.md) para detalhes completos da arquitetura.

## 📚 Deploy

Ver documentação completa em:
- `ESTRUTURA.md` - Arquitetura e organização do projeto
- `DEPLOY_TVMAZE.md` - Guia completo de deploy
- `ARQUITETURA.md` - Diagrama da aplicação

### Deploy Rápido (Railway)
```bash
railway login
railway init
railway up
```

## 📚 Documentação Completa

- **ESTRUTURA.md** - 🆕 Arquitetura profissional do projeto
- **DEPLOY_TVMAZE.md** - Guia de deploy em várias plataformas
- **README_TVMAZE.md** - Documentação detalhada da API
- **TVMAZE_RESUMO.md** - Resumo completo do projeto
- **ARQUITETURA.md** - Diagrama e explicação da arquitetura

## 🎯 Conceitos Go Implementados

- ✅ Arquitetura em camadas (Clean Architecture)
- ✅ Dependency Injection
- ✅ Separation of Concerns
- ✅ Structs e JSON tags
- ✅ HTTP Server e Client
- ✅ Error handling
- ✅ Context e graceful shutdown
- ✅ Middleware pattern
- ✅ Testing (unit tests)
- ✅ Goroutines e Channels
- ✅ Package organization (internal, pkg, cmd)

## 🌟 Diferenciais Profissionais

✨ **Arquitetura em Camadas**: Separação clara de responsabilidades  
✨ **Testabilidade**: Cada camada pode ser testada isoladamente  
✨ **Manutenibilidade**: Código organizado e fácil de manter  
✨ **Escalabilidade**: Fácil adicionar novos recursos  
✨ **Padrões de Mercado**: Estrutura seguindo best practices Go  
✨ **Documentação Interativa**: Interface web para testar endpoints  
✨ **Graceful Shutdown**: Encerramento seguro do servidor  
✨ **Production-Ready**: Pronto para ambientes profissionais  

## 📝 Licença

MIT

---

Criado como projeto de aprendizado de Go Lang 🐹  
Refatorado seguindo padrões profissionais de mercado 🚀
