package services

import (
	"fmt"

	"github-api-demo/internal/clients"
	"github-api-demo/internal/models"
)

// GitHubService contém a lógica de negócio para o GitHub
type GitHubService struct {
	client *clients.GitHubClient
}

// NewGitHubService cria uma nova instância do serviço
func NewGitHubService(client *clients.GitHubClient) *GitHubService {
	return &GitHubService{
		client: client,
	}
}

// GetUser retorna dados de um usuário do GitHub
func (s *GitHubService) GetUser(username string) (*models.GitHubUser, error) {
	if username == "" {
		return nil, fmt.Errorf("username não pode ser vazio")
	}
	return s.client.GetUser(username)
}
