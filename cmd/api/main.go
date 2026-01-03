package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github-api-demo/internal/clients"
	"github-api-demo/internal/handlers"
	"github-api-demo/internal/router"
	"github-api-demo/internal/services"
)

func main() {
	// Inicializar clientes
	tvmazeClient := clients.NewTVMazeClient()
	githubClient := clients.NewGitHubClient()
	
	// Inicializar serviços
	tvmazeService := services.NewTVMazeService(tvmazeClient)
	githubService := services.NewGitHubService(githubClient)
	
	// Inicializar handlers
	tvmazeHandler := handlers.NewTVMazeHandler(tvmazeService)
	githubHandler := handlers.NewGitHubHandler(githubService)
	
	// Configurar rotas
	mux := router.Setup(tvmazeHandler, githubHandler)
	
	// Configurar servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	
	// Canal para capturar sinais de shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	
	// Iniciar servidor em goroutine
	go func() {
		log.Printf("🚀 Servidor iniciado na porta %s", port)
		log.Printf("📚 Documentação: http://localhost:%s/docs", port)
		log.Printf("📡 API TVMaze: http://localhost:%s/schedule", port)
		log.Printf("🐙 API GitHub: http://localhost:%s/api/user?username=patrickbathu", port)
		
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("❌ Erro ao iniciar servidor: %v", err)
		}
	}()
	
	// Aguardar sinal de shutdown
	<-quit
	log.Println("🛑 Shutdown solicitado...")
	
	// Graceful shutdown com timeout de 30 segundos
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("❌ Erro durante shutdown: %v", err)
	}
	
	log.Println("✅ Servidor encerrado com sucesso")
}
