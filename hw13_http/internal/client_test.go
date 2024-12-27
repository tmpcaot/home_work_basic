package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/tmpcaot/HOME_WORK_BASIC/hw13_http/pkg/client"
)

// TestClient_GET проверяет работу метода GET.
func TestClient_GET(t *testing.T) {
	// Создаем тестовый сервер.
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte(`{"status": 200, "message": "Test GET"}`))
	}))
	defer server.Close() // Закрываем сервер после завершения теста.

	url := server.URL + "/api/example"
	method := "GET"

	// Создаем новый контекст.
	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		t.Fatalf("Ошибка создания запроса: %v\n", err)
	}

	client := &http.Client{}    // Используем стандартный HTTP клиент.
	resp, err := client.Do(req) // Отправляем запрос.
	if err != nil {
		t.Fatalf("Ошибка отправки запроса: %v\n", err)
	}
	defer resp.Body.Close() // Закрываем тело ответа после его обработки.

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Ошибка чтения тела ответа: %v\n", err)
	}

	expectedBody := `{"status": 200, "message": "Test GET"}`
	if strings.TrimSpace(string(body)) != expectedBody { // Убираем лишние пробелы и сравниваем.
		t.Errorf("Ожидалось: %s, получено: %s", expectedBody, string(body))
	}
}

// TestClient_POST проверяет работу метода POST.
func TestClient_POST(t *testing.T) {
	// Создаем тестовый сервер.
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte(`{"status": 201, "message": "Test POST"}`))
	}))
	defer server.Close() // Закрываем сервер после завершения теста.

	url := server.URL + "/api/example"
	method := "POST"

	postData := &client.PostData{
		Message: "Hello from test client!",
	}
	jsonData, err := json.Marshal(postData)
	if err != nil {
		t.Fatalf("Ошибка маршалинга данных: %v\n", err)
	}

	// Создаем новый контекст.
	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Ошибка создания запроса: %v\n", err)
	}
	req.Header.Set("Content-Type", "application/json") // Устанавливаем заголовок Content-Type.

	client := &http.Client{}    // Используем стандартный HTTP клиент.
	resp, err := client.Do(req) // Отправляем запрос.
	if err != nil {
		t.Fatalf("Ошибка отправки запроса: %v\n", err)
	}
	defer resp.Body.Close() // Закрываем тело ответа после его обработки.

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Ошибка чтения тела ответа: %v\n", err)
	}

	expectedBody := `{"status": 201, "message": "Test POST"}`
	if strings.TrimSpace(string(body)) != expectedBody { // Убираем лишние пробелы и сравниваем.
		t.Errorf("Ожидалось: %s, получено: %s", expectedBody, string(body))
	}
}
