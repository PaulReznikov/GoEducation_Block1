package main

import "testing"

func TestFilteringAndSumm(t *testing.T) {
	tests := []struct {
		name           string
		input          []int
		expectedOutput int
	}{
		{
			name:           "Пустой слайс",
			input:          []int{},
			expectedOutput: 0,
		},
		{
			name:           "Слайс без отрицательных чисел",
			input:          []int{1, 2, 3},
			expectedOutput: 1*1 + 2*2 + 3*3, // 14
		},
		{
			name:           "Слайс с отрицательными числами",
			input:          []int{1, -2, 3, -4, 5},
			expectedOutput: 1*1 + 3*3 + 5*5, // 35
		},
		{
			name:           "Слайс, содержащий только отрицательные числа",
			input:          []int{-1, -2, -3, -4},
			expectedOutput: 0,
		},
		{
			name:           "Слайс с одним числом",
			input:          []int{4},
			expectedOutput: 4 * 4, // 16
		},
		{
			name:           "Слайс с одним отрицательным числом",
			input:          []int{-4},
			expectedOutput: 0,
		},
		{
			name:           "Слайс с нулями и отрицательными числами",
			input:          []int{-1, 0, -3, 0, -4},
			expectedOutput: 0 + 0, // 0
		},
		{
			name:           "Слайс с положительными и отрицательными числами и нулями",
			input:          []int{-1, 0, 2, -3, 0, 4, -5},
			expectedOutput: 2*2 + 4*4, // 20
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := FilteringAndSumm(tt.input)
			if output != tt.expectedOutput {
				t.Errorf("FilteringAndSumm(%v) = %v; expected %v", tt.input, output, tt.expectedOutput)
			}
		})
	}
}
