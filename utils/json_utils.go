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

	switch d := data.(type) {
	case []byte:
		// If data is a byte slice, unmarshal it to a data structure
		var jsonData interface{}
		err := json.Unmarshal(d, &jsonData)
		if err != nil {
			return err
		}

		dataBytes, err = json.MarshalIndent(jsonData, "", indentString)
		if err != nil {
			return err
		}
	case interface{}:
		// If data is not a byte slice, marshal it to JSON
		var err error
		dataBytes, err = json.MarshalIndent(d, "", indentString)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported data type")
	}

	_, err = file.Write(dataBytes)
	if err != nil {
		return err
	}

	return nil
}
func WriteFormattoJSONFile(data interface{}, filename string){
	WriteCustomFormattedJSONToFile(data,filename,"    ")
}
