package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// Conv определяет формат и возвращает конвертированный результат
func Conv(input string) (string, error) {
	//убираем лишние пробелы вначале и в конце строки
	input = strings.TrimSpace(input)

	if input == "" {
		return "", errors.New("пустая строка")
	}

	//Конвертируем строку
	//если только точки, тире и пробелы
	if isMorse(input) {

		//переводим морзе в текст
		return morse.ToText(input), nil
	}

	//иначе переводим текст в морзе
	return morse.ToMorse(input), nil
}

// Проверяем морзе ли это
func isMorse(s string) bool {
	for _, r := range s {
		//возврвщвем false в том случае если символ НЕ точка, НЕ тире, и НЕ пробел
		if r != '.' && r != '-' && r != ' ' {
			return false
		}
	}
	return true
}
