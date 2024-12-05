package validate

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

func MinLength(input string, minLen int) bool {
	len := utf8.RuneCountInString(input)
	return len >= minLen
}

func MaxLength(input string, maxLen int) bool {
	len := utf8.RuneCountInString(input)
	return len <= maxLen
}

func MinMaxLength(input string, minLen, maxLen int) bool {
	len := utf8.RuneCountInString(input)
	return len >= minLen && len <= maxLen
}

func OnlyContainsRunes(input, allowedChars string) bool {
	for _, char := range input {
		if !strings.ContainsRune(allowedChars, char) {
			return false
		}
	}
	return true
}

func MatchRegex(input, regex string) bool {
	match, _ := regexp.MatchString(regex, input)
	return match
}