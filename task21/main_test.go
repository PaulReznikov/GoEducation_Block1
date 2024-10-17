package main

import (
	"reflect"
	"testing"
)

// TestModificationAndFiltering проверяет корректность работы функции ModificationAndFiltering.
func TestModificationAndFiltering(t *testing.T) {
	tests := []struct {
		name           string
		input          []int
		expectedOutput []int
	}{
		{
			name:           "Пустой слайс",
			input:          []int{},
			expectedOutput: []int{},
		},
		{
			name:           "Один элемент, меньше суммы",
			input:          []int{5},
			expectedOutput: []int{},
		},
		{
			name:           "Один элемент, больше суммы",
			input:          []int{0},
			expectedOutput: []int{},
		},
		{
			name:           "Несколько элементов, все остаются после фильтрации",
			input:          []int{1, 3, 5},
			expectedOutput: []int{10},
		},
		{
			name:           "Несколько элементов, часть удаляется",
			input:          []int{1, 2, 3, 4, 5},
			expectedOutput: []int{20},
		},
		{
			name:           "Все элементы равны",
			input:          []int{2, 2, 2, 2},
			expectedOutput: []int{},
		},
		{
			name:           "Слайс с отрицательными и нулевыми элементами",
			input:          []int{0, -1, 2, 3, -4},
			expectedOutput: []int{4, 9},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			output := ModificationAndFiltering(tt.input)
			if !reflect.DeepEqual(output, tt.expectedOutput) {
				t.Errorf("ModificationAndFiltering(%v) = %v; want %v", tt.input, output, tt.expectedOutput)
			}
		})
	}
}
