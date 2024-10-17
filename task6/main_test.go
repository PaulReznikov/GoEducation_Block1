package main

import (
	"reflect"
	"testing"
)

// TestPrivateDictionary проверяет функцию PrivateDictionary на различных наборах входных данных.
func TestPrivateDictionary(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]int
	}{
		{
			name:     "Пустая строка",
			input:    "",
			expected: map[string]int{},
		},
		{
			name:     "Все слова короче 3 символов",
			input:    "a an by",
			expected: map[string]int{},
		},
		{
			name:     "Нет слов короче 3 символов",
			input:    "cat dog mouse",
			expected: map[string]int{"cat": 1, "dog": 1, "mouse": 1},
		},
		{
			name:     "Смешанные слова с разной длиной",
			input:    "hi cat dog a an mouse",
			expected: map[string]int{"cat": 1, "dog": 1, "mouse": 1},
		},
		{
			name:     "Слова ровно 3 символа",
			input:    "cat dog bat",
			expected: map[string]int{"cat": 1, "dog": 1, "bat": 1},
		},
		{
			name:     "Повторяющиеся слова",
			input:    "cat dog cat mouse dog dog",
			expected: map[string]int{"cat": 2, "dog": 3, "mouse": 1},
		},
		{
			name:     "Юникодные символы",
			input:    "привет мир мир до",
			expected: map[string]int{"привет": 1, "мир": 2},
		},
		{
			name:     "Множественные пробелы",
			input:    "cat    dog  mouse",
			expected: map[string]int{"cat": 1, "dog": 1, "mouse": 1},
		},
		{
			name:     "Одно слово длиннее 3 символов",
			input:    "elephant",
			expected: map[string]int{"elephant": 1},
		},
		{
			name:     "Одно слово короче 3 символов",
			input:    "hi",
			expected: map[string]int{},
		},
		{
			name:     "Слова с пунктуацией",
			input:    "hello, world! hi.",
			expected: map[string]int{"hello,": 1, "world!": 1, "hi.": 1},
		},
		{
			name:     "Слова ровно 3 и более символов",
			input:    "cat dog bat elephant",
			expected: map[string]int{"cat": 1, "dog": 1, "bat": 1, "elephant": 1},
		},
		{
			name:     "Слова с разным регистром",
			input:    "Cat cat CAT dog DOG",
			expected: map[string]int{"Cat": 1, "cat": 1, "CAT": 1, "dog": 1, "DOG": 1},
		},
		{
			name:     "Слова с не-ASCII символами",
			input:    "café café café",
			expected: map[string]int{"café": 3},
		},
		{
			name:     "Смешанные юникодные и ASCII символы",
			input:    "hello привет hello мир",
			expected: map[string]int{"hello": 2, "привет": 1, "мир": 1},
		},
		{
			name:     "Слова с длиной ровно 2 символа",
			input:    "a an be",
			expected: map[string]int{},
		},
		{
			name:     "Слова с длиной ровно 3 символа и более",
			input:    "see bee fee",
			expected: map[string]int{"see": 1, "bee": 1, "fee": 1},
		},
		{
			name:     "Слова с смешанной длиной и повторениями",
			input:    "apple a ap app apple apple a",
			expected: map[string]int{"apple": 3, "app": 1},
		},
		{
			name:     "Слова с смешанными пробелами и табуляциями",
			input:    "cat\tdog  mouse\nelephant",
			expected: map[string]int{"cat": 1, "dog": 1, "mouse": 1, "elephant": 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := PrivateDictionary(tt.input)
			if !reflect.DeepEqual(output, tt.expected) {
				t.Errorf("PrivateDictionary(%q) = %v; want %v", tt.input, output, tt.expected)
			}
		})
	}
}
