package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// HandlerIndex возвращает index.html
func HandlerIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../index.html")
}

// HandlerUpload принимает файл, обрабатывает и возвращает результат
func HandlerUpload(w http.ResponseWriter, r *http.Request) {

	//Ограничиваем метод
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return

	}
	//Парсим форму и устанавливаем лимит для файла (10 МБ)
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Ошибка при парсинге формы", http.StatusInternalServerError)
		return
	}

	//Получаем файл
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Ошибка при получении файла", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	//Читаем содержимое файла
	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Ошибка при чтении файла", http.StatusInternalServerError)
		return
	}

	//Конвертируем данные при помощи пакета service
	result, err := service.Conv(string(data))
	if err != nil {
		http.Error(w, "Ошибка при обработке данных: "+err.Error(), http.StatusInternalServerError)
		return
	}

	//Создаем имя выходного файла
	ext := filepath.Ext(handler.Filename) // расширение исходного файла
	outputName := fmt.Sprintf("output_%d%s", time.Now().Unix(), ext)

	//Создаем локальный файл
	outFile, err := os.Create(outputName)
	if err != nil {
		http.Error(w, "Ошибка при создании файла", http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	//Записываем результат конвертации(из морзе в текст или наоборот) в файл
	_, err = outFile.WriteString(result)
	if err != nil {
		http.Error(w, "Ошибка при записи файла", http.StatusInternalServerError)
		return
	}

	//Возвращаем результат пользователю
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(result))
}
