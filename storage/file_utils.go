package storage

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	dataDir      = "data"
	booksFile    = "data/books.txt"
	usersFile    = "data/users.txt"
	userBooksDir = "data/user_books"
)

func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil
}

func createFileIfNotExists(filePath string) error {
	if FileExists(filePath) {
		return fmt.Errorf("Такой файл %s уже существует", filePath)
	}

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("Ошибка в создании файла %s: %w", filePath, err)
	}

	file.Close()
	return nil
}

func InitStorage() error {
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return fmt.Errorf("Ошибка при иницилизации хранилища %s: %w", dataDir, err)
	}

	if err := os.MkdirAll(userBooksDir, 0755); err != nil {
		return fmt.Errorf("Ошибка при иницилизации хранилища %s: %w", userBooksDir, err)
	}

	if err := createFileIfNotExists(booksFile); err != nil {
		return err
	}

	if err := createFileIfNotExists(usersFile); err != nil {
		return err
	}

	return nil
}

func ReadLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("Не удалось открыть файл %s, %w", filePath, err)
	}
	defer file.Close()

	var lines []string
	var buffer *bufio.Scanner = bufio.NewScanner(file)
	for buffer.Scan() {
		line := strings.TrimSpace(buffer.Text())
		if len(line) != 0 {
			lines = append(lines, line)
		}
	}

	if err := buffer.Err(); err != nil {
		return nil, fmt.Errorf("Ошибка при чтении файла %s, %w", filePath, err)
	}

	return lines, nil
}

func WriteLines(filePath string, lines []string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("Не удалось создать файл %s: %w", filePath, err)
	}
	defer file.Close()

	var buffer *bufio.Writer = bufio.NewWriter(file)
	for _, line := range lines {
		if _, err := buffer.WriteString(line + "\n"); err != nil {
			return fmt.Errorf("Ошибка при записи данных %s: %w", filePath, err)
		}
	}

	if err := buffer.Flush(); err != nil {
		return fmt.Errorf("Ошибка при сохранении в файл %s: %w", filePath, err)
	}

	return nil
}

func AppendLine(filePath string, line string) error {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("Не удалось открыть файл %s: %w", filePath, err)
	}
	defer file.Close()

	var buffer *bufio.Writer = bufio.NewWriter(file)
	if _, err := buffer.WriteString(line + "\n"); err != nil {
		return fmt.Errorf("Не удалось записать в файл %s: %w", filePath, err)
	}

	if err := buffer.Flush(); err != nil {
		return fmt.Errorf("Не удалось сохранить в файл %s: %w", filePath, err)
	}

	return nil
}
