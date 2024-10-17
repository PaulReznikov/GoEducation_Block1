package main

import (
	"reflect"
	"sort"
	"testing"
)

// TestCommonCharacters проверяет функцию CommonCharacters на различных наборах входных данных.
func TestCommonCharacters(t *testing.T) {
	tests := []struct {
		name           string
		inputA         string
		inputB         string
		expectedResult []rune
	}{
		{
			name:           "Basic case with common characters",
			inputA:         "hello",
			inputB:         "world",
			expectedResult: []rune{'o', 'l'},
		},
		{
			name:           "No common characters",
			inputA:         "abc",
			inputB:         "def",
			expectedResult: []rune{},
		},
		{
			name:           "One empty string",
			inputA:         "",
			inputB:         "nonempty",
			expectedResult: []rune{},
		},
		{
			name:           "Both strings empty",
			inputA:         "",
			inputB:         "",
			expectedResult: []rune{},
		},
		{
			name:           "Multiple common characters with duplicates",
			inputA:         "aabbcc",
			inputB:         "bccddee",
			expectedResult: []rune{'b', 'c'},
		},
		{
			name:           "Unicode characters",
			inputA:         "héllo",
			inputB:         "worldhé",
			expectedResult: []rune{'h', 'é', 'l', 'o'},
		},
		{
			name:           "Case sensitivity",
			inputA:         "Hello",
			inputB:         "hELLo",
			expectedResult: []rune{'h', 'e', 'l', 'o'}, // Приведение регистра не осуществляется
		},
		{
			name:           "Special characters",
			inputA:         "123!@#",
			inputB:         "!@#456",
			expectedResult: []rune{'!', '@', '#'},
		},
		{
			name:           "Repeated characters in both strings",
			inputA:         "aaa",
			inputB:         "aaabbb",
			expectedResult: []rune{'a'},
		},
		{
			name:           "Long strings with multiple common characters",
			inputA:         "thequickbrownfox",
			inputB:         "jumpsoverthelazydog",
			expectedResult: []rune{'t', 'h', 'e', 'o', 'r', 'u'},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CommonCharacters(tt.inputA, tt.inputB)

			// Поскольку порядок элементов в результате не гарантирован, сортируем оба среза перед сравнением
			sortRunes := func(runes []rune) []rune {
				sorted := make([]rune, len(runes))
				copy(sorted, runes)
				sort.Slice(sorted, func(i, j int) bool { return sorted[i] < sorted[j] })
				return sorted
			}

			sortedResult := sortRunes(result)
			sortedExpected := sortRunes(tt.expectedResult)

			if !reflect.DeepEqual(sortedResult, sortedExpected) {
				t.Errorf("CommonCharacters(%q, %q) = %v; want %v", tt.inputA, tt.inputB, sortedResult, sortedExpected)
			}
		})
	}
}
