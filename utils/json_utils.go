package utils

import "encoding/json"

func FilterJSONByPredicate(inputJSON []byte, predicate func(key string, value interface{}) bool) ([]byte, error) {
	var jsonData map[string]interface{}

	if err := json.Unmarshal(inputJSON, &jsonData); err != nil {
		return nil, err
	}

	filteredData := make(map[string]interface{})

	for key, value := range jsonData {
		if predicate(key, value) {
			filteredData[key] = value
		}
	}

	filteredJSON, err := json.Marshal(filteredData)
	if err != nil {
		return nil, err
	}

	return filteredJSON, nil
}
