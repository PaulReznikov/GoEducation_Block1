package main

import (
	"reflect"
	"testing"
)

func TestSortingDictionary(t *testing.T) {
	tests := []struct {
		testName string
		input    string
		expected []string
	}{
		{
			testName: "Проверка на пустую строку",
			input:    "",
			expected: []string{},
		},
		{
			testName: "Проверка на наличие одного слова",
			input:    "Abbbbb",
			expected: []string{"Abbbbb"},
		},
		{
			testName: "Проверка на наличие нескольких слов",
			input:    "AA BBB CCC",
			expected: []string{"AA", "BBB", "CCC"},
		},
		{
			testName: "Проверка на наличие специфичных символов в словах",
			input:    "A%%% B! ^ C#",
			expected: []string{"^", "B!", "C#", "A%%%"},
		},
		{
			testName: "Проврека на слова одинаковой длины",
			input:    "AA BB CC DD",
			expected: []string{"AA", "BB", "CC", "DD"},
		},
		{
			testName: "Проверка на слова с разным регистром слов",
			input:    "aA Aa BB bb",
			expected: []string{"aA", "Aa", "BB", "bb"},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.testName, func(t *testing.T) {
			result := SortingDictionary(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("SortingDictionary(%v) = %v, wont = %v", tt.input, result, tt.expected)
			}
		})
	}
}
