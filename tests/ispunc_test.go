package handlers

import (
	"testing"

	"go-reloaded/handlers"
)

func TestIsPunctuation(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{input: ".", expected: true},
		{input: "!", expected: true},
		{input: "a", expected: false},
		{input: "123", expected: false},
	}

	for _, tc := range testCases {
		result := handlers.IsPunctuation(tc.input)
		if result != tc.expected {
			t.Errorf("Expected IsPunctuation(%s) to be %v, but got %v", tc.input, tc.expected, result)
		}
	}
}
