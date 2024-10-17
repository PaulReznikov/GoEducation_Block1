package main

import (
	"reflect"
	"testing"
)

// TestSliceOptimization проверяет корректность работы функции SliceOptimization.
func TestSliceOptimization(t *testing.T) {
	tests := []struct {
		name           string
		input          []int
		expectedOutput []int
	}{
		{
			name:           "Слайс с четными и нечетными индексами",
			input:          []int{1, 25, 30, 5, 60, 35},
			expectedOutput: []int{120, 60, 2},
		},
		{
			name:           "Пустой слайс",
			input:          []int{},
			expectedOutput: []int{},
		},
		{
			name:           "Слайс с одним элементом",
			input:          []int{10},
			expectedOutput: []int{20},
		},
		{
			name:           "Все элементы одинаковые",
			input:          []int{5, 5, 5, 5, 5, 5},
			expectedOutput: []int{10, 10, 10},
		},
		{
			name:           "Слайс с отрицательными числами",
			input:          []int{-2, -3, -4, -1, 0, -6},
			expectedOutput: []int{0, -4, -8},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := SliceOptimization(tt.input)
			if !reflect.DeepEqual(output, tt.expectedOutput) {
				t.Errorf("SliceOptimization(%v) = %v; want %v", tt.input, output, tt.expectedOutput)
			}
		})
	}
}
