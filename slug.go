package slug

import (
	"strings"
	"unicode"
)

func Slug(str string) string {
	return slug(strings.ToLower(str))
}

func slug(str string) string {
	str = strings.TrimFunc(str, trimRune)
	return strings.Map(mapRune, str)
}

const (
	underscore = '_'
	minus      = '-'
)

func trimRune(c rune) bool {
	return unicode.IsSpace(c) || c == minus || c == underscore
}

func mapRune(c rune) rune {
	switch {
	case unicode.IsSpace(c):
		return underscore
	case keep(c):
		return c
	default:
		return -1
	}
}

func keep(c rune) bool {
	return unicode.IsDigit(c) || unicode.IsLower(c) || unicode.IsUpper(c) || c == underscore || c == minus
}
