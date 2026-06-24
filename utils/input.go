package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)

func ReadChoice() (int, error) {
	if !scanner.Scan() {
		return 0, fmt.Errorf("Ошибка чтения данных.")
	}

	input := strings.TrimSpace(scanner.Text())
	if len(input) == 0 {
		return 0, fmt.Errorf("Ввод не может быть пустым.")
	}

	choice, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("Неверный формат ответа.")
	}

	return choice, nil
}

func ReadString(prompt string) (string, error) {
	Display(prompt)

	if !scanner.Scan() {
		return "", fmt.Errorf("Ошибка чтения данных")
	}

	var input string = strings.TrimSpace(scanner.Text())
	if len(input) == 0 {
		return "", fmt.Errorf("Ввод не может быть пустым.")
	}

	return input, nil
}

func ReadInt(prompt string) (int, error) {
	Display(prompt)

	if !scanner.Scan() {
		return 0, fmt.Errorf("Ошибка чтения данных.")
	}

	input := strings.TrimSpace(scanner.Text())
	if len(input) == 0 {
		return 0, fmt.Errorf("Ввод не может быть пустым.")
	}

	value, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("Неверный формат ответа.")
	}

	return value, nil
}

func ReadYear(prompt string) (int, error) {
	Display(prompt)

	if !scanner.Scan() {
		return 0, fmt.Errorf("Ошибка чтения данных.")
	}

	input := strings.TrimSpace(scanner.Text())
	if len(input) == 0 {
		return 0, fmt.Errorf("Ввод не может быть пустым.")
	}

	year, err := strconv.Atoi(input)
	switch {
	case err != nil:
		return 0, fmt.Errorf("Неверный формат ответа.")
	case year <= 1000 || year >= 2026:
		return 0, fmt.Errorf("Неверный диапазон ответа (1000÷2026)")
	}

	return year, nil
}

func ReadEmail(prompt string) (string, error) {
	Display(prompt)

	if !scanner.Scan() {
		return "", fmt.Errorf("Ошибка чтения данных.")
	}

	var email string = strings.TrimSpace(scanner.Text())

	if isValidate, err := IsValidateEmail(email); !isValidate && err != nil {
		return "", fmt.Errorf("Ошибка валидации email: %w", err)
	}

	return email, nil
}

const (
	Agree    = "да"
	Disagree = "нет"
)

func ReadStatus(prompt string) (string, error) {
	Display(prompt)

	if !scanner.Scan() {
		return "", fmt.Errorf("Ошибка чтения данных.")
	}

	status := strings.ToLower(strings.TrimSpace(scanner.Text()))
	if len(status) == 0 {
		return "", fmt.Errorf("Ввод не может быть пустым.")
	}

	if status != Agree && status != Disagree {
		return "", fmt.Errorf("Неверный формат ввода (только \"да\" или \"нет\")")
	}

	return status, nil
}
