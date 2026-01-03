package services

import (
	"fmt"
	"strings"
	"time"

	"github-api-demo/internal/clients"
	"github-api-demo/internal/models"
)

// TVMazeService contém a lógica de negócio para o TVMaze
type TVMazeService struct {
	client *clients.TVMazeClient
}

// NewTVMazeService cria uma nova instância do serviço
func NewTVMazeService(client *clients.TVMazeClient) *TVMazeService {
	return &TVMazeService{
		client: client,
	}
}

// GetTodaySchedule retorna a programação de hoje para um país
func (s *TVMazeService) GetTodaySchedule(country string) ([]models.Schedule, error) {
	today := time.Now().Format("2006-01-02")
	return s.client.GetSchedule(country, today)
}

// SearchShows busca shows pelo nome
func (s *TVMazeService) SearchShows(query string) ([]map[string]interface{}, error) {
	if query == "" {
		return nil, fmt.Errorf("query não pode ser vazia")
	}
	return s.client.SearchShows(query)
}

// GetShowByID retorna os detalhes de um show
func (s *TVMazeService) GetShowByID(id string) (*models.Show, error) {
	if id == "" {
		return nil, fmt.Errorf("ID não pode ser vazio")
	}
	return s.client.GetShowByID(id)
}

// GetScheduleByGenre retorna a programação filtrada por gênero
func (s *TVMazeService) GetScheduleByGenre(country, genre string) ([]models.Schedule, error) {
	if genre == "" {
		return nil, fmt.Errorf("gênero não pode ser vazio")
	}

	schedule, err := s.GetTodaySchedule(country)
	if err != nil {
		return nil, err
	}

	// Filtrar por gênero
	var filtered []models.Schedule
	genreLower := strings.ToLower(genre)
	
	for _, item := range schedule {
		for _, g := range item.Show.Genres {
			if strings.ToLower(g) == genreLower || strings.Contains(strings.ToLower(g), genreLower) {
				filtered = append(filtered, item)
				break
			}
		}
	}

	return filtered, nil
}

// GetNowPlaying retorna os programas que estão passando agora
func (s *TVMazeService) GetNowPlaying(country string) ([]models.Schedule, error) {
	schedule, err := s.GetTodaySchedule(country)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	currentTime := now.Format("15:04")
	
	var nowPlaying []models.Schedule
	
	for _, item := range schedule {
		if item.Airtime == "" {
			continue
		}
		
		// Parse do horário do programa
		airtime, err := time.Parse("15:04", item.Airtime)
		if err != nil {
			continue
		}
		
		// Calcular horário de fim (airtime + runtime)
		runtime := 60 // runtime padrão de 60 minutos
		if item.Episode != nil && item.Episode.Runtime > 0 {
			runtime = item.Episode.Runtime
		}
		
		endTime := airtime.Add(time.Duration(runtime) * time.Minute)
		currentTimeParsed, _ := time.Parse("15:04", currentTime)
		
		// Verificar se está no ar agora
		if !currentTimeParsed.Before(airtime) && currentTimeParsed.Before(endTime) {
			nowPlaying = append(nowPlaying, item)
		}
	}
	
	return nowPlaying, nil
}
