package handlers

import (
	"reflect"
	"testing"

	"go-reloaded/handlers"
)

func TestPunctuationCorrecting(t *testing.T) {
	testCases := []struct {
		input    []string
		expected []string
	}{
		{
			input:    []string{"'", "m", "happy"},
			expected: []string{"'m", "happy"},
		},
		{
			input:    []string{"'", "s", "a", "great", "day", "'"},
			expected: []string{"'s", "a", "great", "day'"},
		},
	}

	for _, tc := range testCases {
		result := handlers.PunctuationCorrecting(tc.input)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Expected PunctuationCorrecting(%v) to be %v, but got %v", tc.input, tc.expected, result)
		}
	}
}
