package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestHomeHandler testa o endpoint raiz
func TestHomeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(homeHandler)

	handler.ServeHTTP(rr, req)

	// Verificar status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler retornou código errado: got %v want %v",
			status, http.StatusOK)
	}

	// Verificar Content-Type
	expected := "application/json"
	if ctype := rr.Header().Get("Content-Type"); ctype != expected {
		t.Errorf("Content-Type incorreto: got %v want %v",
			ctype, expected)
	}

	// Verificar se é JSON válido
	var response map[string]interface{}
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf("Resposta não é JSON válido: %v", err)
	}

	// Verificar se tem a chave "message"
	if _, ok := response["message"]; !ok {
		t.Error("Resposta não contém a chave 'message'")
	}
}

// TestUserHandler_MissingUsername testa validação de parâmetro
func TestUserHandler_MissingUsername(t *testing.T) {
	req, err := http.NewRequest("GET", "/user", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler)

	handler.ServeHTTP(rr, req)

	// Deve retornar Bad Request
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler retornou código errado: got %v want %v",
			status, http.StatusBadRequest)
	}

	// Verificar resposta de erro
	var response Response
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf("Resposta não é JSON válido: %v", err)
	}

	if response.Success {
		t.Error("Success deveria ser false para requisição inválida")
	}

	if response.Error == "" {
		t.Error("Error message está vazio")
	}
}

// TestUserHandler_ValidUser testa busca de usuário válido
func TestUserHandler_ValidUser(t *testing.T) {
	req, err := http.NewRequest("GET", "/user?username=torvalds", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler)

	handler.ServeHTTP(rr, req)

	// Verificar status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler retornou código errado: got %v want %v",
			status, http.StatusOK)
	}

	// Verificar resposta
	var response Response
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf("Resposta não é JSON válido: %v", err)
	}

	if !response.Success {
		t.Errorf("Success deveria ser true, erro: %v", response.Error)
	}

	if response.Data == nil {
		t.Error("Data não deveria ser nil")
	}
}

// TestUserHandler_InvalidUser testa busca de usuário inválido
func TestUserHandler_InvalidUser(t *testing.T) {
	req, err := http.NewRequest("GET", "/user?username=usuarioquenaoexiste123456789", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler)

	handler.ServeHTTP(rr, req)

	// Verificar que não é 200
	if status := rr.Code; status == http.StatusOK {
		t.Error("handler deveria retornar erro para usuário inexistente")
	}

	// Verificar resposta de erro
	var response Response
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf("Resposta não é JSON válido: %v", err)
	}

	if response.Success {
		t.Error("Success deveria ser false para usuário inexistente")
	}
}

// TestGetGitHubUser testa a função de busca de usuário
func TestGetGitHubUser(t *testing.T) {
	// Teste com usuário válido
	user, err := getGitHubUser("torvalds")
	if err != nil {
		t.Errorf("Erro ao buscar usuário válido: %v", err)
	}

	if user == nil {
		t.Fatal("User não deveria ser nil")
	}

	if user.Login != "torvalds" {
		t.Errorf("Login incorreto: got %v want torvalds", user.Login)
	}

	if user.Name == "" {
		t.Error("Name não deveria estar vazio")
	}

	// Teste com usuário inválido
	_, err = getGitHubUser("usuarioquenaoexiste123456789")
	if err == nil {
		t.Error("Deveria retornar erro para usuário inexistente")
	}
}

// Benchmark para testar performance
func BenchmarkUserHandler(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user?username=torvalds", nil)
	
	for i := 0; i < b.N; i++ {
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(userHandler)
		handler.ServeHTTP(rr, req)
	}
}
