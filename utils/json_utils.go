package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

)

func ParseJSONFromString[T any](jsonString string, target *T) error {
	err := json.Unmarshal([]byte(jsonString), target)
	if err != nil {
		return fmt.Errorf("error parsing JSON: %v", err)
	}
	return nil
}


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
	var out bytes.Buffer
	switch d := data.(type) {
	case []byte:
		dataBytes = d

	default:
		var err error
		dataBytes, err = json.MarshalIndent(d, "", indentString)
		if err != nil {
			return err
		}
	}
	json.Indent(&out, UnicodeUnquote(dataBytes), "", indentString)

	out.WriteTo(file)
	if err != nil {
		return err
	}

	return nil
}
func WriteFormattoJSONFile(data interface{}, filename string){
	WriteCustomFormattedJSONToFile(data,filename,"    ")
}



func RemoveComments(data []byte) ([]byte, error) {
	inString := false
	inSingleLineComment := false
	inMultiLineComment := false
	var result []byte

	for i := 0; i < len(data); i++ {
		if inSingleLineComment {
			if data[i] == '\n' {
				inSingleLineComment = false
			}
			continue
		}

		if inMultiLineComment {
			if data[i] == '*' && i+1 < len(data) && data[i+1] == '/' {
				inMultiLineComment = false
				i++ // Skip the '/'
			}
			continue
		}

		if inString {
			if data[i] == '"' && data[i-1] != '\\' {
				inString = false
			}
			result = append(result, data[i])
			continue
		}

		if data[i] == '"' {
			inString = true
			result = append(result, data[i])
			continue
		}

		if data[i] == '/' && i+1 < len(data) {
			if data[i+1] == '/' {
				inSingleLineComment = true
				i++ // Skip the next '/'
				continue
			} else if data[i+1] == '*' {
				inMultiLineComment = true
				i++ // Skip the next '*'
				continue
			}
		}

		result = append(result, data[i])
	}

	if !json.Valid(result) {
		return nil, fmt.Errorf("Hey there your JSON is no good  after removing comments. Try to paste the json into an online JSON validator to find the issue (Comments should work fine its usually you have a comma at the end of the last element in an array in your JSON which is invalid)")
	}

	return result, nil
}
