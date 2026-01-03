#!/bin/bash

# Script de deploy para diferentes plataformas cloud

echo "🚀 Deploy Script - GitHub API Microservice"
echo ""

# Função para deploy no Heroku
deploy_heroku() {
    echo "📦 Deploy no Heroku..."
    heroku create github-api-go-service
    git push heroku main
    heroku open
}

# Função para deploy no Railway
deploy_railway() {
    echo "🚂 Deploy no Railway..."
    railway init
    railway up
    railway open
}

# Função para deploy no Fly.io
deploy_fly() {
    echo "🪰 Deploy no Fly.io..."
    fly launch --name github-api-service
    fly deploy
    fly open
}

# Função para deploy no Render
deploy_render() {
    echo "🎨 Deploy no Render..."
    echo "1. Conecte seu repositório em https://render.com"
    echo "2. Selecione 'New Web Service'"
    echo "3. Use o arquivo render.yaml para configuração automática"
}

# Função para deploy no Google Cloud Run
deploy_gcloud() {
    echo "☁️  Deploy no Google Cloud Run..."
    gcloud builds submit --tag gcr.io/PROJECT_ID/github-api
    gcloud run deploy github-api \
        --image gcr.io/PROJECT_ID/github-api \
        --platform managed \
        --region us-central1 \
        --allow-unauthenticated
}

# Menu
echo "Escolha a plataforma de deploy:"
echo "1) Heroku"
echo "2) Railway"
echo "3) Fly.io"
echo "4) Render"
echo "5) Google Cloud Run"
echo "6) Cancelar"
read -p "Opção: " choice

case $choice in
    1) deploy_heroku ;;
    2) deploy_railway ;;
    3) deploy_fly ;;
    4) deploy_render ;;
    5) deploy_gcloud ;;
    6) echo "Cancelado." ;;
    *) echo "Opção inválida" ;;
esac
