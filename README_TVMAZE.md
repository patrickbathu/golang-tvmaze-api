# 📺 TVMaze API - Go Microservice

API REST em Go que consulta a programação de TV do dia usando a [TVMaze API](https://www.tvmaze.com/api).

## 🚀 Características

- ✅ Consulta programação de TV por país
- ✅ Busca de shows por nome
- ✅ Detalhes de shows específicos
- ✅ Graceful shutdown
- ✅ Middleware de logging
- ✅ CORS habilitado
- ✅ Testes unitários
- ✅ Timeout configurado
- ✅ Resposta padronizada em JSON

## 📋 Pré-requisitos

- Go 1.21+ instalado
- Conexão com internet (para acessar TVMaze API)

## 🏃 Como Executar

### Executar localmente

```bash
# Executar diretamente
go run tvmaze-api.go

# Ou compilar e executar
go build -o tvmaze-server tvmaze-api.go
./tvmaze-server
```

### Executar com porta customizada

```bash
PORT=3000 go run tvmaze-api.go
```

## 🔌 Endpoints

### 1. Informações da API
```bash
GET /
```

**Resposta:**
```json
{
  "message": "📺 API Go - TVMaze Schedule",
  "version": "1.0.0",
  "date": "2026-01-03",
  "endpoints": {...},
  "examples": [...]
}
```

### 2. Programação de Hoje
```bash
GET /schedule?country=US
```

**Parâmetros:**
- `country` (opcional): Código do país (padrão: US)
  - US (Estados Unidos)
  - BR (Brasil)
  - GB (Reino Unido)
  - etc.

**Exemplo:**
```bash
curl "http://localhost:8080/schedule?country=US"
```

**Resposta:**
```json
{
  "success": true,
  "data": [
    {
      "id": 3487624,
      "airdate": "2026-01-03",
      "airtime": "20:00",
      "show": {
        "id": 431,
        "name": "Friends",
        "type": "Scripted",
        "language": "English",
        "genres": ["Comedy", "Romance"],
        "status": "Ended",
        "premiered": "1994-09-22",
        "summary": "...",
        "image": {...},
        "network": {...}
      }
    }
  ],
  "count": 100
}
```

### 3. Buscar Shows
```bash
GET /search?q=NOME_DO_SHOW
```

**Parâmetros:**
- `q` (obrigatório): Nome do show a buscar

**Exemplo:**
```bash
curl "http://localhost:8080/search?q=friends"
curl "http://localhost:8080/search?q=game+of+thrones"
```

**Resposta:**
```json
{
  "success": true,
  "data": [
    {
      "score": 0.9036184,
      "show": {
        "id": 431,
        "name": "Friends",
        "type": "Scripted",
        "language": "English",
        "genres": ["Comedy", "Romance"],
        ...
      }
    }
  ],
  "count": 10
}
```

### 4. Detalhes de um Show
```bash
GET /show?id=ID_DO_SHOW
```

**Parâmetros:**
- `id` (obrigatório): ID do show na TVMaze

**Exemplo:**
```bash
curl "http://localhost:8080/show?id=431"
```

**Resposta:**
```json
{
  "success": true,
  "data": {
    "id": 431,
    "name": "Friends",
    "type": "Scripted",
    "language": "English",
    "genres": ["Comedy", "Romance"],
    "status": "Ended",
    "premiered": "1994-09-22",
    "summary": "Six young people from New York City...",
    "image": {
      "medium": "https://...",
      "original": "https://..."
    },
    "network": {
      "id": 1,
      "name": "NBC",
      "country": {
        "name": "United States",
        "code": "US"
      }
    }
  }
}
```

## 🧪 Testes

Execute os testes unitários:

```bash
# Rodar todos os testes
go test -v tvmaze-api.go tvmaze-api_test.go

# Rodar com coverage
go test -cover tvmaze-api.go tvmaze-api_test.go

# Gerar relatório de coverage
go test -coverprofile=coverage.out tvmaze-api.go tvmaze-api_test.go
go tool cover -html=coverage.out
```

## 📦 Estrutura do Projeto

```
.
├── tvmaze-api.go           # Código principal da API
├── tvmaze-api_test.go      # Testes unitários
├── api.go                  # API GitHub (outro exemplo)
├── api_test.go             # Testes da API GitHub
├── primeiroGoLang.go       # Primeiro exemplo Hello World
├── go.mod                  # Dependências do módulo
├── Dockerfile              # Container Docker
├── docker-compose.yml      # Orquestração Docker
├── Makefile                # Comandos úteis
└── README.md               # Esta documentação
```

## 🐳 Docker

### Build da imagem

```bash
docker build -t tvmaze-api .
```

### Executar container

```bash
docker run -p 8080:8080 tvmaze-api
```

### Docker Compose

```bash
docker-compose up
```

## 🔧 Variáveis de Ambiente

| Variável | Descrição | Padrão |
|----------|-----------|--------|
| `PORT` | Porta do servidor | `8080` |

## 📊 Exemplos de Uso

### Usando cURL

```bash
# Programação de hoje nos EUA
curl "http://localhost:8080/schedule?country=US"

# Programação de hoje no Brasil
curl "http://localhost:8080/schedule?country=BR"

# Buscar Breaking Bad
curl "http://localhost:8080/search?q=breaking+bad"

# Detalhes de Friends (ID 431)
curl "http://localhost:8080/show?id=431"
```

### Usando JavaScript (fetch)

```javascript
// Buscar programação
fetch('http://localhost:8080/schedule?country=US')
  .then(res => res.json())
  .then(data => console.log(data));

// Buscar show
fetch('http://localhost:8080/search?q=friends')
  .then(res => res.json())
  .then(data => console.log(data));
```

### Usando Python (requests)

```python
import requests

# Buscar programação
response = requests.get('http://localhost:8080/schedule?country=US')
data = response.json()
print(data)

# Buscar show
response = requests.get('http://localhost:8080/search?q=friends')
data = response.json()
print(data)
```

## 🎯 Conceitos Go Implementados

1. **Structs e JSON Tags** - Modelagem de dados
2. **HTTP Client** - Requisições para API externa
3. **HTTP Server** - Criação de endpoints REST
4. **Error Handling** - Tratamento robusto de erros
5. **Middleware** - Logging de requisições
6. **Context** - Graceful shutdown
7. **Testing** - Testes unitários completos
8. **Goroutines** - Servidor HTTP assíncrono
9. **Channels** - Comunicação entre goroutines
10. **Time** - Timeouts e formatação de datas

## 📚 Próximos Passos

- [ ] Adicionar cache Redis
- [ ] Implementar rate limiting
- [ ] Adicionar autenticação JWT
- [ ] Criar endpoints para episódios
- [ ] Adicionar banco de dados (favoritos)
- [ ] Implementar pagination
- [ ] Adicionar Swagger/OpenAPI
- [ ] Criar dashboard web
- [ ] Deploy em cloud (Railway, Render, Fly.io)
- [ ] CI/CD com GitHub Actions

## 🤝 Contribuindo

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues ou pull requests.

## 📝 Licença

Este projeto é open source e está disponível sob a licença MIT.

## 🔗 Links Úteis

- [TVMaze API Documentation](https://www.tvmaze.com/api)
- [Go Documentation](https://go.dev/doc/)
- [Go by Example](https://gobyexample.com/)

## 👨‍💻 Autor

Criado como projeto de aprendizado de Go Lang.

---

⭐ Se este projeto te ajudou, considere dar uma estrela!
