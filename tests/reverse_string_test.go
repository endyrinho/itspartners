package reverse_string_test

import (
	"itspartners_task/reverse_string"
	"testing"
)

func TestReverseString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "olleh"},
		{"", ""},
		{"a", "a"},
		{"abcdefghijklmnopqrstuvwxyz", "zyxwvutsrqponmlkjihgfedcba"},
		{"аав", "ваа"},
	}

	for _, test := range tests {
		result := reverse_string.ReverseString(test.input)
		if result != test.expected {
			t.Errorf("ReverseString(%q) returned %q, expected %q", test.input, result, test.expected)
		}
	}
}
