# Build stage
FROM golang:1.21-alpine AS builder

# Instalar certificados SSL para requisições HTTPS
RUN apk add --no-cache ca-certificates git

WORKDIR /app

# Copiar arquivos de dependências
COPY go.mod ./
RUN go mod download

# Copiar código fonte
COPY . .

# Build da aplicação
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api-server cmd/api/main.go

# Runtime stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copiar binário do build stage
COPY --from=builder /app/api-server .

# Expor porta
EXPOSE 8080

# Variáveis de ambiente
ENV PORT=8080

# Comando para iniciar
CMD ["./api-server"]
