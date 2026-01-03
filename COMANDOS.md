# 🛠️ Comandos Úteis - Go TVMaze API

## 🚀 Desenvolvimento

### Executar aplicação
```bash
# Método 1: go run
go run cmd/api/main.go

# Método 2: make
make run

# Método 3: compilar e executar
make build
./bin/api-server

# Com variável de ambiente (porta customizada)
PORT=3000 go run cmd/api/main.go
```

### Desenvolvimento com auto-reload (usando air)
```bash
# Instalar air
go install github.com/cosmtrek/air@latest

# Executar
air
```

## 🧪 Testes

```bash
# Todos os testes
go test ./...

# Testes verbose
go test -v ./...

# Testes com coverage
go test -cover ./...

# Coverage HTML
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Testar pacote específico
go test ./internal/services/...

# Executar teste específico
go test -run TestNewTVMazeService ./internal/services/

# Com make
make test
make test-coverage
```

## 🔍 Análise de Código

```bash
# Formatar código
go fmt ./...

# Verificar problemas
go vet ./...

# Lint (requer golangci-lint)
golangci-lint run

# Verificar imports não usados
goimports -l .

# Verificar vulnerabilidades
go list -json -m all | nancy sleuth
```

## 📦 Dependências

```bash
# Atualizar dependências
go mod tidy

# Verificar dependências
go mod verify

# Ver dependências
go list -m all

# Limpar cache
go clean -modcache

# Adicionar nova dependência
go get github.com/package/name
```

## 🐳 Docker

```bash
# Build da imagem
docker build -t tvmaze-api:latest .

# Executar container
docker run -p 8080:8080 tvmaze-api:latest

# Com docker-compose
docker-compose up
docker-compose up -d  # background
docker-compose down
docker-compose logs -f

# Com make
make docker-build
make docker-run
make docker-compose-up
make docker-compose-down
```

## 📊 Benchmarks

```bash
# Executar benchmarks
go test -bench=. ./...

# Com memory stats
go test -bench=. -benchmem ./...

# Benchmark específico
go test -bench=BenchmarkGetSchedule ./internal/services/
```

## 🔧 Build

```bash
# Build simples
go build -o bin/api-server cmd/api/main.go

# Build com flags de otimização
go build -ldflags="-s -w" -o bin/api-server cmd/api/main.go

# Cross-compile para Linux
GOOS=linux GOARCH=amd64 go build -o bin/api-server-linux cmd/api/main.go

# Cross-compile para Windows
GOOS=windows GOARCH=amd64 go build -o bin/api-server.exe cmd/api/main.go

# Build para múltiplas plataformas
make build-all  # se configurado no Makefile
```

## 🧹 Limpeza

```bash
# Limpar binários
rm -rf bin/

# Limpar cache de testes
go clean -testcache

# Limpar tudo
go clean -i -r -cache -testcache -modcache

# Com make
make clean
```

## 📡 Testar API (curl)

```bash
# Home
curl http://localhost:8080/

# Schedule
curl "http://localhost:8080/schedule?country=US"
curl "http://localhost:8080/schedule?country=BR"

# Search
curl "http://localhost:8080/search?q=friends"

# Show details
curl "http://localhost:8080/show?id=431"

# Genre
curl "http://localhost:8080/genre?genre=Drama&country=US"

# Now playing
curl "http://localhost:8080/now?country=US"

# GitHub user
curl "http://localhost:8080/api/user?username=torvalds"

# Com formatação JSON (jq)
curl -s http://localhost:8080/ | jq .

# Salvar resposta
curl -o response.json "http://localhost:8080/schedule?country=BR"

# Ver headers
curl -I http://localhost:8080/

# Verbose
curl -v http://localhost:8080/
```

## 📊 Profiling

```bash
# CPU profiling
go test -cpuprofile=cpu.prof -bench=. ./...
go tool pprof cpu.prof

# Memory profiling
go test -memprofile=mem.prof -bench=. ./...
go tool pprof mem.prof

# Executar com profiling
go run -cpuprofile=cpu.prof cmd/api/main.go
```

## 🌐 Desenvolvimento Web

```bash
# Abrir documentação no browser
open http://localhost:8080/docs

# Verificar se está rodando
curl -f http://localhost:8080/ && echo "✅ OK" || echo "❌ Falhou"

# Monitorar logs
tail -f logs/app.log  # se configurado

# Live reload (browser-sync)
browser-sync start --proxy "localhost:8080" --files "**/*"
```

## 📝 Git

```bash
# Verificar status
git status

# Adicionar arquivos da nova estrutura
git add cmd/ internal/ pkg/
git add ESTRUTURA.md MIGRACAO.md ARQUITETURA_V3.md

# Commit
git commit -m "refactor: migrar para arquitetura em camadas v3.0"

# Tag de versão
git tag v3.0.0
git push origin v3.0.0

# Ver mudanças
git diff
git log --oneline
```

## 🔍 Debugger

```bash
# Com delve
go install github.com/go-delve/delve/cmd/dlv@latest
dlv debug cmd/api/main.go

# Debugger com breakpoint
dlv debug cmd/api/main.go -- --config=dev.yaml
```

## 📈 Métricas e Monitoramento

```bash
# Expor métricas Prometheus (se implementado)
curl http://localhost:8080/metrics

# Health check
curl http://localhost:8080/health

# Ver processos Go
ps aux | grep api-server

# Usar pprof (se habilitado)
go tool pprof http://localhost:8080/debug/pprof/heap
```

## 🚀 Deploy

```bash
# Railway
railway up

# Heroku
git push heroku main

# Google Cloud Run
gcloud run deploy --source .

# Fly.io
fly deploy

# AWS Elastic Beanstalk
eb deploy
```

## 💡 Dicas Úteis

```bash
# Ver todas as variáveis de ambiente Go
go env

# Ver onde Go está instalado
which go

# Versão do Go
go version

# Documentação de um pacote
go doc fmt
go doc net/http

# Executar go mod tidy automaticamente
watch -n 5 go mod tidy

# Monitorar mudanças e executar testes
watch -n 2 go test ./...
```

## 📚 Recursos

- [Go Documentation](https://golang.org/doc/)
- [Effective Go](https://golang.org/doc/effective_go)
- [Go by Example](https://gobyexample.com/)
- [Project Layout](https://github.com/golang-standards/project-layout)

---

**💡 Dica**: Adicione esses comandos ao seu `.bashrc` ou `.zshrc` como aliases!

```bash
alias gorun='go run cmd/api/main.go'
alias gotest='go test -v ./...'
alias gobuild='go build -o bin/api-server cmd/api/main.go'
alias gofmt='go fmt ./...'
```
