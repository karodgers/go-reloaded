package handlers

import (
	"testing"

	"go-reloaded/handlers"
)

func TestMovePunctuation(t *testing.T) {
	testCases := []struct {
		input    []string
		expected []string
	}{
		{
			input:    []string{"Hello", ",", "world", "!"},
			expected: []string{"Hello,", "world!"},
		},
		{
			input:    []string{"This", "is", "a", "test", "."},
			expected: []string{"This", "is", "a", "test."},
		},
	}

	for _, tc := range testCases {
		result := handlers.MovePunctuation(tc.input)
		if !areSlicesEqual(result, tc.expected) {
			t.Errorf("Expected MovePunctuation(%v) to be %v, but got %v", tc.input, tc.expected, result)
		}
	}
}

// Checks if two slices of strings are equal
func areSlicesEqual(slice1, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}
