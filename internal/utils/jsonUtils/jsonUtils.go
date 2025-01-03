package jsonUtils

import "encoding/json"

func ConvertInterestsToList(interests json.RawMessage) ([]string, error) {
	var interestsList []string
	// Сначала распарсить сам массив JSON в строку
	err := json.Unmarshal(interests, &interestsList)
	if err != nil {
		return nil, err
	}
	return interestsList, nil
}
