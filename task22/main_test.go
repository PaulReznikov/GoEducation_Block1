package main

import (
	"reflect"
	"testing"
)

// TestWordsIndexing проверяет корректность работы функции WordsIndexing.
func TestWordsIndexing(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput map[string][]int
	}{
		{
			name:           "Пустая строка",
			input:          "",
			expectedOutput: map[string][]int{},
		},
		{
			name:  "Строка с уникальными словами",
			input: "apple banana cherry",
			expectedOutput: map[string][]int{
				"apple":  {0},
				"banana": {1},
				"cherry": {2},
			},
		},
		{
			name:  "Строка с повторяющимися словами",
			input: "I love Go Go Go",
			expectedOutput: map[string][]int{
				"I":    {0},
				"love": {1},
				"Go":   {2, 3, 4},
			},
		},
		{
			name:  "Строка с разным регистром",
			input: "Hello hello HeLLo",
			expectedOutput: map[string][]int{
				"Hello": {0},
				"hello": {1},
				"HeLLo": {2},
			},
		},
		{
			name:  "Строка с символами",
			input: "word1 word2 word3 word1",
			expectedOutput: map[string][]int{
				"word1": {0, 3},
				"word2": {1},
				"word3": {2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := WordsIndexing(tt.input)
			if !reflect.DeepEqual(output, tt.expectedOutput) {
				t.Errorf("WordsIndexing(%q) = %v; want %v", tt.input, output, tt.expectedOutput)
			}
		})
	}
}
