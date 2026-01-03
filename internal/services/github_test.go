package services

import (
	"testing"

	"github-api-demo/internal/clients"
)

func TestNewGitHubService(t *testing.T) {
	client := clients.NewGitHubClient()
	service := NewGitHubService(client)
	
	if service == nil {
		t.Error("NewGitHubService deve retornar uma instância válida")
	}
	
	if service.client == nil {
		t.Error("Cliente não pode ser nil")
	}
}

func TestGetUser_EmptyUsername(t *testing.T) {
	client := clients.NewGitHubClient()
	service := NewGitHubService(client)
	
	_, err := service.GetUser("")
	if err == nil {
		t.Error("GetUser deve retornar erro para username vazio")
	}
	
	if err.Error() != "username não pode ser vazio" {
		t.Errorf("Mensagem de erro incorreta: %v", err)
	}
}
