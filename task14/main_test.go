package main

import (
	"reflect"
	"testing"
)

// TestGrouppingSort проверяет функцию GrouppingSort на различных наборах входных данных.
func TestGrouppingSort(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput [][]string
	}{
		{
			name:           "Пустая строка",
			input:          "",
			expectedOutput: [][]string{},
		},
		{
			name:  "Одна группа слов одинаковой длины",
			input: "cat dog bat",
			expectedOutput: [][]string{
				{"cat", "dog", "bat"},
			},
		},
		{
			name:  "Разные длины слов",
			input: "a bb ccc dddd eeeee",
			expectedOutput: [][]string{
				{"a"},
				{"bb"},
				{"ccc"},
				{"dddd"},
				{"eeeee"},
			},
		},
		{
			name:  "Смешанные длины и повторяющиеся слова",
			input: "A bb ccc bb dddd ccc ccc eeeee",
			expectedOutput: [][]string{
				{"A"},
				{"bb", "bb"},
				{"ccc", "ccc", "ccc"},
				{"dddd"},
				{"eeeee"},
			},
		},
		{
			name:  "Слова с разными длинами, включая пробелы",
			input: "aaa  bb cc",
			expectedOutput: [][]string{
				{"bb", "cc"},
				{"aaa"},
			},
		},
		{
			name:  "Слова с одинаковой длиной и разным регистром",
			input: "cat CAT bat BAT",
			expectedOutput: [][]string{
				{"cat", "CAT", "bat", "BAT"},
			},
		},
		{
			name:  "Слова с неанглийскими символами",
			input: "кошка собака кот",
			expectedOutput: [][]string{
				{"кот"},
				{"кошка"},
				{"собака"},
			},
		},
	}

	for _, tt := range tests {
		tt := tt // захват переменной цикла
		t.Run(tt.name, func(t *testing.T) {
			output := GrouppingSort(tt.input)
			if !reflect.DeepEqual(output, tt.expectedOutput) {
				t.Errorf("GrouppingSort(%q) = %v; want %v", tt.input, output, tt.expectedOutput)
			}
		})
	}
}
