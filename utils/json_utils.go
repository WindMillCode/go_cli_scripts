package utils

import (
	"encoding/json"
	"os"
)

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



func WriteCustomFormattedJSONToFile(data interface{}, filename string, indentString string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", indentString)

	if err := encoder.Encode(data); err != nil {
		return err
	}

	return nil
}

func WriteFormattoJSONFile(data interface{}, filename string){
	WriteCustomFormattedJSONToFile(data,filename,"    ")
}
