package handlers

import (
	"reflect"
	"testing"

	"go-reloaded/handlers"
)

func TestAToAn(t *testing.T) {
	testCases := []struct {
		input    []string
		expected []string
	}{
		{input: []string{"a", "apple"}, expected: []string{"an", "apple"}},
		{input: []string{"A", "banana"}, expected: []string{"A", "banana"}},
		{input: []string{"a", "car"}, expected: []string{"a", "car"}},
		{input: []string{"A", "dog"}, expected: []string{"A", "dog"}},
	}

	for _, tc := range testCases {
		result := handlers.AToAn(tc.input)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Expected AToAn(%v) to be %v, but got %v", tc.input, tc.expected, result)
		}
	}
}
