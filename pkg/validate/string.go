package validate

import (
	"log"
	"math"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/agnivade/levenshtein"
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

func RemoveRepeatedChars(input string, maxRepeat int) string {
	var builder strings.Builder
	for i := 0; i < len(input); i++ {
		// Append the current character only if it's not the same as the previous one
		if i <= maxRepeat-1 || input[i] != input[i-maxRepeat] {
			builder.WriteByte(input[i])
		}
	}
	return builder.String()
}

func Normalise(input string) string {
	normalised := strings.ToLower(input)

	// Replace common substitutions
    substitutions := map[string]string{
        "1": "i", "0": "o", "@": "a", "$": "s", "3": "e", "4": "a", "5": "s", "8": "b",
    }
    for old, new := range substitutions {
        normalised = strings.ReplaceAll(normalised, old, new)
    }

	// Normalise repeated characters
    normalised = RemoveRepeatedChars(normalised, 2)
	
	// Remove all non alpha characters
    reNonAlpha := regexp.MustCompile(`[^a-zA-Z]`)
    normalised = reNonAlpha.ReplaceAllString(normalised, "")

    return normalised
}

func IsRestricted(input string, blacklist, whitelist []string) ([]string, bool) {
	normalised := Normalise(input)
	log.Printf("\n\nChecking Against Input: '%s'", input)

	output := []string{}
    for _, blWord := range blacklist {
		// Is Restricted
		log.Printf("Exact Substring: '%s' against blacklisted word '%s'", normalised, blWord)
		if strings.Contains(normalised, blWord) { return []string{blWord}, true} // { return blWord, true }

		// Sliding Window Fuzzy Search
		thresholdDist := 1
        wordLength := len(blWord)
		windowSize := wordLength + thresholdDist
		for i := 0; i <= len(normalised)-windowSize; i++ {
			// Extract the substring from input based on the window size
			substr := normalised[i : i+windowSize]

			// Compare each possible substring within the sliding window
			for j := 0; j <= len(substr)-wordLength; j++ {
				window := substr[j : j+wordLength]
				// Match Found
				log.Printf("Fuzzy Searching: '%s' against blacklisted word '%s'", window, blWord)
				if levenshtein.ComputeDistance(window, blWord) <= thresholdDist {
					isWhitelisted := false
					for _, wlWord := range whitelist {
						// Only proceed if detected window could fit within the whitelist word
						if len(window) <= len(wlWord) {
							// Adjust window boundaries to check the larger context
							lenDiff := len(wlWord) - len(window)
							start := int(math.Max(0, float64(i+j-lenDiff)))
							end := int(math.Min(float64(len(normalised)), float64(i+j+wordLength+lenDiff)))

							contextWindow := normalised[start:end]
							log.Printf("Checking: '%s' against whitelisted word '%s'", contextWindow, wlWord)

							// Check if the expanded window contains the whitelist word
							if strings.Contains(contextWindow, wlWord) {
								isWhitelisted = true

								// Move outer loop index to the end of the whitelisted word
								i = start + strings.Index(contextWindow, wlWord) + len(wlWord)
								break
							}
						}
					}

					if (!isWhitelisted) {
						output = append(output, blWord)
					}
				}
			}
		}
	}

	// Default Return
    return output, false
}