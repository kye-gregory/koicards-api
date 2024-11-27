package validate

import (
	"testing"
)

func TestMinLength(t *testing.T) {
	tests := []struct {
		input    string
		minLen   int
		expected bool
	}{
		{"hello", 3, true},
		{"hi", 3, false},
		{"", -3, true},
		{"😊😊😊", 4, false},
	}

	for _, test := range tests {
		result := MinLength(test.input, test.minLen)
		if result != test.expected {
			t.Errorf("MinLength(%q, %d) = %v; want %v", test.input, test.minLen, result, test.expected)
		}
	}
}

func TestMaxLength(t *testing.T) {
	tests := []struct {
		input    string
		maxLen   int
		expected bool
	}{
		{"hello", 10, true},
		{"longstring", -2, false},
		{"", 0, true},
		{"😊😊😊", 2, false},
	}

	for _, test := range tests {
		result := MaxLength(test.input, test.maxLen)
		if result != test.expected {
			t.Errorf("MaxLength(%q, %d) = %v; want %v", test.input, test.maxLen, result, test.expected)
		}
	}
}

func TestMinMaxLength(t *testing.T) {
	tests := []struct {
		input    string
		minLen   int
		maxLen   int
		expected bool
	}{
		{"hello", 3, 10, true},
		{"hi", -3, 1, false},
		{"longstring", 3, 5, false},
		{"😊😊", 1, 3, true},
	}

	for _, test := range tests {
		result := MinMaxLength(test.input, test.minLen, test.maxLen)
		if result != test.expected {
			t.Errorf("MinMaxLength(%q, %d, %d) = %v; want %v", test.input, test.minLen, test.maxLen, result, test.expected)
		}
	}
}

func TestOnlyContainsRunes(t *testing.T) {
	tests := []struct {
		input       string
		allowed     string
		expected    bool
	}{
		{"abc", "abc", true},
		{"abc123", "abc", false},
		{"", "abc", true},
		{"😊", "😊", true},
		{"😊👍", "😊", false},
	}

	for _, test := range tests {
		result := OnlyContainsRunes(test.input, test.allowed)
		if result != test.expected {
			t.Errorf("OnlyContainsRunes(%q, %q) = %v; want %v", test.input, test.allowed, result, test.expected)
		}
	}
}

func TestMatchRegex(t *testing.T) {
	tests := []struct {
		input    string
		regex    string
		expected bool
	}{
		{"123", `^\d+$`, true},
		{"abc123", `^\d+$`, false},
		{"hello", `^[a-z]+$`, true},
		{"HELLO", `^[a-z]+$`, false},
		{"😊", `^.$`, true},
	}

	for _, test := range tests {
		result := MatchRegex(test.input, test.regex)
		if result != test.expected {
			t.Errorf("MatchRegex(%q, %q) = %v; want %v", test.input, test.regex, result, test.expected)
		}
	}
}