package main

import (
	"reflect"
	"testing"
)

// TestFrequencyFiltering проверяет корректность работы функции FrequencyFiltering на различных наборах данных.
func TestFrequencyFiltering(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput map[string]int
	}{
		{
			name:           "Пустая строка",
			input:          "",
			expectedOutput: map[string]int{},
		},
		{
			name:           "Строка с уникальными словами",
			input:          "one two three",
			expectedOutput: map[string]int{},
		},
		{
			name:  "Строка с повторяющимися словами",
			input: "apple apple orange orange banana",
			expectedOutput: map[string]int{
				"apple":  2,
				"orange": 2,
			},
		},
		{
			name:  "Смешанные уникальные и повторяющиеся слова",
			input: "car bike car truck truck bike",
			expectedOutput: map[string]int{
				"car":   2,
				"bike":  2,
				"truck": 2,
			},
		},
		{
			name:           "Слова с учётом регистра",
			input:          "Dog dog DOG",
			expectedOutput: map[string]int{},
		},
		{
			name:  "Строка с повторяющимися словами в разных местах",
			input: "cat dog cat elephant dog cat",
			expectedOutput: map[string]int{
				"cat": 3,
				"dog": 2,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			output := FrequencyFiltering(tt.input)
			if !reflect.DeepEqual(output, tt.expectedOutput) {
				t.Errorf("FrequencyFiltering(%q) = %v; want %v", tt.input, output, tt.expectedOutput)
			}
		})
	}
}
