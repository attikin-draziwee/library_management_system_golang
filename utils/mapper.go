package utils

import "fmt"

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
	if len(status) == 0 {
		return false, fmt.Errorf("Статус не может быть пустым.")
	}

	if status != Agree && status != Disagree {
		return false, fmt.Errorf("Неверный формат ввода (только \"да\" или \"нет\")")
	}

	switch status {
	case "да":
		return true, nil
	case "нет":
		return false, nil
	default:
		return false, fmt.Errorf("Ошибка логики")
	}
}
