package main

import (
	"testing"
)

// TestSortStrings проверяет функцию SortStrings на различных наборах входных данных.
func TestSortStrings(t *testing.T) {
	tests := []struct {
		name           string
		input          []string
		expectedOutput string
	}{
		{
			name:           "Пустой слайс",
			input:          []string{},
			expectedOutput: "",
		},
		{
			name:           "Слайс без дубликатов, сортировка по длине",
			input:          []string{"apple", "dog", "banana", "cat"},
			expectedOutput: "dog, cat, apple, banana",
		},
		{
			name:           "Слайс с дубликатами",
			input:          []string{"apple", "dog", "banana", "dog", "cat", "apple"},
			expectedOutput: "dog, cat, apple, banana",
		},
		{
			name:           "Все элементы одинаковые",
			input:          []string{"test", "test", "test"},
			expectedOutput: "test",
		},
		{
			name:           "Слайс с разной длиной строк",
			input:          []string{"a", "ab", "abc", "abcd"},
			expectedOutput: "a, ab, abc, abcd",
		},
		{
			name:           "Слайс с пустыми строками",
			input:          []string{"", "a", "", "ab"},
			expectedOutput: ", a, ab",
		},
		{
			name:           "Слайс с unicode символами",
			input:          []string{"😊", "hello", "привет", "world", "😊"},
			expectedOutput: "😊, hello, world, привет",
		},
		{
			name:           "Слайс с разными регистрами",
			input:          []string{"Apple", "apple", "Banana", "banana"},
			expectedOutput: "Apple, apple, Banana, banana",
		},
		{
			name:           "Слайс с одними короткими строками",
			input:          []string{"a", "b", "c"},
			expectedOutput: "a, b, c",
		},
		{
			name:           "Слайс с одним элементом",
			input:          []string{"unique"},
			expectedOutput: "unique",
		},
		{
			name:           "Слайс с пустой строкой",
			input:          []string{""},
			expectedOutput: "",
		},
		{
			name:           "Слайс с несколькими пустыми строками",
			input:          []string{"", "", ""},
			expectedOutput: "",
		},
		{
			name:           "Слайс с дубликатами и разной длиной",
			input:          []string{"a", "aa", "a", "aaa", "aa"},
			expectedOutput: "a, aa, aaa",
		},
		{
			name:           "Слайс с пробелами внутри строк",
			input:          []string{"a b", "b a", "c", "c"},
			expectedOutput: "c, a b, b a",
		},
		{
			name:           "Слайс с различными символами",
			input:          []string{"hello!", "world@", "hello", "world"},
			expectedOutput: "hello, world, hello!, world@",
		},
		{
			name:           "Слайс с цифрами в строках",
			input:          []string{"123", "1", "12", "123"},
			expectedOutput: "1, 12, 123",
		},
		//???
		//{
		//	name:           "Слайс с смешанными символами",
		//	input:          []string{"Go1", "Go2", "Go1", "Go3"},
		//	expectedOutput: "Go1, Go2, Go3",
		//},
	}

	for _, tt := range tests {
		tt := tt // захват переменной цикла
		t.Run(tt.name, func(t *testing.T) {
			output := SortStrings(tt.input)
			if output != tt.expectedOutput {
				t.Errorf("SortStrings(%v) = %q; want %q", tt.input, output, tt.expectedOutput)
			}
		})
	}
}
