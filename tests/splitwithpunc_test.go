package handlers

import (
	"testing"

	"go-reloaded/handlers"
)

func TestSplitStringWithPunctuation(t *testing.T) {
	testCases := []struct {
		input    []string
		expected []string
	}{
		{
			input:    []string{"Hello,world!", "How", "are", "you?"},
			expected: []string{"Hello, world!", "How", "are", "you?"},
		},
	}

	for _, tc := range testCases {
		result := handlers.SplitStringWithPunctuation(tc.input)
		if !areSlicesEqual(result, tc.expected) {
			t.Errorf("Expected SplitStringWithPunctuation(%v) to be %v, but got %v", tc.input, tc.expected, result)
		}
	}
}

// Check if two slices of strings are equal
func AreSlicesEqual(slice1, slice2 []string) bool {
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
