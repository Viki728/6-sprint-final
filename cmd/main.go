package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	//Создаем логгер
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)

	//Создаем сервер
	s := server.NewServer(logger)
	logger.Println("Сервер запускается на порту 8080")
	err := s.HTTP.ListenAndServe()
	if err != nil {
		logger.Fatal(err)
	}
}
