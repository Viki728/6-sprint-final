package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

// Создаем структуру сервера
type Server struct {
	Logger *log.Logger
	HTTP   *http.Server
}

// Функция создания сервера
func NewServer(logger *log.Logger) *Server {
	//Роутер
	mux := http.NewServeMux()

	//Регистрируем хендлеры
	mux.HandleFunc("/", handlers.HandlerIndex)
	mux.HandleFunc("/upload", handlers.HandlerUpload)

	//Создаем экземпляр стуктуры http.Server
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	//Возвращаем ссылку на сервер
	return &Server{
		Logger: logger,
		HTTP:   srv,
	}
}
