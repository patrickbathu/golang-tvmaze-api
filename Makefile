# Makefile para facilitar comandos comuns

.PHONY: help build run docker-build docker-run docker-stop docker-logs test clean

help: ## Mostrar esta mensagem de ajuda
	@echo "Comandos disponíveis:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

build: ## Compilar a aplicação
	go build -o bin/api-server api.go

run: ## Executar a aplicação localmente
	go run api.go

docker-build: ## Build da imagem Docker
	docker build -t github-api:latest .

docker-run: ## Rodar container Docker
	docker run -d -p 8080:8080 --name github-api-service github-api:latest

docker-compose-up: ## Iniciar com docker-compose
	docker-compose up -d

docker-compose-down: ## Parar docker-compose
	docker-compose down

docker-stop: ## Parar container Docker
	docker stop github-api-service || true
	docker rm github-api-service || true

docker-logs: ## Ver logs do container
	docker logs -f github-api-service

test: ## Executar testes unitários
	go test -v

test-coverage: ## Executar testes com cobertura
	go test -v -cover -coverprofile=coverage.out
	go tool cover -html=coverage.out

test-api: ## Testar a API via HTTP
	@echo "Testando endpoint raiz..."
	@curl -s http://localhost:8080/ | python3 -m json.tool
	@echo "\nTestando busca de usuário..."
	@curl -s "http://localhost:8080/user?username=torvalds" | python3 -m json.tool

benchmark: ## Executar benchmarks
	go test -bench=. -benchmem

clean: ## Limpar arquivos gerados
	rm -rf bin/
	docker rmi github-api:latest || true
