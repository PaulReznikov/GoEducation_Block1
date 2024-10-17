package main

import (
	"reflect"
	"testing"
)

// TestAccumulationOfResults проверяет функцию AccumulationOfResults на различных наборах данных.
func TestAccumulationOfResults(t *testing.T) {
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
			name:           "Слайс с одним элементом",
			input:          []int{5},
			expectedOutput: []int{5}, // один элемент остается, так как нет среднего
		},
		{
			name:           "Слайс с положительными числами",
			input:          []int{1, 2, 3, 4, 5, 6, 7},
			expectedOutput: []int{24, 48}, // элементы больше среднего
		},
		{
			name:           "Слайс с отрицательными числами",
			input:          []int{-1, -2, -3, -4, -5},
			expectedOutput: []int{-1, -2, -3}, // один элемент остаётся, так как он больше среднего
		},
		{
			name:           "Смешанный слайс (положительные и отрицательные числа)",
			input:          []int{-3, 1, 4, -2, 7, -1},
			expectedOutput: []int{-3, 1, -2, -4}, // только один элемент превышает среднее значение
		},
		{
			name:           "Слайс из одинаковых чисел",
			input:          []int{5, 5, 5, 5, 5},
			expectedOutput: []int{20, 40}, // некоторые элементы будут оставлены, так как все элементы равны
		},
	}

	for _, tt := range tests {
		tt := tt // захват переменной цикла
		t.Run(tt.name, func(t *testing.T) {
			output := AccumulationOfResults(tt.input)
			if !reflect.DeepEqual(output, tt.expectedOutput) {
				t.Errorf("AccumulationOfResults(%v) = %v; want %v", tt.input, output, tt.expectedOutput)
			}
		})
	}
}
