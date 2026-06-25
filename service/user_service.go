package service

import (
	"fmt"
	"strings"

	"github.com/attikin-draziwee/library_management_system_golang/storage"
	"github.com/attikin-draziwee/library_management_system_golang/utils"
)

func GenerateNextUserID() (int, error) {
	lines, err := storage.ReadAllUsers()
	if err != nil {
		return 0, err
	}

	var maxID int = 0
	for _, line := range lines {
		id, _, _, err := utils.ParseUserLine(line)
		if err == nil && id > maxID {
			maxID = id
		}
		continue
	}

	return maxID + 1, nil
}

func AddUser(name, email string) error {
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)
	validEmail, emailErr := utils.IsValidateEmail(email)

	switch {
	case len(name) == 0 || len(email) == 0:
		return fmt.Errorf("Пустое значение для имени и/или почты")
	case validEmail == false || emailErr != nil:
		return fmt.Errorf("Неверный формат email: %w", emailErr)
	}

	maxID, err := GenerateNextUserID()
	if err != nil {
		return err
	}

	userLine := utils.FormatUserLine(maxID, name, email)

	if err := storage.AppendUser(userLine); err != nil {
		return err
	}

	utils.DisplayFormatter("Пользователь '%s' успешно добавлен (ID: %d)", name, maxID)
	return nil
}

func GetUserBookCount(userID int) (int, error) {
	bookIDs, err := storage.ReadUserBooks(userID)
	if err != nil {
		return 0, fmt.Errorf("Ошибка при получении книг у пользователя %d: %w", userID, err)
	}

	return len(bookIDs), nil
}

func ListUsers() error {
	lines, err := storage.ReadAllUsers()
	switch {
	case err != nil:
		return err
	case len(lines) == 0:
		utils.DisplayLine("\nСписок пользователя пуст")
		return nil
	}

	utils.DisplayLine("\n---Список пользователей---")
	utils.DisplayFormatter("%-5s | %-20s | %-64s | %-12s\n", "ID", "Имя", "Email", "Кол-во книг")
	utils.DisplayLine(strings.Repeat("-", 110))

	var validUsers int = 0
	for _, line := range lines {
		id, name, email, err := utils.ParseUserLine(line)
		if err != nil {
			utils.DisplayFormatter("Предупреждение: некорректная строка - %s\n", line)
			continue
		}
		booksCount, _ := GetUserBookCount(id)

		validUsers += 1
		utils.DisplayFormatter("%-5d | %-20s | %-64s | %-12d\n", id, name, email, booksCount)
	}

	utils.DisplayFormatter("\n%sКоличество пользователей: %d\n", strings.Repeat("-", 29), validUsers)

	return nil
}

func SearchUsers(query string) error {
	query = strings.TrimSpace(strings.ToLower(query))
	if len(query) == 0 {
		return fmt.Errorf("Поисковый запрос не может быть пустым")
	}

	results, err := storage.SearchUsersByQuery(query)
	if err != nil {
		return err
	}

	if len(results) == 0 {
		utils.DisplayFormatter("По вашему запросу %s ничего не найдено.", query)
		return nil
	}

	utils.DisplayFormatter("\n---Результаты поиска по запросу: %s---\n", query)
	utils.DisplayFormatter("%-5s | %-20s | %-64s | %-12s\n", "ID", "Имя", "Email", "Кол-во книг")
	utils.DisplayLine(strings.Repeat("-", 110))

	var foundCount int = 0
	for _, line := range results {
		id, name, email, err := utils.ParseUserLine(line)
		if err != nil {
			continue
		}
		foundCount += 1
		booksCount, _ := GetUserBookCount(id)
		utils.DisplayFormatter("%-5d | %-20s | %-64s | %-12d\n", id, name, email, booksCount)
	}

	utils.DisplayFormatter("\n%sКоличество записей: %d\n", strings.Repeat("-", 29), foundCount)
	return nil
}

func UserExists(userID int) (bool, string, error) {
	name, _, err := storage.GetUserInfo(userID)
	if err != nil {
		if strings.Contains(err.Error(), "не найден") {
			return false, "", nil
		}
		return false, "", err
	}

	return true, name, nil
}

func DeleteUser(userID int) error {
	exists, name, err := UserExists(userID)
	switch {
	case err != nil:
		return err
	case !exists:
		return fmt.Errorf("Пользователь с ID %d не найден", userID)
	}

	if err := storage.DeleteUserBooksFile(userID); err != nil {
		return fmt.Errorf("Ошибка при удалении списка книг у пользователя %d: %w", userID, err)
	}

	if err := storage.RemoveUser(userID); err != nil {
		return fmt.Errorf("Ошибка при удалении пользователя %d: %w", userID, err)
	}

	utils.DisplayFormatter("Пользователь '%s' (ID: %d) был удален", name, userID)

	return nil
}
