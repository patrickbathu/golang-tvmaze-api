# 📦 Guia Completo de Deploy - GitHub API Microservice

## 🎯 Visão Geral

Este guia mostra como deixar sua API rodando 24/7 como um microserviço em produção.

---

## 🚀 Opções de Deploy (Do mais fácil ao mais complexo)

### 1️⃣ **Railway.app** (RECOMENDADO - GRATUITO)

✅ **Mais fácil e rápido**  
✅ **Deploy automático via Git**  
✅ **Free tier generoso**

```bash
# 1. Instalar Railway CLI
brew install railway

# 2. Login
railway login

# 3. Inicializar projeto
railway init

# 4. Deploy (automático!)
railway up

# 5. Abrir aplicação
railway open
```

**Ou via Web (sem CLI):**
1. Acesse: https://railway.app
2. Clique em "Start a New Project"
3. Conecte seu GitHub
4. Selecione o repositório
5. Railway detecta automaticamente o Dockerfile
6. Deploy automático! 🎉

---

### 2️⃣ **Render.com** (GRATUITO)

✅ **Interface simples**  
✅ **Free tier disponível**  
✅ **SSL automático**

**Passos:**
1. Acesse: https://render.com
2. Clique em "New +" → "Web Service"
3. Conecte seu repositório GitHub
4. Configurações automáticas do `render.yaml`
5. Clique em "Create Web Service"

**Configuração Manual:**
- **Build Command:** `go build -o bin/api-server api.go`
- **Start Command:** `./bin/api-server`
- **Environment:** Docker ou Go

---

### 3️⃣ **Fly.io** (GRATUITO)

✅ **Muito rápido**  
✅ **Global edge network**  
✅ **Free tier: 3 VMs compartilhadas**

```bash
# 1. Instalar Fly CLI
brew install flyctl

# 2. Login
fly auth login

# 3. Launch (cria fly.toml automaticamente)
fly launch --name github-api-service

# 4. Deploy
fly deploy

# 5. Abrir aplicação
fly open

# 6. Ver logs
fly logs
```

---

### 4️⃣ **Heroku** (Pago após Nov 2022)

```bash
# 1. Instalar Heroku CLI
brew tap heroku/brew && brew install heroku

# 2. Login
heroku login

# 3. Criar app
heroku create github-api-go-service

# 4. Deploy via Git
git push heroku main

# 5. Abrir app
heroku open

# 6. Ver logs
heroku logs --tail
```

---

### 5️⃣ **Google Cloud Run** (Pago após free tier)

✅ **Escala automático**  
✅ **Pay-per-use**  
✅ **Free tier: 2M requests/mês**

```bash
# 1. Instalar Google Cloud SDK
brew install --cask google-cloud-sdk

# 2. Login e configurar projeto
gcloud auth login
gcloud config set project SEU-PROJECT-ID

# 3. Build e push para Container Registry
gcloud builds submit --tag gcr.io/SEU-PROJECT-ID/github-api

# 4. Deploy no Cloud Run
gcloud run deploy github-api \
  --image gcr.io/SEU-PROJECT-ID/github-api \
  --platform managed \
  --region us-central1 \
  --allow-unauthenticated \
  --port 8080

# 5. Acessar URL fornecida
```

---

### 6️⃣ **Docker Compose (VPS/Servidor Próprio)**

Se você tem um servidor (DigitalOcean, AWS EC2, Linode, etc):

```bash
# 1. SSH no servidor
ssh usuario@seu-servidor.com

# 2. Clonar repositório
git clone seu-repositorio.git
cd goLang

# 3. Iniciar com Docker Compose
docker-compose up -d

# 4. Verificar
docker-compose ps
docker-compose logs -f

# 5. Configurar nginx como reverse proxy (opcional)
# Ver seção "Nginx Reverse Proxy" abaixo
```

---

### 7️⃣ **Systemd (Linux VPS sem Docker)**

Para servidores Linux tradicionais:

```bash
# 1. SSH no servidor
ssh usuario@seu-servidor.com

# 2. Instalar Go
wget https://go.dev/dl/go1.21.5.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin

# 3. Clonar e compilar
git clone seu-repositorio.git
cd goLang
go build -o api-server api.go

# 4. Criar diretório de produção
sudo mkdir -p /opt/github-api
sudo cp api-server /opt/github-api/

# 5. Instalar service
sudo cp github-api.service /etc/systemd/system/
sudo systemctl daemon-reload
sudo systemctl enable github-api
sudo systemctl start github-api

# 6. Verificar status
sudo systemctl status github-api
sudo journalctl -u github-api -f
```

---

### 8️⃣ **AWS (Amazon Web Services)**

#### Opção A: AWS Elastic Beanstalk
```bash
# 1. Instalar EB CLI
pip install awsebcli

# 2. Inicializar
eb init -p docker github-api-app

# 3. Criar ambiente e deploy
eb create production-env
eb open
```

#### Opção B: AWS ECS (Container Service)
1. Push da imagem para ECR
2. Criar task definition
3. Criar service no ECS
4. Configurar load balancer

---

## 🔧 Configurações Adicionais

### Nginx Reverse Proxy

```nginx
# /etc/nginx/sites-available/github-api
server {
    listen 80;
    server_name seu-dominio.com;

    location / {
        proxy_pass http://localhost:8080;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection 'upgrade';
        proxy_set_header Host $host;
        proxy_cache_bypass $http_upgrade;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    }
}
```

```bash
# Habilitar configuração
sudo ln -s /etc/nginx/sites-available/github-api /etc/nginx/sites-enabled/
sudo nginx -t
sudo systemctl reload nginx

# SSL com Let's Encrypt
sudo apt install certbot python3-certbot-nginx
sudo certbot --nginx -d seu-dominio.com
```

---

## 📊 Monitoramento e Logs

### Railway
```bash
railway logs
```

### Render
- Dashboard web com logs em tempo real

### Fly.io
```bash
fly logs
fly status
```

### Docker Compose
```bash
docker-compose logs -f github-api
```

### Systemd
```bash
sudo journalctl -u github-api -f
```

---

## 🔐 Variáveis de Ambiente

Para produção, você pode adicionar:

```bash
# .env (NÃO commitar no Git)
PORT=8080
GITHUB_TOKEN=seu-token-opcional
ENV=production
```

Modifique `api.go` para usar:
```go
import "github.com/joho/godotenv"

func init() {
    godotenv.Load()
}
```

---

## 🎯 Recomendação Final

Para **começar rapidamente**:
1. **Railway.app** - Melhor para protótipos e MVPs
2. **Render.com** - Ótima interface, fácil de usar
3. **Fly.io** - Performance excelente, global

Para **produção séria**:
1. **Google Cloud Run** - Escala automática, pay-per-use
2. **AWS ECS/Fargate** - Mais controle, integração AWS
3. **VPS próprio** - Máximo controle e customização

---

## ✅ Checklist de Deploy

- [ ] Código commitado no Git
- [ ] Dockerfile testado localmente
- [ ] Variáveis de ambiente configuradas
- [ ] Plataforma escolhida
- [ ] Deploy realizado
- [ ] URL acessível
- [ ] Healthcheck funcionando
- [ ] Logs verificados
- [ ] Monitoramento configurado
- [ ] Domínio customizado (opcional)
- [ ] SSL/HTTPS ativado

---

## 🆘 Problemas Comuns

### Porta incorreta
- Certifique-se que a variável `PORT` está configurada
- Plataformas cloud geralmente fornecem `PORT` automaticamente

### Timeout
- Aumente o timeout no código (já configurado para 15s)
- Configure healthcheck adequadamente

### Build falha
- Verifique se `go.mod` está presente
- Certifique-se que o Dockerfile está correto

---

## 📞 Suporte

Para ajuda adicional:
- Railway: https://railway.app/help
- Render: https://render.com/docs
- Fly.io: https://fly.io/docs
- GCP: https://cloud.google.com/run/docs

Boa sorte com seu deploy! 🚀
