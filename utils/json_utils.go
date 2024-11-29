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
func WriteFormattoJSONFile(data interface{}, filename string) {
	WriteCustomFormattedJSONToFile(data, filename, "    ")
}

func CleanJSON(data []byte) ([]byte, error) {
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

	// Remove trailing commas from arrays and objects
	cleanedResult, err := removeTrailingCommas(result)
	if err != nil {
		return nil, err
	}

	if !json.Valid(cleanedResult) {

		return nil, fmt.Errorf("Hey there your JSON is no good after removing comments and trailing commas. Try to paste the JSON into an online JSON validator to find the issue (Comments should work fine; it's usually because you have a comma at the end of the last element in an array or object in your JSON, which is invalid).")
	}

	return cleanedResult, nil
}

// Helper function to remove trailing commas for the last element in arrays and objects
func removeTrailingCommas(data []byte) ([]byte, error) {
	var result []byte
	inString := false
	bracketStack := []byte{}

	for i := 0; i < len(data); i++ {
		char := data[i]

		if char == '"' {
			// Count backslashes preceding the quote
			backslashCount := 0
			j := i - 1
			for j >= 0 && data[j] == '\\' {
				backslashCount++
				j--
			}
			// Toggle inString only if backslashes count is even
			if backslashCount%2 == 0 {
				inString = !inString
			}
		}

		if !inString {
			// Track the context (array or object) using a stack
			if char == '{' || char == '[' {
				bracketStack = append(bracketStack, char)
			} else if char == '}' || char == ']' {
				// Remove any trailing commas before the closing bracket
				lastIndex := len(result) - 1
				// Skip backward over whitespace
				for lastIndex >= 0 && (result[lastIndex] == ' ' || result[lastIndex] == '\n' || result[lastIndex] == '\t' || result[lastIndex] == '\r') {
					lastIndex--
				}
				if lastIndex >= 0 && result[lastIndex] == ',' {
					// Remove the comma and any whitespace after it
					result = result[:lastIndex]
					// Remove any whitespace between the comma and the current position
					i = i - 1
					for i >= 0 && (data[i] == ' ' || data[i] == '\n' || data[i] == '\t' || data[i] == '\r') {
						i--
					}
					// Adjust index to current position
					i++
				}
				bracketStack = bracketStack[:len(bracketStack)-1]
			} else if char == ',' {
				// Skip the comma if the next non-whitespace character is a closing bracket
				skipComma := false
				j := i + 1
				for j < len(data) {
					if data[j] == ' ' || data[j] == '\n' || data[j] == '\t' || data[j] == '\r' {
						j++
						continue
					} else if data[j] == '}' || data[j] == ']' {
						skipComma = true
						break
					} else {
						break
					}
				}
				if skipComma {
					continue // Skip the comma
				}
			}
		}

		// Append current character to the result
		result = append(result, char)
	}

	// Validate matching brackets
	if len(bracketStack) != 0 {
		return nil, fmt.Errorf("Invalid JSON structure: unmatched brackets")
	}

	return result, nil
}
