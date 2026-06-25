package storage

import (
	"fmt"
	"strings"

	"github.com/attikin-draziwee/library_management_system_golang/utils"
)

func ReadAllBooks() ([]string, error) {
	return ReadLines(booksFile)
}

func AppendBook(bookLine string) error {
	return AppendLine(booksFile, bookLine)
}

func RemoveBook(bookID int) error {
	lines, err := ReadAllBooks()
	if err != nil {
		return err
	}

	var found bool = false
	var newLines []string
	for _, line := range lines {
		id, _, _, _, _, err := utils.ParseBookLine(line)
		if err != nil {
			newLines = append(newLines, line)
			continue
		}
		if id == bookID {
			found = true
			continue
		}
		newLines = append(newLines, line)
	}

	if !found {
		return fmt.Errorf("Книга с ID %d не найдена", bookID)
	}
	return WriteLines(booksFile, newLines)
}

func UpdateBookStatusInFile(bookID int, status string) error {
	newStatus, err := utils.StatusStringToBoolean(status)
	if err != nil {
		return err
	}

	lines, err := ReadAllBooks()
	if err != nil {
		return err
	}

	var found bool = false
	var newLines []string
	for _, line := range lines {
		id, title, author, year, _, err := utils.ParseBookLine(line)
		if err != nil {
			newLines = append(newLines, line)
			continue
		}
		if id == bookID {
			found = true
			var updatedLine string = utils.FormatBookLine(id, title, author, year, newStatus)
			newLines = append(newLines, updatedLine)
		} else {
			newLines = append(newLines, line)
		}
	}

	if !found {
		return fmt.Errorf("Книга с ID %d не найдена", bookID)
	}
	return WriteLines(booksFile, newLines)
}

func FindBookLineByID(bookID int) (string, error) {
	lines, err := ReadAllBooks()
	if err != nil {
		return "", err
	}

	for _, line := range lines {
		id, _, _, _, _, err := utils.ParseBookLine(line)
		if err != nil {
			continue
		}
		if id == bookID {
			return line, nil
		}
	}

	return "", fmt.Errorf("Книга с ID %d не найдена", bookID)
}

func GetBookInfo(bookID int) (title, author string, year int, available bool, err error) {
	line, err := FindBookLineByID(bookID)
	if err != nil {
		return "", "", 0, false, err
	}

	_, title, author, year, available, err = utils.ParseBookLine(line)
	if err != nil {
		return "", "", 0, false, err
	}

	return title, author, year, available, nil
}

func SearchBooksByQuery(query string) ([]string, error) {
	lines, err := ReadAllBooks()
	if err != nil {
		return nil, err
	}

	query = strings.ToLower(query)
	var resultList []string
	for _, line := range lines {
		_, title, author, _, _, err := utils.ParseBookLine(line)
		if err != nil {
			continue
		}
		if strings.Contains(strings.ToLower(title), query) || strings.Contains(strings.ToLower(author), query) {
			resultList = append(resultList, line)
		}
	}

	return resultList, nil
}
