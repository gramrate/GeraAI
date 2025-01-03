package parser

import (
	"fmt"
	"strconv"
	"strings"
)

// StringToArrayStr превращает строку в массив []string
func StringToArrayStr(input string) []string {
	trimmed := strings.Trim(input, "[]")
	if trimmed == "" {
		return []string{}
	}
	return strings.Split(trimmed, ";")
}

// StrArrayToString превращает массив []string в строку
func StrArrayToString(arr []string) string {
	// Соединяем элементы массива через запятую
	return "[" + strings.Join(arr, ";") + "]"
}

// StringToArrayUint превращает строку в массив []uint
func StringToArrayUint(input string) ([]uint, error) {
	trimmed := strings.Trim(input, "[]") // Убираем квадратные скобки
	if trimmed == "" {
		return []uint{}, nil
	}

	parts := strings.Split(trimmed, ";")
	result := make([]uint, len(parts))

	for i, part := range parts {
		num, err := strconv.ParseUint(strings.TrimSpace(part), 10, 64) // Парсим каждую строку как uint
		if err != nil {
			return nil, fmt.Errorf("Error converting string to uint: %v\n", err)
		}
		result[i] = uint(num)
	}

	return result, nil
}

// UintArrayToString превращает массив []uint в строку
func UintArrayToString(arr []uint) string {
	parts := make([]string, len(arr))
	for i, num := range arr {
		parts[i] = strconv.FormatUint(uint64(num), 10) // Преобразуем uint в строку
	}
	return "[" + strings.Join(parts, ";") + "]"
}

func StringToUint(str string) (uint, error) {
	numInt, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return uint(numInt), nil
}
