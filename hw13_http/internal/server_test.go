package internal

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/tmpcaot/HOME_WORK_BASIC/hw13_http/pkg/server"
)

// TestHandleRequest проверяет обработку запросов методом GET.
func TestHandleRequest_GET(t *testing.T) {
	// Создаем тестовый сервер
	req, err := http.NewRequestWithContext(context.Background(), "GET", "/api/example", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.HandleRequest)

	// Выполняем запрос.
	handler.ServeHTTP(rr, req)

	// Проверяем статус-код.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Проверяем тело ответа.
	expected := `{"status":200,"message":"Received GET request to /api/example"}`
	actual := rr.Body.String()
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

// TestHandleRequest проверяет обработку запросов методом POST.
func TestHandleRequest_POST(t *testing.T) {
	// Создаем тестовый сервер.
	req, err := http.NewRequestWithContext(context.Background(), "POST", "/api/example",
		bytes.NewBuffer([]byte(`{"message":"Hello!"}`)))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.HandleRequest)

	// Выполняем запрос.
	handler.ServeHTTP(rr, req)

	// Проверяем статус-код.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Проверяем тело ответа.
	expected := `{"status":200,"message":"Received POST request to /api/example"}`
	actual := rr.Body.String()
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

// TestRunServer проверяет запуск сервера.
func TestRunServer(t *testing.T) {
	addr := ":8080"
	go func() {
		server.RunServer(addr)
	}()

	// Ждем немного времени, чтобы сервер успел запуститься.
	time.Sleep(time.Millisecond * 100)

	// Создаем контекст.
	ctx := context.Background()

	// Проверяем доступность сервера.
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("http://%s/api/example", addr), nil)
	if err != nil {
		t.Errorf("error creating request: %v", err)
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("error getting response: %v", err)
	}
	defer res.Body.Close()

	// Проверяем статус-код.
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, res.StatusCode)
	}

	// Читаем тело ответа.
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("error reading response body: %v", err)
	}

	// Проверяем содержимое тела ответа.
	expected := `{"status":200,"message":"Received GET request to /api/example"}`
	if !reflect.DeepEqual(expected, string(body)) {
		t.Errorf("unexpected response body: got %q, want %q", string(body), expected)
	}
}
