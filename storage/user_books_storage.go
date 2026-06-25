package storage

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
)

func GetUserBooksFilePath(userID int) string {
	filepath := filepath.Join(userBooksDir, strconv.Itoa(userID))
	return filepath
}

func ReadUserBooks(userID int) ([]int, error) {
	var filepath string = GetUserBooksFilePath(userID)
	if !FileExists(filepath) {
		return []int{}, nil
	}

	lines, err := ReadLines(filepath)
	if err != nil {
		return nil, err
	}

	var bookIDs []int = make([]int, 0)
	for _, line := range lines {
		bookID, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil {
			continue
		}
		bookIDs = append(bookIDs, bookID)
	}

	return bookIDs, nil
}

func WriteUserBooks(userID int, bookIDs []int) error {
	var filepath string = GetUserBooksFilePath(userID)
	var lines []string

	for _, bookID := range bookIDs {
		lines = append(lines, strconv.Itoa(bookID)+"\n")
	}

	return WriteLines(filepath, lines)
}

func AddUserBook(userID int, bookID int) error {
	bookIDs, err := ReadUserBooks(userID)
	if err != nil {
		return err
	}

	if slices.Contains(bookIDs, bookID) {
		return fmt.Errorf("Книга с ID %d уже есть у пользователя с ID %d", bookID, userID)
	}

	bookIDs = append(bookIDs, bookID)
	return WriteUserBooks(userID, bookIDs)
}

func RemoveUserBook(userID int, bookID int) error {
	bookIDs, err := ReadUserBooks(userID)
	if err != nil {
		return err
	}

	if !slices.Contains(bookIDs, bookID) {
		return fmt.Errorf("У пользователя ID %d нет такой книги.", userID)
	}

	var found bool = false
	var modifiedBookIDs []int
	for _, id := range bookIDs {
		if id == bookID {
			found = true
			continue
		}
		modifiedBookIDs = append(modifiedBookIDs, id)
	}

	if !found {
		return fmt.Errorf("Книга с ID %d не найдена у пользователя %d", bookID, userID)
	}
	return WriteUserBooks(userID, modifiedBookIDs)
}

func DeleteUserBooksFile(userID int) error {
	filepath := GetUserBooksFilePath(userID)
	if !FileExists(filepath) {
		return nil
	}

	if err := os.Remove(filepath); err != nil {
		return fmt.Errorf("Не удалось удалить файл книг пользователя ID %d: %w", userID, err)
	}

	return nil
}

func UserHasBook(userID int, bookID int) (bool, error) {
	bookIDs, err := ReadUserBooks(userID)
	if err != nil {
		return false, err
	}

	if slices.Contains(bookIDs, bookID) {
		return true, nil
	}

	return false, nil
}
