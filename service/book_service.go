package service

import (
	"fmt"
	"strings"

	"github.com/attikin-draziwee/library_management_system_golang/storage"
	"github.com/attikin-draziwee/library_management_system_golang/utils"
)

func GenerateNextBookID() (int, error) {
	lines, err := storage.ReadAllBooks()
	if err != nil {
		return 0, err
	}

	var maxID int = 0
	for _, line := range lines {
		id, _, _, _, _, err := utils.ParseBookLine(line)
		if err == nil && id > maxID {
			maxID = id
		}
		continue
	}

	return maxID + 1, nil
}

func IsBookAvailable(bookID int) (bool, error) {
	_, _, _, available, err := storage.GetBookInfo(bookID)
	if err != nil {
		return false, err
	}
	return available, nil
}

func BookExists(bookID int) (bool, string, error) {
	title, _, _, _, err := storage.GetBookInfo(bookID)
	if err != nil {
		return false, "", err
	}
	return true, title, nil
}

func AddBook(title, author string, year int) error {
	title = strings.TrimSpace(title)
	author = strings.TrimSpace(author)
	switch {
	case len(title) == 0:
		return fmt.Errorf("Пустое значение для название книги")
	case len(author) == 0:
		return fmt.Errorf("Пустое значение для имени автора книги")
	case year <= 1000 || year >= 2026:
		return fmt.Errorf("Неверный диапазон ответа (1000÷2026)")
	}

	newID, err := GenerateNextBookID()
	if err != nil {
		return err
	}

	var bookLine string = utils.FormatBookLine(newID, title, author, year, true)

	if err := storage.AppendBook(bookLine); err != nil {
		return err
	}

	utils.DisplayFormatter("Книга '%s' успешно добавлена (ID: %d)", title, newID)
	return nil
}

func ListBook() error {
	lines, err := storage.ReadAllBooks()
	switch {
	case err != nil:
		return err
	case len(lines) == 0:
		utils.DisplayLine("\nСписок книг пуст")
		return nil
	}

	utils.DisplayLine("\n---Список книг---")
	utils.DisplayFormatter("%-5s | %-32s | %-32s | %-6s | %-12s\n", "ID", "Название", "Автор", "Год", "Доступна")
	utils.DisplayLine(strings.Repeat("-", 96))

	var validBooks int = 0
	for _, line := range lines {
		id, title, author, year, available, err := utils.ParseBookLine(line)

		if err != nil {
			utils.DisplayFormatter("Предупреждение: некорректная строка - %s\n", line)
			continue
		}

		validBooks += 1
		utils.DisplayFormatter("%-5d | %-32s | %-32s | %-6d | %-12t\n", id, title, author, year, available)
	}

	utils.DisplayFormatter("\n%sКоличество книг: %d\n", strings.Repeat("-", 29), validBooks)

	return nil
}

func SearchBooks(query string) error {
	query = strings.TrimSpace(strings.ToLower(query))
	if len(query) == 0 {
		return fmt.Errorf("Поисковый запрос не может быть пустым")
	}

	results, err := storage.SearchBooksByQuery(query)
	if err != nil {
		return err
	}

	if len(results) == 0 {
		utils.DisplayFormatter("По вашему запросу %s ничего не найдено.", query)
		return nil
	}

	utils.DisplayFormatter("\n---Результаты поиска по запросу: %s---\n", query)
	utils.DisplayFormatter("%-5s | %-32s | %-32s | %-6s | %-12s\n", "ID", "Название", "Автор", "Год", "Доступна")
	utils.DisplayLine(strings.Repeat("-", 110))

	var foundCount int = 0
	for _, line := range results {
		id, title, author, year, available, err := utils.ParseBookLine(line)
		if err != nil {
			continue
		}

		foundCount += 1
		utils.DisplayFormatter("%-5d | %-32s | %-32s | %-6d | %-12t\n", id, title, author, year, available)
	}

	utils.DisplayFormatter("\n%sКоличество записей: %d\n", strings.Repeat("-", 29), foundCount)
	return nil
}

func DeleteBook(bookID int) error {
	exists, title, err := BookExists(bookID)
	switch {
	case err != nil:
		return err
	case !exists:
		return fmt.Errorf("Книга с ID %d не найдена.", bookID)
	}

	err = storage.RemoveBook(bookID)
	if err == nil {
		utils.DisplayFormatter("Книга '%s' (ID %d) успешно удалена\n", title, bookID)
		return err
	}
	return err
}

func UpdateBookStatus(bookID int, newStatus string) error {
	exists, title, err := BookExists(bookID)
	switch {
	case err != nil:
		return err
	case !exists:
		return fmt.Errorf("Книга с ID %d не найдена.", bookID)
	}

	err = storage.UpdateBookStatusInFile(bookID, newStatus)
	if err == nil {
		utils.DisplayFormatter("Статус книги '%s' (ID %d) успешно обновлен на '%s'", title, bookID, newStatus)
		return err
	}
	return err
}

func BorrowBook(userID int, bookID int) error {
	userExists, userName, err := UserExists(userID)
	switch {
	case err != nil:
		return err
	case !userExists:
		return fmt.Errorf("Пользователь с ID %d не найден", userID)
	}

	bookExists, bookTitle, err := BookExists(bookID)
	switch {
	case err != nil:
		return err
	case !bookExists:
		return fmt.Errorf("Книга с ID %d не найден", userID)
	}

	available, err := IsBookAvailable(bookID)
	switch {
	case err != nil:
		return err
	case !available:
		return fmt.Errorf("Книга с ID %d не доступна (кто-то её взял)", userID)
	}

	if err := storage.AddUserBook(userID, bookID); err != nil {
		return err
	}

	if err := storage.UpdateBookStatusInFile(bookID, "нет"); err != nil {
		storage.RemoveUserBook(userID, bookID)
		return fmt.Errorf("Не удалось обновить статус у книги: %w", err)
	}

	utils.DisplayFormatter("Книга '%s' (ID: %d) выдана пользователю %s (ID: %d)", bookTitle, bookID, userName, userID)
	return nil
}

func ReturnBook(userID int, bookID int) error {
	userExists, userName, err := UserExists(userID)
	switch {
	case err != nil:
		return err
	case !userExists:
		return fmt.Errorf("Пользователь с ID %d не найден", userID)
	}

	bookExists, bookTitle, err := BookExists(bookID)
	switch {
	case err != nil:
		return err
	case !bookExists:
		return fmt.Errorf("Книга с ID %d не найден", userID)
	}

	hasBook, err := storage.UserHasBook(userID, bookID)
	switch {
	case err != nil:
		return err
	case !hasBook:
		return fmt.Errorf("Книга '%s' не находится у пользователя %s", bookTitle, userName)
	}

	if err := storage.RemoveUserBook(userID, bookID); err != nil {
		return err
	}

	if err := storage.UpdateBookStatusInFile(bookID, "да"); err != nil {
		storage.AddUserBook(userID, bookID)
		return fmt.Errorf("Не удалось обновить статус книги: %w", err)
	}

	utils.DisplayFormatter("Книга '%s' (ID: %d) возвращена пользователем %s (ID %d)", bookTitle, bookID, userName, userID)
	return nil
}

func ListUserBooks(userID int) error {
	userExists, userName, err := UserExists(userID)
	switch {
	case err != nil:
		return err
	case !userExists:
		return fmt.Errorf("Пользователь с ID %d не найден", userID)
	}

	bookIDs, err := storage.ReadUserBooks(userID)
	switch {
	case err != nil:
		return err
	case len(bookIDs) == 0:
		utils.DisplayFormatter("\nПользователь %s (ID: %d) не взял ни одной книги\n", userName, userID)
		return nil
	}

	utils.DisplayFormatter("\n---Книги пользователя %s (ID: %d)---\n", userName, userID)
	utils.DisplayFormatter("%-5s | %-32s | %-32s | %-6s\n", "ID", "Название", "Автор", "Год")
	utils.DisplayLine(strings.Repeat("-", 83))

	var foundCount int = 0
	for _, bookID := range bookIDs {
		title, author, year, _, err := storage.GetBookInfo(bookID)
		if err != nil {
			continue
		}

		foundCount += 1
		utils.DisplayFormatter("%-5d | %-32s | %-32s | %-6d\n", bookID, title, author, year)
	}

	utils.DisplayFormatter("\n%sКоличество записей: %d\n", strings.Repeat("-", 29), foundCount)
	return nil
}
