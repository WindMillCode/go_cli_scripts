package utils

import (
	"encoding/json"
	"fmt"
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

	var dataBytes []byte

	// Check the type of the data parameter
	switch d := data.(type) {
	case []byte:
		dataBytes = d
	case interface{}:
		var err error
		dataBytes, err = json.Marshal(d)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported data type")
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", indentString)

	if err := encoder.Encode(dataBytes); err != nil {
		return err
	}

	return nil
}

func WriteFormattoJSONFile(data interface{}, filename string){
	WriteCustomFormattedJSONToFile(data,filename,"    ")
}
