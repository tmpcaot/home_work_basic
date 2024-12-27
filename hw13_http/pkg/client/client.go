package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// PostData структура данных для отправки через POST-запрос.
type PostData struct {
	Message string `json:"message"`
}

// RunClient запускает клиентское приложение.
func RunClient(serverURL, resourcePath, method string) error {
	var req *http.Request
	var err error

	// Создадим контекст с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // Убедимся, что cancel будет вызван

	switch method {
	case "GET":
		req, err = http.NewRequestWithContext(ctx, http.MethodGet, serverURL+resourcePath, nil)
		if err != nil {
			return fmt.Errorf("error creating request: %w", err)
		}
	case "POST":
		data := &PostData{
			Message: "Hello from client!",
		}
		// Используем просто присваивание для существующей переменной err
		jsonData, marshalErr := json.Marshal(data)
		if marshalErr != nil {
			return fmt.Errorf("error marshaling JSON: %w", marshalErr)
		}
		req, err = http.NewRequestWithContext(ctx, http.MethodPost, serverURL+resourcePath, bytes.NewBuffer(jsonData))
		if err != nil {
			return fmt.Errorf("error creating request: %w", err)
		}
		req.Header.Set("Content-Type", "application/json")
	default:
		return fmt.Errorf("unknown method: %s", method)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}

	fmt.Println(string(body))
	return nil
}
