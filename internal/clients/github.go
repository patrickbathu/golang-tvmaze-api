package clients

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github-api-demo/internal/models"
)

// GitHubClient é o cliente para a API do GitHub
type GitHubClient struct {
	httpClient *http.Client
	baseURL    string
}

// NewGitHubClient cria uma nova instância do cliente GitHub
func NewGitHubClient() *GitHubClient {
	return &GitHubClient{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		baseURL: "https://api.github.com",
	}
}

// GetUser busca dados de um usuário do GitHub
func (c *GitHubClient) GetUser(username string) (*models.GitHubUser, error) {
	url := fmt.Sprintf("%s/users/%s", c.baseURL, username)
	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar requisição: %w", err)
	}
	
	req.Header.Set("User-Agent", "GoLang-TVMaze-API")
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer requisição: %w", err)
	}
	defer resp.Body.Close()
	
	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("usuário não encontrado")
	}
	
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %d", resp.StatusCode)
	}
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler resposta: %w", err)
	}
	
	var user models.GitHubUser
	if err := json.Unmarshal(body, &user); err != nil {
		return nil, fmt.Errorf("erro ao decodificar JSON: %w", err)
	}
	
	return &user, nil
}
