package middleware

import (
	"log"
	"net/http"
	"time"
)

// Logging middleware para registrar requisições HTTP
func Logging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("📨 %s %s", r.Method, r.URL.Path)
		next(w, r)
		log.Printf("✅ %s %s - %v", r.Method, r.URL.Path, time.Since(start))
	}
}

// CORS middleware para configurar headers de CORS
func CORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		next(w, r)
	}
}
