package main

import (
	"reflect"
	"testing"
)

// TestSort2DSlice проверяет функцию Sort2DSlice на различных наборах входных данных.
func TestSort2DSlice(t *testing.T) {
	tests := []struct {
		name           string
		input          [][]int
		expectedOutput []int
		expectPanic    bool
	}{
		{
			name:           "Пустой внешний слайс",
			input:          [][]int{},
			expectedOutput: []int{},
			expectPanic:    false,
		},
		{
			name:           "Внешний слайс с пустыми внутренними слайсами",
			input:          [][]int{{}, {}, {}},
			expectedOutput: nil, // Функция вызовет панику, так как пытается получить slc2D[i][0]
			expectPanic:    true,
		},
		{
			name: "Один внутренний слайс с одним элементом",
			input: [][]int{
				{5},
			},
			expectedOutput: []int{5},
			expectPanic:    false,
		},
		{
			name: "Несколько внутренних слайсов с разным количеством элементов",
			input: [][]int{
				{1, 3, 2},
				{4, 4, 4, 4},
				{0, -1, -2},
			},
			expectedOutput: []int{0, 3, 4},
			expectPanic:    false,
		},
		{
			name: "Все внутренние слайсы содержат одинаковые элементы",
			input: [][]int{
				{2, 2, 2},
				{2, 2},
				{2},
			},
			expectedOutput: []int{2, 2, 2},
			expectPanic:    false,
		},
		{
			name: "Внутренние слайсы с отрицательными числами",
			input: [][]int{
				{-3, -1, -7},
				{-2, -2, -2},
				{-5, -4},
			},
			expectedOutput: []int{-4, -2, -1},
			expectPanic:    false,
		},
		{
			name: "Внутренние слайсы с положительными и отрицательными числами",
			input: [][]int{
				{-1, 0, 1},
				{2, -2, 3},
				{4, -4, 5},
			},
			expectedOutput: []int{1, 3, 5},
			expectPanic:    false,
		},
		{
			name: "Внутренние слайсы с одинаковыми максимумами",
			input: [][]int{
				{1, 2, 3},
				{3, 3, 3},
				{2, 3},
			},
			expectedOutput: []int{3, 3, 3},
			expectPanic:    false,
		},
		{
			name: "Внутренний слайс с большим числом",
			input: [][]int{
				{1, 2, 999999},
				{7, 8, 9},
			},
			expectedOutput: []int{9, 999999},
			expectPanic:    false,
		},
		{
			name: "Внутренний слайс с нулем",
			input: [][]int{
				{0, 0, 0},
				{0},
			},
			expectedOutput: []int{0, 0},
			expectPanic:    false,
		},
		{
			name: "Внутренние слайсы с одним элементом каждый",
			input: [][]int{
				{10},
				{20},
				{30},
			},
			expectedOutput: []int{10, 20, 30},
			expectPanic:    false,
		},
		{
			name: "Внутренние слайсы с возрастающими числами",
			input: [][]int{
				{1, 2, 3, 4},
				{2, 3, 4, 5},
				{3, 4, 5, 6},
			},
			expectedOutput: []int{4, 5, 6},
			expectPanic:    false,
		},
		{
			name: "Внутренние слайсы с убывающими числами",
			input: [][]int{
				{4, 3, 2, 1},
				{5, 4, 3, 2},
				{6, 5, 4, 3},
			},
			expectedOutput: []int{4, 5, 6},
			expectPanic:    false,
		},
		{
			name: "Внутренние слайсы с повторяющимися максимумами",
			input: [][]int{
				{1, 3, 3},
				{3, 3, 1},
				{1, 3, 3},
			},
			expectedOutput: []int{3, 3, 3},
			expectPanic:    false,
		},
		{
			name: "Внутренние слайсы с разными типами чисел",
			input: [][]int{
				{1, -1, 2},
				{3, -3, 4},
				{-5, 5, 6},
			},
			expectedOutput: []int{2, 4, 6},
			expectPanic:    false,
		},
	}

	for _, tt := range tests {
		tt := tt // capture range variable
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if !tt.expectPanic {
						t.Errorf("Sort2DSlice(%v) panicked unexpectedly: %v", tt.input, r)
					}
				}
			}()

			output := Sort2DSlice(tt.input)
			if !tt.expectPanic {
				// Для корректного сравнения срезов используем reflect.DeepEqual
				if !reflect.DeepEqual(output, tt.expectedOutput) {
					t.Errorf("Sort2DSlice(%v) = %v; want %v", tt.input, output, tt.expectedOutput)
				}
			}
		})
	}
}
