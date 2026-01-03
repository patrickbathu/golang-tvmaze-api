package clients

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github-api-demo/internal/models"
)

// TVMazeClient é o cliente para a API do TVMaze
type TVMazeClient struct {
	httpClient *http.Client
	baseURL    string
}

// NewTVMazeClient cria uma nova instância do cliente TVMaze
func NewTVMazeClient() *TVMazeClient {
	return &TVMazeClient{
		httpClient: &http.Client{
			Timeout: 15 * time.Second,
		},
		baseURL: "https://api.tvmaze.com",
	}
}

// GetSchedule busca a programação de um país e data
func (c *TVMazeClient) GetSchedule(country, date string) ([]models.Schedule, error) {
	url := fmt.Sprintf("%s/schedule?country=%s&date=%s", c.baseURL, country, date)
	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar requisição: %w", err)
	}
	
	req.Header.Set("User-Agent", "GoLang-TVMaze-API")
	
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer requisição: %w", err)
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %d", resp.StatusCode)
	}
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler resposta: %w", err)
	}
	
	var schedule []models.Schedule
	if err := json.Unmarshal(body, &schedule); err != nil {
		return nil, fmt.Errorf("erro ao decodificar JSON: %w", err)
	}
	
	return schedule, nil
}

// SearchShows busca shows pelo nome
func (c *TVMazeClient) SearchShows(query string) ([]map[string]interface{}, error) {
	url := fmt.Sprintf("%s/search/shows?q=%s", c.baseURL, query)
	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar requisição: %w", err)
	}
	
	req.Header.Set("User-Agent", "GoLang-TVMaze-API")
	
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer requisição: %w", err)
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %d", resp.StatusCode)
	}
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler resposta: %w", err)
	}
	
	var results []map[string]interface{}
	if err := json.Unmarshal(body, &results); err != nil {
		return nil, fmt.Errorf("erro ao decodificar JSON: %w", err)
	}
	
	return results, nil
}

// GetShowByID busca um show específico pelo ID
func (c *TVMazeClient) GetShowByID(id string) (*models.Show, error) {
	url := fmt.Sprintf("%s/shows/%s", c.baseURL, id)
	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar requisição: %w", err)
	}
	
	req.Header.Set("User-Agent", "GoLang-TVMaze-API")
	
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer requisição: %w", err)
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %d", resp.StatusCode)
	}
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler resposta: %w", err)
	}
	
	var show models.Show
	if err := json.Unmarshal(body, &show); err != nil {
		return nil, fmt.Errorf("erro ao decodificar JSON: %w", err)
	}
	
	return &show, nil
}
