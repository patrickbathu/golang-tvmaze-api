package router

import (
	"net/http"

	"github-api-demo/internal/handlers"
	"github-api-demo/internal/middleware"
)

// Setup configura todas as rotas da aplicação
func Setup(tvmazeHandler *handlers.TVMazeHandler, githubHandler *handlers.GitHubHandler) *http.ServeMux {
	mux := http.NewServeMux()
	
	// Rotas TVMaze
	mux.HandleFunc("/", middleware.Logging(tvmazeHandler.Home))
	mux.HandleFunc("/docs", middleware.Logging(handlers.DocsHandler))
	mux.HandleFunc("/schedule", middleware.Logging(tvmazeHandler.Schedule))
	mux.HandleFunc("/search", middleware.Logging(tvmazeHandler.Search))
	mux.HandleFunc("/show", middleware.Logging(tvmazeHandler.ShowDetails))
	mux.HandleFunc("/genre", middleware.Logging(tvmazeHandler.Genre))
	mux.HandleFunc("/now", middleware.Logging(tvmazeHandler.NowPlaying))
	
	// Rotas GitHub
	mux.HandleFunc("/api/", middleware.Logging(githubHandler.Home))
	mux.HandleFunc("/api/user", middleware.Logging(githubHandler.GetUser))
	
	return mux
}
