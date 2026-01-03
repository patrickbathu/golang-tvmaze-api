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

// Episode representa um episódio na TVMaze API
type Episode struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Season   int    `json:"season"`
	Number   int    `json:"number"`
	Airdate  string `json:"airdate"`
	Airtime  string `json:"airtime"`
	Runtime  int    `json:"runtime"`
	Summary  string `json:"summary"`
	Image    *Image `json:"image"`
}

// Show representa um show de TV
type Show struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Type     string   `json:"type"`
	Language string   `json:"language"`
	Genres   []string `json:"genres"`
	Status   string   `json:"status"`
	Premiered string  `json:"premiered"`
	Summary  string   `json:"summary"`
	Image    *Image   `json:"image"`
	Network  *Network `json:"network"`
}

// Network representa a rede de TV
type Network struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Country Country `json:"country"`
}

// Country representa o país
type Country struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

// Image representa as imagens
type Image struct {
	Medium   string `json:"medium"`
	Original string `json:"original"`
}

// Schedule representa um item da programação
type Schedule struct {
	ID       int      `json:"id"`
	Airdate  string   `json:"airdate"`
	Airtime  string   `json:"airtime"`
	Show     Show     `json:"show"`
	Episode  *Episode `json:"episode,omitempty"`
}

// Response representa a resposta padrão da API
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
	Count   int         `json:"count,omitempty"`
}

// Cliente HTTP global
var httpClient = &http.Client{
	Timeout: 15 * time.Second,
}

// getTodaySchedule busca a programação de hoje na TVMaze
func getTodaySchedule(country string) ([]Schedule, error) {
	// Formato de data: YYYY-MM-DD
	today := time.Now().Format("2006-01-02")
	
	url := fmt.Sprintf("https://api.tvmaze.com/schedule?country=%s&date=%s", country, today)
	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	
	req.Header.Set("User-Agent", "GoLang-TVMaze-API")
	
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erro na API TVMaze: status %d", resp.StatusCode)
	}
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	
	var schedule []Schedule
	err = json.Unmarshal(body, &schedule)
	if err != nil {
		return nil, err
	}
	
	return schedule, nil
}

// searchShow busca shows pelo nome
func searchShow(query string) ([]map[string]interface{}, error) {
	url := fmt.Sprintf("https://api.tvmaze.com/search/shows?q=%s", query)
	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	
	req.Header.Set("User-Agent", "GoLang-TVMaze-API")
	
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erro na API TVMaze: status %d", resp.StatusCode)
	}
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	
	var results []map[string]interface{}
	err = json.Unmarshal(body, &results)
	if err != nil {
		return nil, err
	}
	
	return results, nil
}

// getShowByID busca informações de um show pelo ID
func getShowByID(id string) (*Show, error) {
	url := fmt.Sprintf("https://api.tvmaze.com/shows/%s", id)
	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	
	req.Header.Set("User-Agent", "GoLang-TVMaze-API")
	
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("show não encontrado")
	}
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	
	var show Show
	err = json.Unmarshal(body, &show)
	if err != nil {
		return nil, err
	}
	
	return &show, nil
}

// Handler para a rota raiz
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	info := map[string]interface{}{
		"message": "📺 API Go - TVMaze Schedule",
		"version": "1.0.0",
		"date":    time.Now().Format("2006-01-02"),
		"endpoints": map[string]string{
			"GET /":                    "Informações da API",
			"GET /schedule":            "Programação de hoje (país padrão: US)",
			"GET /schedule?country=BR": "Programação de hoje no Brasil",
			"GET /search?q=NOME":       "Buscar shows por nome",
			"GET /show/:id":            "Detalhes de um show específico",
		},
		"examples": []string{
			"http://localhost:8080/schedule",
			"http://localhost:8080/schedule?country=BR",
			"http://localhost:8080/search?q=friends",
			"http://localhost:8080/show/431",
		},
	}
	
	json.NewEncoder(w).Encode(info)
}

// Handler para /schedule
func scheduleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	
	// Obter país da query string (padrão: US)
	country := r.URL.Query().Get("country")
	if country == "" {
		country = "US"
	}
	
	// Buscar programação
	schedule, err := getTodaySchedule(country)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{
			Success: false,
			Error:   err.Error(),
		})
		return
	}
	
	json.NewEncoder(w).Encode(Response{
		Success: true,
		Data:    schedule,
		Count:   len(schedule),
	})
}

// Handler para /search
func searchHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	
	query := r.URL.Query().Get("q")
	if query == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Success: false,
			Error:   "Parâmetro 'q' é obrigatório. Use: /search?q=NOME",
		})
		return
	}
	
	results, err := searchShow(query)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{
			Success: false,
			Error:   err.Error(),
		})
		return
	}
	
	json.NewEncoder(w).Encode(Response{
		Success: true,
		Data:    results,
		Count:   len(results),
	})
}

// Handler para /show/:id
func showHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	
	id := r.URL.Query().Get("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Success: false,
			Error:   "Parâmetro 'id' é obrigatório. Use: /show?id=123",
		})
		return
	}
	
	show, err := getShowByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Response{
			Success: false,
			Error:   err.Error(),
		})
		return
	}
	
	json.NewEncoder(w).Encode(Response{
		Success: true,
		Data:    show,
	})
}

// Middleware de logging
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("📨 %s %s", r.Method, r.URL.Path)
		next(w, r)
		log.Printf("✅ %s %s - %v", r.Method, r.URL.Path, time.Since(start))
	}
}

func main() {
	// Obter porta da variável de ambiente ou usar padrão
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	// Registrar rotas com middleware
	mux := http.NewServeMux()
	mux.HandleFunc("/", loggingMiddleware(homeHandler))
	mux.HandleFunc("/schedule", loggingMiddleware(scheduleHandler))
	mux.HandleFunc("/search", loggingMiddleware(searchHandler))
	mux.HandleFunc("/show", loggingMiddleware(showHandler))
	
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
		log.Printf("📺 Servidor TVMaze API rodando em http://localhost:%s\n", port)
		log.Println("📚 Endpoints disponíveis:")
		log.Println("   GET / - Informações da API")
		log.Println("   GET /schedule - Programação de hoje")
		log.Println("   GET /schedule?country=BR - Programação do Brasil")
		log.Println("   GET /search?q=NOME - Buscar shows")
		log.Println("   GET /show?id=ID - Detalhes do show")
		log.Printf("\n💡 Exemplo: http://localhost:%s/schedule?country=US\n", port)
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
