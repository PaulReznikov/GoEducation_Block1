package main

import (
	"testing"
)

// TestFilteringSortSearch проверяет корректность работы функции FilteringSortSearch.
func TestFilteringSortSearch(t *testing.T) {
	tests := []struct {
		name           string
		words          []string
		searchLen      int
		expectedOutput string
	}{
		{
			name:           "Строки с цифрами фильтруются, первая строка длины больше 5",
			words:          []string{"hhhh4", "a", "bbb", "c4", "cccc", "eeeeee", "fffffffffff"},
			searchLen:      5,
			expectedOutput: "eeeeee",
		},
		{
			name:           "Нет строки больше длины 10",
			words:          []string{"hhhh4", "a", "bbb", "c4", "cccc", "eeeeee", "fffffffffff"},
			searchLen:      10,
			expectedOutput: "fffffffffff",
		},
		{
			name:           "Все строки содержат цифры, пустой результат",
			words:          []string{"123", "h2ello", "go4lang"},
			searchLen:      2,
			expectedOutput: "",
		},
		{
			name:           "Нет строк с цифрами, сортировка по длине, длина больше 2",
			words:          []string{"go", "hello", "gopher"},
			searchLen:      2,
			expectedOutput: "hello",
		},
		{
			name:           "Пустой слайс, должен вернуть пустую строку",
			words:          []string{},
			searchLen:      5,
			expectedOutput: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := FilteringSortSearch(tt.words, tt.searchLen)
			if output != tt.expectedOutput {
				t.Errorf("FilteringSortSearch(%v, %d) = %q; want %q", tt.words, tt.searchLen, output, tt.expectedOutput)
			}
		})
	}
}
