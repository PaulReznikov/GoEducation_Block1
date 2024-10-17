package main

import (
	"reflect"
	"testing"
)

// TestCountingWords проверяет функцию CountingWords на различных наборах данных.
func TestCountingWords(t *testing.T) {
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
			name:  "Одинарное слово",
			input: "hello",
			expectedOutput: map[string]int{
				"hello": 1,
			},
		},
		{
			name:  "Повторение одного слова",
			input: "hello hello hello",
			expectedOutput: map[string]int{
				"hello": 3,
			},
		},
		{
			name:  "Слова с разным регистром",
			input: "Hello HELLO hello",
			expectedOutput: map[string]int{
				"hello": 3,
			},
		},
		{
			name:  "Слова со знаками препинания",
			input: "hello, world! hello. world?",
			expectedOutput: map[string]int{
				"hello": 2,
				"world": 2,
			},
		},
		{
			name:  "Сложное предложение",
			input: "I am go I!!!?? restricted: ;directory machine PC am ....I PC? go!?",
			expectedOutput: map[string]int{
				"i":          3,
				"am":         2,
				"go":         2,
				"restricted": 1,
				"directory":  1,
				"machine":    1,
				"pc":         2,
			},
		},
		{
			name:           "Знаки препинания и пустое слово", ///???
			input:          "...,?!;;",
			expectedOutput: map[string]int{},
		},
	}

	for _, tt := range tests {
		tt := tt // захват переменной цикла
		t.Run(tt.name, func(t *testing.T) {
			output := CountingWords(tt.input)
			if !reflect.DeepEqual(output, tt.expectedOutput) {
				t.Errorf("CountingWords(%v) = %v; want %v", tt.input, output, tt.expectedOutput)
			}
		})
	}
}
