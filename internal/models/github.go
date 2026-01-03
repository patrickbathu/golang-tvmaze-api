package models

// GitHubUser representa os dados do usuário do GitHub
type GitHubUser struct {
	Login       string `json:"login"`
	Name        string `json:"name"`
	Bio         string `json:"bio"`
	Location    string `json:"location"`
	Followers   int    `json:"followers"`
	Following   int    `json:"following"`
	PublicRepos int    `json:"public_repos"`
	AvatarURL   string `json:"avatar_url"`
	CreatedAt   string `json:"created_at"`
}
