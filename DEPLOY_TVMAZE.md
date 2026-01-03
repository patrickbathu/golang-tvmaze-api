# 🚀 Guia de Deploy - TVMaze API

Este guia mostra várias formas de fazer deploy da API TVMaze como um microserviço em produção.

## 📋 Conteúdo

1. [Deploy Local (Systemd)](#1-deploy-local-systemd)
2. [Deploy com Docker](#2-deploy-com-docker)
3. [Deploy no Railway](#3-deploy-no-railway)
4. [Deploy no Render](#4-deploy-no-render)
5. [Deploy no Fly.io](#5-deploy-no-flyio)
6. [Deploy no Google Cloud Run](#6-deploy-no-google-cloud-run)

---

## 1. Deploy Local (Systemd)

Para rodar como serviço no Linux:

### Passo 1: Compilar o binário

```bash
go build -o /usr/local/bin/tvmaze-api tvmaze-api.go
```

### Passo 2: Criar arquivo de serviço

Criar arquivo `/etc/systemd/system/tvmaze-api.service`:

```ini
[Unit]
Description=TVMaze API Service
After=network.target

[Service]
Type=simple
User=www-data
WorkingDirectory=/opt/tvmaze-api
ExecStart=/usr/local/bin/tvmaze-api
Restart=always
RestartSec=10
Environment="PORT=8080"

# Logging
StandardOutput=journal
StandardError=journal
SyslogIdentifier=tvmaze-api

[Install]
WantedBy=multi-user.target
```

### Passo 3: Habilitar e iniciar

```bash
# Recarregar systemd
sudo systemctl daemon-reload

# Habilitar serviço
sudo systemctl enable tvmaze-api

# Iniciar serviço
sudo systemctl start tvmaze-api

# Ver status
sudo systemctl status tvmaze-api

# Ver logs
sudo journalctl -u tvmaze-api -f
```

---

## 2. Deploy com Docker

### Build e Run

```bash
# Build da imagem
docker build -t tvmaze-api .

# Executar container
docker run -d \
  --name tvmaze-api \
  -p 8080:8080 \
  --restart unless-stopped \
  tvmaze-api

# Ver logs
docker logs -f tvmaze-api

# Parar container
docker stop tvmaze-api

# Remover container
docker rm tvmaze-api
```

### Docker Compose

```bash
# Iniciar
docker-compose up -d

# Ver logs
docker-compose logs -f

# Parar
docker-compose down
```

---

## 3. Deploy no Railway

[Railway](https://railway.app/) é uma plataforma moderna de deploy com tier gratuito.

### Método 1: Via GitHub

1. Faça push do código para GitHub
2. Acesse https://railway.app/
3. Clique em "New Project"
4. Selecione "Deploy from GitHub repo"
5. Escolha seu repositório
6. Railway detectará automaticamente Go
7. Configure variáveis de ambiente (se necessário)

### Método 2: Via Railway CLI

```bash
# Instalar Railway CLI
npm install -g @railway/cli

# Login
railway login

# Inicializar projeto
railway init

# Deploy
railway up

# Ver logs
railway logs

# Abrir no navegador
railway open
```

### Configuração Railway (`railway.toml`):

```toml
[build]
builder = "nixpacks"

[deploy]
startCommand = "go run tvmaze-api.go"
restartPolicyType = "ON_FAILURE"
restartPolicyMaxRetries = 10
```

---

## 4. Deploy no Render

[Render](https://render.com/) oferece deploy gratuito para aplicações web.

### Passos:

1. Acesse https://render.com/
2. Clique em "New +" → "Web Service"
3. Conecte seu repositório GitHub
4. Configure:
   - **Name**: tvmaze-api
   - **Environment**: Go
   - **Build Command**: `go build -o server tvmaze-api.go`
   - **Start Command**: `./server`
   - **Instance Type**: Free

### Configuração Render (`render.yaml`):

```yaml
services:
  - type: web
    name: tvmaze-api
    env: go
    buildCommand: go build -o server tvmaze-api.go
    startCommand: ./server
    envVars:
      - key: PORT
        value: 8080
```

---

## 5. Deploy no Fly.io

[Fly.io](https://fly.io/) permite deploy global de aplicações.

### Passos:

```bash
# Instalar flyctl
curl -L https://fly.io/install.sh | sh

# Login
fly auth login

# Inicializar app
fly launch

# Deploy
fly deploy

# Abrir app
fly open

# Ver logs
fly logs

# Ver status
fly status
```

### Configuração Fly.io (`fly.toml`):

```toml
app = "tvmaze-api"
primary_region = "gru" # São Paulo

[build]
  builder = "paketobuildpacks/builder:base"

[env]
  PORT = "8080"

[http_service]
  internal_port = 8080
  force_https = true
  auto_stop_machines = true
  auto_start_machines = true
  min_machines_running = 0

[[vm]]
  cpu_kind = "shared"
  cpus = 1
  memory_mb = 256
```

---

## 6. Deploy no Google Cloud Run

Cloud Run é serverless e escala automaticamente.

### Passos:

```bash
# Instalar gcloud CLI
# https://cloud.google.com/sdk/docs/install

# Login
gcloud auth login

# Configurar projeto
gcloud config set project MEU_PROJETO_ID

# Build com Cloud Build
gcloud builds submit --tag gcr.io/MEU_PROJETO_ID/tvmaze-api

# Deploy no Cloud Run
gcloud run deploy tvmaze-api \
  --image gcr.io/MEU_PROJETO_ID/tvmaze-api \
  --platform managed \
  --region us-central1 \
  --allow-unauthenticated \
  --port 8080

# Ver logs
gcloud run logs read tvmaze-api
```

---

## 🔧 Configurações Recomendadas

### Variáveis de Ambiente

```bash
PORT=8080                    # Porta do servidor
GO_ENV=production            # Ambiente
LOG_LEVEL=info              # Nível de log
```

### Nginx Reverse Proxy (opcional)

Se você quiser usar Nginx na frente:

```nginx
server {
    listen 80;
    server_name api.seudominio.com;

    location / {
        proxy_pass http://localhost:8080;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection 'upgrade';
        proxy_set_header Host $host;
        proxy_cache_bypass $http_upgrade;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
```

---

## 📊 Monitoramento

### Health Check Endpoint

Adicione um endpoint de health check:

```go
func healthHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "status": "healthy",
        "time": time.Now().Format(time.RFC3339),
    })
}
```

### Logs

Use ferramentas como:
- **Datadog** - Monitoramento completo
- **New Relic** - APM
- **Sentry** - Error tracking
- **Prometheus + Grafana** - Métricas

---

## 🔒 Segurança

1. **HTTPS**: Use sempre HTTPS em produção
2. **Rate Limiting**: Implemente rate limiting
3. **CORS**: Configure CORS adequadamente
4. **API Keys**: Considere autenticação
5. **Firewall**: Configure firewall adequado

---

## 📈 Otimizações

1. **Cache**: Use Redis para cache de respostas
2. **CDN**: Use CloudFlare ou similar
3. **Compression**: Habilite gzip
4. **Connection Pooling**: Reuse HTTP connections
5. **Graceful Shutdown**: Já implementado ✅

---

## 🎯 Comparação de Plataformas

| Plataforma | Tier Gratuito | Facilidade | Escala | Custo |
|------------|---------------|------------|--------|-------|
| Railway | ✅ $5/mês | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐ | 💰💰 |
| Render | ✅ Limitado | ⭐⭐⭐⭐⭐ | ⭐⭐⭐ | 💰💰 |
| Fly.io | ✅ Sim | ⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ | 💰💰 |
| Cloud Run | ✅ Sim | ⭐⭐⭐ | ⭐⭐⭐⭐⭐ | 💰💰💰 |
| VPS | ❌ Não | ⭐⭐ | ⭐⭐⭐⭐ | 💰 |

---

## 🆘 Troubleshooting

### API não inicia

```bash
# Verificar se porta está em uso
lsof -i :8080

# Verificar logs
journalctl -u tvmaze-api -n 50

# Testar localmente
go run tvmaze-api.go
```

### Timeout em requisições

- Aumentar timeout do HTTP client
- Verificar conectividade com TVMaze API
- Usar retry logic

### Memory/CPU alto

- Implementar cache
- Usar connection pooling
- Profile com pprof

---

## 📚 Recursos Adicionais

- [Go Deployment Best Practices](https://go.dev/doc/)
- [Docker Best Practices](https://docs.docker.com/develop/dev-best-practices/)
- [Railway Docs](https://docs.railway.app/)
- [Render Docs](https://render.com/docs)
- [Fly.io Docs](https://fly.io/docs/)

---

🎉 **Pronto!** Agora você tem várias opções para fazer deploy da sua API TVMaze!
