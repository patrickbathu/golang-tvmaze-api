package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github-api-demo/internal/models"
	"github-api-demo/internal/services"
)

// TVMazeHandler contém os handlers para TVMaze
type TVMazeHandler struct {
	service *services.TVMazeService
}

// NewTVMazeHandler cria uma nova instância do handler
func NewTVMazeHandler(service *services.TVMazeService) *TVMazeHandler {
	return &TVMazeHandler{
		service: service,
	}
}

// Home retorna informações sobre a API
func (h *TVMazeHandler) Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	info := map[string]interface{}{
		"message": "📺 API Go - TVMaze Schedule",
		"version": "3.0.0",
		"date":    time.Now().Format("2006-01-02"),
		"time":    time.Now().Format("15:04"),
		"docs":    "/docs - 📚 Documentação Interativa",
		"endpoints": map[string]string{
			"GET /":                      "Informações da API",
			"GET /docs":                  "📚 Documentação Interativa (Swagger-like)",
			"GET /schedule":              "Programação de hoje (país padrão: US)",
			"GET /schedule?country=BR":   "Programação de hoje no Brasil",
			"GET /search?q=NOME":         "Buscar shows por nome",
			"GET /show?id=ID":            "Detalhes de um show específico",
			"GET /genre?genre=GENERO":    "Programação filtrada por gênero/categoria",
			"GET /now":                   "O que está passando agora",
			"GET /api/user?username=USER": "Informações de usuário do GitHub",
		},
		"examples": []string{
			"/docs",
			"/schedule",
			"/schedule?country=BR",
			"/search?q=friends",
			"/show?id=431",
			"/genre?genre=Sports&country=US",
			"/genre?genre=Drama&country=BR",
			"/now?country=US",
			"/api/user?username=patrickbathu",
		},
		"genres": []string{
			"Sports", "Drama", "Comedy", "Action", "Thriller",
			"Horror", "Romance", "Science-Fiction", "Fantasy",
			"Mystery", "Crime", "Documentary", "News",
		},
	}
	
	json.NewEncoder(w).Encode(info)
}

// Schedule retorna a programação de hoje
func (h *TVMazeHandler) Schedule(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	
	country := r.URL.Query().Get("country")
	if country == "" {
		country = "US"
	}
	
	schedule, err := h.service.GetTodaySchedule(country)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.Response{
			Success: false,
			Error:   err.Error(),
		})
		return
	}
	
	json.NewEncoder(w).Encode(models.Response{
		Success: true,
		Data:    schedule,
		Count:   len(schedule),
	})
}

// Search busca shows
func (h *TVMazeHandler) Search(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	
	query := r.URL.Query().Get("q")
	if query == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Success: false,
			Error:   "Parâmetro 'q' é obrigatório. Use: /search?q=NOME",
		})
		return
	}
	
	results, err := h.service.SearchShows(query)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.Response{
			Success: false,
			Error:   err.Error(),
		})
		return
	}
	
	json.NewEncoder(w).Encode(models.Response{
		Success: true,
		Data:    results,
		Count:   len(results),
	})
}

// ShowDetails retorna detalhes de um show
func (h *TVMazeHandler) ShowDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	
	id := r.URL.Query().Get("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Success: false,
			Error:   "Parâmetro 'id' é obrigatório. Use: /show?id=123",
		})
		return
	}
	
	show, err := h.service.GetShowByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(models.Response{
			Success: false,
			Error:   err.Error(),
		})
		return
	}
	
	json.NewEncoder(w).Encode(models.Response{
		Success: true,
		Data:    show,
	})
}

// Genre retorna programação por gênero
func (h *TVMazeHandler) Genre(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	
	genre := r.URL.Query().Get("genre")
	if genre == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Success: false,
			Error:   "Parâmetro 'genre' é obrigatório. Use: /genre?genre=Sports&country=US",
		})
		return
	}
	
	country := r.URL.Query().Get("country")
	if country == "" {
		country = "US"
	}
	
	schedule, err := h.service.GetScheduleByGenre(country, genre)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.Response{
			Success: false,
			Error:   err.Error(),
		})
		return
	}
	
	json.NewEncoder(w).Encode(models.Response{
		Success: true,
		Data:    schedule,
		Count:   len(schedule),
	})
}

// NowPlaying retorna o que está passando agora
func (h *TVMazeHandler) NowPlaying(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	
	country := r.URL.Query().Get("country")
	if country == "" {
		country = "US"
	}
	
	nowPlaying, err := h.service.GetNowPlaying(country)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.Response{
			Success: false,
			Error:   err.Error(),
		})
		return
	}
	
	response := map[string]interface{}{
		"success":      true,
		"current_time": time.Now().Format("15:04"),
		"country":      country,
		"data":         nowPlaying,
		"count":        len(nowPlaying),
	}
	
	json.NewEncoder(w).Encode(response)
}
