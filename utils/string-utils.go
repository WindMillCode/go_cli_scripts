package utils

import (
	"regexp"
	"strconv"
	"unicode/utf8"

	"github.com/chzyer/readline/runes"
)

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
