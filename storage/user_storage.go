package storage

import (
	"fmt"
	"strings"

	"github.com/attikin-draziwee/library_management_system_golang/utils"
)

func ReadAllUsers() ([]string, error) {
	return ReadLines(usersFile)
}

func AppendUser(userLine string) error {
	return AppendLine(usersFile, userLine)
}

func RemoveUser(userID int) error {
	lines, err := ReadAllUsers()
	if err != nil {
		return err
	}

	var found bool = false
	var newLines []string
	for _, line := range lines {
		id, _, _, err := utils.ParseUserLine(line)
		if err != nil {
			newLines = append(newLines, line)
			continue
		}
		if id == userID {
			found = true
			continue
		}
		newLines = append(newLines, line)
	}

	if !found {
		return fmt.Errorf("Пользователь с ID %d не найдена", userID)
	}
	return WriteLines(usersFile, newLines)
}

func FindUserLineByID(userID int) (string, error) {
	lines, err := ReadAllUsers()
	if err != nil {
		return "", err
	}

	for _, line := range lines {
		id, _, _, err := utils.ParseUserLine(line)
		if err != nil {
			continue
		}
		if id == userID {
			return line, nil
		}
	}

	return "", fmt.Errorf("Пользователь с ID %d не найдена", userID)
}

func GetUserInfo(userID int) (name, email string, err error) {
	line, err := FindUserLineByID(userID)
	if err != nil {
		return "", "", err
	}

	_, name, email, err = utils.ParseUserLine(line)
	if err != nil {
		return "", "", err
	}

	return name, email, nil
}

func SearchUsersByQuery(query string) ([]string, error) {
	lines, err := ReadAllUsers()
	if err != nil {
		return nil, err
	}

	query = strings.ToLower(query)
	var resultList []string
	for _, line := range lines {
		_, name, email, err := utils.ParseUserLine(line)
		if err != nil {
			continue
		}
		if strings.Contains(strings.ToLower(name), query) || strings.Contains(strings.ToLower(email), query) {
			resultList = append(resultList, line)
		}
	}

	return resultList, nil
}
