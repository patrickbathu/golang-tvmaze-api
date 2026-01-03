package handlers

import (
	"encoding/json"
	"net/http"

	"github-api-demo/internal/models"
	"github-api-demo/internal/services"
)

// GitHubHandler contém os handlers para GitHub
type GitHubHandler struct {
	service *services.GitHubService
}

// NewGitHubHandler cria uma nova instância do handler
func NewGitHubHandler(service *services.GitHubService) *GitHubHandler {
	return &GitHubHandler{
		service: service,
	}
}

// Home retorna informações sobre a API do GitHub
func (h *GitHubHandler) Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	info := map[string]interface{}{
		"message": "🐙 GitHub User API",
		"version": "2.0.0",
		"endpoints": map[string]string{
			"GET /api/user?username=USER": "Buscar informações de usuário do GitHub",
		},
		"example": "/api/user?username=patrickbathu",
	}
	
	json.NewEncoder(w).Encode(info)
}

// GetUser retorna informações de um usuário do GitHub
func (h *GitHubHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	
	username := r.URL.Query().Get("username")
	if username == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Success: false,
			Error:   "Parâmetro 'username' é obrigatório. Use: /api/user?username=USERNAME",
		})
		return
	}
	
	user, err := h.service.GetUser(username)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "usuário não encontrado" {
			statusCode = http.StatusNotFound
		}
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(models.Response{
			Success: false,
			Error:   err.Error(),
		})
		return
	}
	
	json.NewEncoder(w).Encode(models.Response{
		Success: true,
		Data:    user,
	})
}
