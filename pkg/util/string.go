package util

import (
	"math/rand/v2"

	"github.com/kye-gregory/koicards-api/pkg/validate"
)

func GenerateRandomString(minLen, maxLen int, charset string) string {
	// Validate lengths
	if (minLen < 0) { minLen = 0 }
	if (maxLen < minLen) { maxLen = minLen }
	if (!validate.MinLength(charset, 1)) { return "" }

	// Convert charset to a slice of runes to handle Unicode characters
	runes := []rune(charset)

	// Generate a random length between minLen and maxLen
	length := rand.IntN(maxLen-minLen+1) + minLen

	// Generate the random string
	randomStr := make([]rune, length)
	for i := range randomStr {
		randomStr[i] = runes[rand.IntN(len(runes))]
	}

	return string(randomStr)
}