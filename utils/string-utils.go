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
