package utils

import (
	"fmt"
	"strings"
)

func StatusBooleanToString(status bool) string {
	switch status {
	case true:
		return "да"
	case false:
		return "нет"
	default:
		return "нет"
	}
}

func StatusStringToBoolean(status string) (bool, error) {
	var trimLowerStatus string = strings.TrimSpace(strings.ToLower(status))

	if len(trimLowerStatus) == 0 {
		return false, fmt.Errorf("Статус не может быть пустым.")
	}

	if trimLowerStatus != Agree && trimLowerStatus != Disagree {
		return false, fmt.Errorf("Неверный формат ввода (только \"да\" или \"нет\")")
	}

	switch trimLowerStatus {
	case "да":
		return true, nil
	case "нет":
		return false, nil
	default:
		return false, fmt.Errorf("Ошибка логики")
	}
}
