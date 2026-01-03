package services

import (
	"testing"

	"github-api-demo/internal/clients"
)

func TestNewTVMazeService(t *testing.T) {
	client := clients.NewTVMazeClient()
	service := NewTVMazeService(client)
	
	if service == nil {
		t.Error("NewTVMazeService deve retornar uma instância válida")
	}
	
	if service.client == nil {
		t.Error("Cliente não pode ser nil")
	}
}

func TestSearchShows_EmptyQuery(t *testing.T) {
	client := clients.NewTVMazeClient()
	service := NewTVMazeService(client)
	
	_, err := service.SearchShows("")
	if err == nil {
		t.Error("SearchShows deve retornar erro para query vazia")
	}
	
	if err.Error() != "query não pode ser vazia" {
		t.Errorf("Mensagem de erro incorreta: %v", err)
	}
}

func TestGetShowByID_EmptyID(t *testing.T) {
	client := clients.NewTVMazeClient()
	service := NewTVMazeService(client)
	
	_, err := service.GetShowByID("")
	if err == nil {
		t.Error("GetShowByID deve retornar erro para ID vazio")
	}
	
	if err.Error() != "ID não pode ser vazio" {
		t.Errorf("Mensagem de erro incorreta: %v", err)
	}
}

func TestGetScheduleByGenre_EmptyGenre(t *testing.T) {
	client := clients.NewTVMazeClient()
	service := NewTVMazeService(client)
	
	_, err := service.GetScheduleByGenre("US", "")
	if err == nil {
		t.Error("GetScheduleByGenre deve retornar erro para gênero vazio")
	}
	
	if err.Error() != "gênero não pode ser vazio" {
		t.Errorf("Mensagem de erro incorreta: %v", err)
	}
}
