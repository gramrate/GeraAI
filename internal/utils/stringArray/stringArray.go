package stringArray

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// StringArray для обработки JSON-массива строк
type StringArray []string

// Реализация sql.Scanner
func (sa *StringArray) Scan(value interface{}) error {
	// Проверяем тип входящего значения
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to parse value as []byte")
	}
	// Десериализация JSON в []string
	return json.Unmarshal(bytes, sa)
}

// Реализация driver.Valuer
func (sa StringArray) Value() (driver.Value, error) {
	// Сериализация []string в JSON
	return json.Marshal(sa)
}
