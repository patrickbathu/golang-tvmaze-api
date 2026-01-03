package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHomeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(homeHandler)
	handler.ServeHTTP(rr, req)
	
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler retornou código errado: got %v want %v", status, http.StatusOK)
	}
	
	var response map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("erro ao fazer parse do JSON: %v", err)
	}
	
	if message, ok := response["message"].(string); !ok || message == "" {
		t.Errorf("resposta deve conter uma mensagem")
	}
}

func TestScheduleHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/schedule?country=US", nil)
	if err != nil {
		t.Fatal(err)
	}
	
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(scheduleHandler)
	handler.ServeHTTP(rr, req)
	
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler retornou código errado: got %v want %v", status, http.StatusOK)
	}
	
	var response Response
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("erro ao fazer parse do JSON: %v", err)
	}
	
	if !response.Success {
		t.Errorf("resposta deveria ser success=true, got: %v", response.Success)
	}
}

func TestSearchHandlerMissingQuery(t *testing.T) {
	req, err := http.NewRequest("GET", "/search", nil)
	if err != nil {
		t.Fatal(err)
	}
	
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(searchHandler)
	handler.ServeHTTP(rr, req)
	
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler retornou código errado: got %v want %v", status, http.StatusBadRequest)
	}
	
	var response Response
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("erro ao fazer parse do JSON: %v", err)
	}
	
	if response.Success {
		t.Errorf("resposta deveria ser success=false para query vazia")
	}
}

func TestSearchHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/search?q=friends", nil)
	if err != nil {
		t.Fatal(err)
	}
	
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(searchHandler)
	handler.ServeHTTP(rr, req)
	
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler retornou código errado: got %v want %v", status, http.StatusOK)
	}
	
	var response Response
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("erro ao fazer parse do JSON: %v", err)
	}
	
	if !response.Success {
		t.Errorf("resposta deveria ser success=true, got: %v", response.Success)
	}
}

func TestShowHandlerMissingID(t *testing.T) {
	req, err := http.NewRequest("GET", "/show", nil)
	if err != nil {
		t.Fatal(err)
	}
	
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(showHandler)
	handler.ServeHTTP(rr, req)
	
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler retornou código errado: got %v want %v", status, http.StatusBadRequest)
	}
}

func TestShowHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/show?id=431", nil)
	if err != nil {
		t.Fatal(err)
	}
	
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(showHandler)
	handler.ServeHTTP(rr, req)
	
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler retornou código errado: got %v want %v", status, http.StatusOK)
	}
	
	var response Response
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("erro ao fazer parse do JSON: %v", err)
	}
	
	if !response.Success {
		t.Errorf("resposta deveria ser success=true, got: %v", response.Success)
	}
}

func TestGetTodaySchedule(t *testing.T) {
	schedule, err := getTodaySchedule("US")
	if err != nil {
		t.Errorf("getTodaySchedule retornou erro: %v", err)
	}
	
	if schedule == nil {
		t.Errorf("schedule não deveria ser nil")
	}
}

func TestLoggingMiddleware(t *testing.T) {
	handler := loggingMiddleware(homeHandler)
	
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("middleware retornou código errado: got %v want %v", status, http.StatusOK)
	}
}

func TestHTTPClient(t *testing.T) {
	if httpClient == nil {
		t.Errorf("httpClient não deveria ser nil")
	}
	
	if httpClient.Timeout != 15*time.Second {
		t.Errorf("timeout esperado: 15s, got: %v", httpClient.Timeout)
	}
}
