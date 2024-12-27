package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

// ResponseData структура данных для ответа.
type ResponseData struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// HandleRequest обработчик запросов.
func HandleRequest(w http.ResponseWriter, r *http.Request) {
	// Выводим информацию о запросе.
	log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL.Path)

	// Формируем данные для ответа.
	response := &ResponseData{
		Status:  http.StatusOK,
		Message: fmt.Sprintf("Received %s request to %s", strings.ToUpper(r.Method), r.URL.Path),
	}

	// Преобразуем данные в JSON.
	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправляем ответ клиенту.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

// RunServer запускает серверное приложение.
func RunServer(addr string) {
	http.HandleFunc("/api/example", HandleRequest)

	server := &http.Server{
		Addr:         addr,
		ReadTimeout:  10 * time.Second, // Устанавливаем таймаут на чтение
		WriteTimeout: 10 * time.Second, // Устанавливаем таймаут на запись
	}

	log.Printf("Starting server on %s...\n", addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Error starting server: %v\n", err)
	}
}
