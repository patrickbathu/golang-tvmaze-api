package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// GitHubUser representa os dados do usuário do GitHub
type GitHubUser struct {
	Login     string `json:"login"`
	Name      string `json:"name"`
	Bio       string `json:"bio"`
	Location  string `json:"location"`
	Followers int    `json:"followers"`
	Following int    `json:"following"`
	PublicRepos int  `json:"public_repos"`
	AvatarURL string `json:"avatar_url"`
	CreatedAt string `json:"created_at"`
}

// Resposta da nossa API
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// Função para buscar dados do usuário do GitHub
func getGitHubUser(username string) (*GitHubUser, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s", username)
	
	// Criar cliente HTTP com timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	
	// Fazer requisição GET
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	
	// Adicionar header User-Agent (GitHub requer isso)
	req.Header.Set("User-Agent", "GoLang-API-Tutorial")
	
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	// Verificar status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("usuário não encontrado ou erro na API do GitHub")
	}
	
	// Ler corpo da resposta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	
	// Fazer parse do JSON
	var user GitHubUser
	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, err
	}
	
	return &user, nil
}

// Handler para a rota /user/:username
func userHandler(w http.ResponseWriter, r *http.Request) {
	// Configurar CORS e Content-Type
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	
	// Extrair username da URL
	username := r.URL.Query().Get("username")
	if username == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Success: false,
			Error:   "Parâmetro 'username' é obrigatório. Use: /user?username=NOME",
		})
		return
	}
	
	// Buscar dados do GitHub
	user, err := getGitHubUser(username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{
			Success: false,
			Error:   err.Error(),
		})
		return
	}
	
	// Retornar sucesso
	json.NewEncoder(w).Encode(Response{
		Success: true,
		Data:    user,
	})
}

// Handler para a rota raiz
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	info := map[string]interface{}{
		"message": "🚀 API Go - Consulta GitHub Users",
		"endpoints": map[string]string{
			"/":              "Informações da API",
			"/user?username=": "Buscar informações de usuário do GitHub",
		},
		"exemplo": "http://localhost:8080/user?username=torvalds",
	}
	
	json.NewEncoder(w).Encode(info)
}

func main() {
	// Obter porta da variável de ambiente ou usar padrão
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	// Registrar rotas
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/user", userHandler)
	
	// Configurar servidor
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	
	// Canal para escutar sinais do sistema
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	
	// Iniciar servidor em goroutine
	go func() {
		log.Printf("🚀 Servidor rodando em http://localhost:%s\n", port)
		log.Println("📚 Endpoints disponíveis:")
		log.Println("   GET / - Informações da API")
		log.Println("   GET /user?username=NOME - Buscar usuário do GitHub")
		log.Printf("\n💡 Exemplo: http://localhost:%s/user?username=torvalds\n", port)
		log.Println("\n⏹  Pressione Ctrl+C para parar o servidor")
		
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("❌ Erro ao iniciar servidor: %v\n", err)
		}
	}()
	
	// Aguardar sinal de interrupção
	<-stop
	
	// Graceful shutdown
	log.Println("\n⏳ Desligando servidor gracefully...")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("❌ Erro durante shutdown: %v\n", err)
	} else {
		log.Println("✅ Servidor encerrado com sucesso!")
	}
}
