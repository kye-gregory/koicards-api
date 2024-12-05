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
		{1, 1, "a"},
		{-5, 10, ""},
		{5, 0, "abc"},
		{0, 1, "123"},
		{3, 5, "😀😃😄"},
		{1, 3, "abc123😀😃😄"},
	}

	for _, test := range tests {
		// Generate Strings
		result := GenerateRandomString(test.minLen, test.maxLen, test.charSet)

		// Test Length
		validLengthA := validate.MinMaxLength(result, test.minLen, test.maxLen) || test.maxLen < test.minLen
		if !validLengthA {
			t.Errorf("GenerateRandomString(%d, %d, %q) = %v; invalid length.", test.minLen, test.maxLen, test.charSet, result)
		}

		// Test Characters
		validCharsA := validate.OnlyContainsRunes(result, test.charSet)
		if !validCharsA {
			t.Errorf("GenerateRandomString(%d, %d, %q) = %v; invalid characters.", test.minLen, test.maxLen, test.charSet, result)
		}
	}

	// Test Variability
	results := []string{}
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789~`!@#$%^&*()-_+=[]{}<>,.'\\/\":; \n\té漢🙂😀😃😄"
	for i := 0; i < 100; i++ {
		result := GenerateRandomString(10, 10, charset)
		isDuplicate := false
		for _, ele := range results {
			if ele == result { isDuplicate = true; break }
		}

		if (!isDuplicate) { results = append(results, result) }
	}

	if len(results) < 95 {
		t.Errorf("GenerateRandomString() duplicates out of 100 = %v; invalid varience.", 100 - len(results))
	}
}