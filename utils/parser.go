package utils

import (
	"fmt"
	"strconv"
	"strings"
)

// "ID|Название книги|Автор|Год|Доступность"
func ParseBookLine(line string) (id int, title, author string, year int, available bool, err error) {
	line = strings.TrimSpace(line)
	if len(line) == 0 {
		return 0, "", "", 0, false, fmt.Errorf("Пустое значение")
	}

	var data []string = strings.Split(line, "|")

	if len(data) != 5 {
		return 0, "", "", 0, false, fmt.Errorf("Неполный формат - всего 5 полей: id, title, author, year и available")
	}

	// Parse ID
	id, err = strconv.Atoi(strings.TrimSpace(data[0]))
	if err != nil {
		return 0, "", "", 0, false, fmt.Errorf("Неверный формат ID: %w", err)
	}

	// Parse Title
	title = strings.TrimSpace(data[1])
	if len(title) == 0 {
		return 0, "", "", 0, false, fmt.Errorf("Название книги не может быть пустое")
	}

	// Parse Author
	author = strings.TrimSpace(data[2])
	if len(title) == 0 {
		return 0, "", "", 0, false, fmt.Errorf("Имя автора не может быть пустым")
	}

	// Parse Year
	year, err = strconv.Atoi(strings.TrimSpace(data[3]))
	switch {
	case err != nil:
		return 0, "", "", 0, false, fmt.Errorf("Неверный формат ответа для года: %w", err)
	case year <= 1000 || year >= 2026:
		return 0, "", "", 0, false, fmt.Errorf("Неверный диапазон ответа (1000÷2026)")
	}

	// Parse available
	var availableString string = strings.ToLower(strings.TrimSpace(scanner.Text()))

	available, err = StatusStringToBoolean(availableString)
	if err != nil {
		return 0, "", "", 0, false, fmt.Errorf("Ошибка парсинга статуса: %w", err)
	}

	return id, title, author, year, available, nil
}

// "ID|Имя пользователя|Почта"
func ParseUserLine(line string) (id int, name, email string, err error) {
	line = strings.TrimSpace(line)
	if len(line) == 0 {
		return 0, "", "", fmt.Errorf("Пустое значение")
	}

	var data []string = strings.Split(line, "|")

	if len(data) != 3 {
		return 0, "", "", fmt.Errorf("Неполный формат - всего 3 поля: id, name, email")
	}

	// Parse ID
	id, err = strconv.Atoi(strings.TrimSpace(data[0]))
	if err != nil {
		return 0, "", "", fmt.Errorf("Неверный формат ID: %w", err)
	}

	// Parse Name
	name = strings.TrimSpace(data[1])
	if len(name) == 0 {
		return 0, "", "", fmt.Errorf("Имя пользователя не может быть пустым")
	}

	// Parse Email
	email = strings.TrimSpace(data[2])

	if isValidate, err := IsValidateEmail(email); !isValidate && err != nil {
		return 0, "", "", fmt.Errorf("Ошибка валидации email: %w", err)
	}

	return id, name, email, nil
}

func FormatBookLine(id int, title, author string, year int, available bool) string {
	var availableString string = StatusBooleanToString(available)
	return fmt.Sprintf("%d|%s|%s|%d|%s", id, title, author, year, availableString)
}

func FormatUserLine(id int, name, email string) string {
	return fmt.Sprintf("%d|%s|%s", id, name, email)
}
