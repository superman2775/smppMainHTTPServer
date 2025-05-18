package utils

import (
	"strings"
	"unicode"
)

func ToUpperCase(s string) string {
	if len(s) == 0 {
		return s
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
func RemoveDash(s string) string {
	return strings.ReplaceAll(s, "-", " ")
}
