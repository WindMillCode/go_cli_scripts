package utils

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/chzyer/readline/runes"
	"github.com/iancoleman/strcase"
)

func IntToStr(val int) string {
  return strconv.Itoa(val)
}

func Reverse(targetArray []string) []string {
	newArray := make([]string, 0, len(targetArray))
	for i := len(targetArray) - 1; i >= 0; i-- {
		newArray = append(newArray, targetArray[i])
	}
	return newArray
}

func UnicodeUnquote(bs []byte) []byte {
	unicodeEscapeRx := regexp.MustCompile(`\\u[0-9a-fA-F]{4}`)
	return unicodeEscapeRx.ReplaceAllFunc(bs, func(code []byte) []byte {
		rune, _, _, _ := strconv.UnquoteChar(string(code), 0)
		width := runes.Width(rune)
		runeBytes := make([]byte, width)
		utf8.EncodeRune(runeBytes, rune)
		return runeBytes
	})
}

func ReplaceAllSubstrings(input string, target string, replacement string) string {
	if target == "" {
		fmt.Println("Target substring cannot be empty.")
		return input
	}

	// Replace all occurrences using strings.ReplaceAll
	result := strings.ReplaceAll(input, target, replacement)
	return result
}

type TruncateStringByRegexOptions struct {
	InputString  string
	RegexPattern string
	Predicate    func(int) bool
}

func TruncateStringByRegex(options TruncateStringByRegexOptions) string {

	regex := regexp.MustCompile(options.RegexPattern)
	matches := regex.FindAllStringIndex(options.InputString, -1)
	currentIndex := 0
	var modifiedString string

	for i := 0; i < len(matches); i++ {
			matchStart, matchEnd := matches[i][0], matches[i][1]

			shouldRemove := true
			if options.Predicate != nil {
				shouldRemove = options.Predicate(i)
			}

			if shouldRemove {
				modifiedString += options.InputString[currentIndex:matchStart]
				currentIndex = matchEnd
			}
	}

	modifiedString += options.InputString[currentIndex:]
	return modifiedString
}

func CreateStringObject(myStr string, entitySuffix string) (CreateStringObjectType, error) {
	if myStr == "" {
		return CreateStringObjectType{}, errors.New("class name is missing or misspelled or the script is having issues finding the class name")
	}

	result := CreateStringObjectType{
		Orig: myStr,
		Prefix: func() string {
			return strings.Split(myStr, entitySuffix)[0]
		},
	}

	result.CamelCase = func(stripSuffix bool, suffix string) string {
		return strcase.ToLowerCamel(grabString(stripSuffix, result))+suffix
	}

	result.Classify = func(stripSuffix bool, suffix string) string {
		return strcase.ToCamel(grabString(stripSuffix, result))+suffix
	}

	result.Capitalize = func(stripSuffix bool, suffix string) string {
		str := grabString(stripSuffix, result)
		if len(str) == 0 {
			return suffix // Return only the suffix if the string is empty
		}
		return strings.ToUpper(str[:1]) + str[1:] + suffix
	}

	result.Dasherize = func(stripSuffix bool, suffix string) string {
		return strcase.ToKebab(grabString(stripSuffix, result))+suffix
	}

	result.Lowercase = func(stripSuffix bool, suffix string) string {
		return strings.ToLower(grabString(stripSuffix, result))+suffix
	}

	result.Uppercase = func(stripSuffix bool, suffix string) string {
		return strings.ToUpper(grabString(stripSuffix, result))+suffix
	}

	result.Snakecase = func(stripSuffix bool, suffix string) string {
		return strcase.ToSnake(grabString(stripSuffix, result))+suffix
	}

	result.KebabCase = func(stripSuffix bool, suffix string) string {
		return strcase.ToKebab(grabString(stripSuffix, result)) + suffix
	}

	result.PascalCase = func(stripSuffix bool, suffix string) string {
		return strcase.ToCamel(grabString(stripSuffix, result)) + suffix
	}





	return result, nil
}

func grabString(stripSuffix bool, result CreateStringObjectType) (string) {
	if stripSuffix {
		return  result.Prefix()
	} else {
		return  result.Orig
	}
}

// CreateStringObjectType represents the structure of the string object.
type CreateStringObjectType struct {
	Orig      string
	Prefix    func() string
	CamelCase func(stripSuffix bool, suffix string) string
	Classify  func(stripSuffix bool, suffix string) string
	Capitalize func(stripSuffix bool, suffix string) string
	Dasherize  func(stripSuffix bool, suffix string) string
	Lowercase  func(stripSuffix bool, suffix string) string
	Uppercase  func(stripSuffix bool, suffix string) string
	Snakecase  func(stripSuffix bool, suffix string) string
	KebabCase  func(stripSuffix bool, suffix string) string
	PascalCase func(stripSuffix bool, suffix string) string
}

func ContainsAny(s string, substrs []string) bool {
	for _, substr := range substrs {
			if strings.Contains(s, substr) {
					return true
			}
	}
	return false
}

func JoinArgs(args []string) string {
	var result string
	for _, arg := range args {
		result += fmt.Sprintf(`"%s" `, arg)
	}
	return result
}

