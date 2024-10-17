package main

import (
	"reflect"
	"testing"
)

// TestMergeSlicesStrings проверяет функцию MergeSlicesStrings на различных наборах входных данных.
func TestMergeSlicesStrings(t *testing.T) {
	tests := []struct {
		name           string
		input1         []string
		input2         []string
		expectedOutput []string
	}{
		{
			name:           "Оба слайса пустые", //???
			input1:         []string{},
			input2:         []string{},
			expectedOutput: []string{},
		},
		{
			name:           "Первый слайс пустой, второй не пустой",
			input1:         []string{},
			input2:         []string{"a", "b", "c"},
			expectedOutput: []string{"a", "b", "c"},
		},
		{
			name:           "Первый слайс не пустой, второй пустой",
			input1:         []string{"x", "y", "z"},
			input2:         []string{},
			expectedOutput: []string{"x", "y", "z"},
		},
		{
			name:           "Нет пересечения, оба слайса уникальны",
			input1:         []string{"a", "b", "c"},
			input2:         []string{"d", "e", "f"},
			expectedOutput: []string{"a", "b", "c", "d", "e", "f"},
		},
		{
			name:           "Пересечение элементов",
			input1:         []string{"a", "b", "c"},
			input2:         []string{"b", "c", "d"},
			expectedOutput: []string{"a", "d"},
		},
		{
			name:           "Повторяющиеся элементы в слайсах",
			input1:         []string{"a", "b", "c", "c"},
			input2:         []string{"b", "d", "e", "d"},
			expectedOutput: []string{"a", "c", "e"},
		},
		{
			name:           "Смешанные регистры", //???
			input1:         []string{"A", "b", "C"},
			input2:         []string{"a", "B", "c"},
			expectedOutput: []string{},
		},
		{
			name:           "Списки содержат только пересекающиеся элементы", //???
			input1:         []string{"x", "y"},
			input2:         []string{"x", "y"},
			expectedOutput: []string{},
		},
		{
			name:           "Смешанные элементы, включая пробелы",
			input1:         []string{"  ", "a", "b"},
			input2:         []string{"b", "  ", "c"},
			expectedOutput: []string{"a", "c"},
		},
		{
			name:           "Списки с разными типами строк",
			input1:         []string{"aaa", "B", "advR", "bb", "AdvR"},
			input2:         []string{"XX", "aaa", "cc", "bb"},
			expectedOutput: []string{"advr", "b", "cc", "xx"},
		},
	}

	for _, tt := range tests {
		tt := tt // захват переменной цикла
		t.Run(tt.name, func(t *testing.T) {
			output := MergeSlicesStrings(tt.input1, tt.input2)
			if !reflect.DeepEqual(output, tt.expectedOutput) {
				t.Errorf("MergeSlicesStrings(%v, %v) = %v; want %v", tt.input1, tt.input2, output, tt.expectedOutput)
			}
		})
	}
}
