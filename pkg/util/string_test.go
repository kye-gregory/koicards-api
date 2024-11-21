package util

import (
	"testing"

	"github.com/kye-gregory/koicards-api/pkg/validate"
)


func TestGenerateRandomString(t *testing.T) {
	tests := []struct {
		minLen   	int
		maxLen   	int
		charSet	 	string
	}{
		{-5, 10, ""},
		{5, 0, "abc"},
		{0, 1, "123"},
		{3, 5, "😀😃😄"},
		{1, 3, "abc123😀😃😄"},
	}

	for _, test := range tests {
		// Generate Strings
		resultA := GenerateRandomString(test.minLen, test.maxLen, test.charSet)
		resultB := GenerateRandomString(test.minLen, test.maxLen, test.charSet)

		// Test Length
		validLengthA := validate.MinLength(resultA, test.minLen) && (validate.MaxLength(resultA, test.maxLen) || test.maxLen < test.minLen)
		validLengthB := validate.MinLength(resultB, test.minLen) && (validate.MaxLength(resultB, test.maxLen) || test.maxLen < test.minLen)
		if !validLengthA {
			t.Errorf("GenerateRandomString(%d, %d, %q) = %v; invalid length.", test.minLen, test.maxLen, test.charSet, resultA)
		}
		if !validLengthB {
			t.Errorf("GenerateRandomString(%d, %d, %q) = %v; invalid length.", test.minLen, test.maxLen, test.charSet, resultB)
		}

		// Test Characters
		validCharsA := validate.OnlyContainsRunes(resultA, test.charSet)
		validCharsB := validate.OnlyContainsRunes(resultB, test.charSet)
		if !validCharsA {
			t.Errorf("GenerateRandomString(%d, %d, %q) = %v; invalid characters.", test.minLen, test.maxLen, test.charSet, resultA)
		}
		if !validCharsB {
			t.Errorf("GenerateRandomString(%d, %d, %q) = %v; invalid characters.", test.minLen, test.maxLen, test.charSet, resultB)
		}

		// Test Variability
		if resultA == resultB && validate.MinLength(test.charSet, 1) {
			t.Errorf("GenerateRandomString(%d, %d, %q); A = %v; B = %v; duplicate results.", test.minLen, test.maxLen, test.charSet, resultA, resultB)
		}
	}
}