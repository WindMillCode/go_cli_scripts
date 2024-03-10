package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"regexp"

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



func RemoveComments(jsonStr string) (string, error) {
	// Remove line comments
	lineCommentsRegex := regexp.MustCompile(`(?m)//.*$`)
	jsonStr = lineCommentsRegex.ReplaceAllString(jsonStr, "")

	// Remove block comments
	blockCommentsRegex := regexp.MustCompile(`(?s)/\*.*?\*/`)
	jsonStr = blockCommentsRegex.ReplaceAllString(jsonStr, "")

	// Check for JSON validity
	if !json.Valid([]byte(jsonStr)) {
		return "", fmt.Errorf("invalid JSON after removing comments")
	}

	return jsonStr, nil
}
