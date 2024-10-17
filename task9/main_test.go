package main

import (
	"reflect"
	"testing"
)

// TestMergeHashTable проверяет функцию MergeHashTable на различных наборах входных данных.
func TestMergeHashTable(t *testing.T) {
	tests := []struct {
		name           string
		inputM1        map[string]int
		inputM2        map[string]int
		expectedOutput map[string]int
		expectPanic    bool
	}{
		{
			name:           "Обе карты пусты",
			inputM1:        map[string]int{},
			inputM2:        map[string]int{},
			expectedOutput: map[string]int{},
			expectPanic:    false,
		},
		{
			name:           "m1 пустая, m2 непустая",
			inputM1:        map[string]int{},
			inputM2:        map[string]int{"x": 5, "y": 15},
			expectedOutput: map[string]int{"x": 5, "y": 15},
			expectPanic:    false,
		},
		{
			name:           "m1 непустая, m2 пустая",
			inputM1:        map[string]int{"x": 10, "y": 20},
			inputM2:        map[string]int{},
			expectedOutput: map[string]int{"x": 10, "y": 20},
			expectPanic:    false,
		},
		{
			name:           "Нет общих ключей",
			inputM1:        map[string]int{"a": 5, "b": 6},
			inputM2:        map[string]int{"c": 7, "d": 8},
			expectedOutput: map[string]int{"a": 5, "b": 6, "c": 7, "d": 8},
			expectPanic:    false,
		},
		{
			name:           "Общие ключи, сумма >10",
			inputM1:        map[string]int{"a": 7, "b": 5},
			inputM2:        map[string]int{"a": 5, "b": 6},
			expectedOutput: map[string]int{"a": 12, "b": 11},
			expectPanic:    false,
		},
		{
			name:           "Общие ключи, сумма <=10",
			inputM1:        map[string]int{"a": 3, "b": 4},
			inputM2:        map[string]int{"a": 5, "b": 6},
			expectedOutput: map[string]int{},
			expectPanic:    false,
		},
		{
			name: "Общие ключи с различными суммами",
			inputM1: map[string]int{
				"a": 7, // 7 + 4 = 11 >10
				"b": 2, // 2 +3 =5 <=10
				"d": 6, //6 +20=26 >10
			},
			inputM2: map[string]int{
				"YYYYYYYY": 1,
				"a":        4,
				"b":        3,
				"d":        20,
				"QQ":       5,
			},
			expectedOutput: map[string]int{
				"YYYYYYYY": 1,
				"a":        11,
				"d":        26,
				"QQ":       5,
			},
			expectPanic: false,
		},
		{
			name:           "m2 изначально nil и m1 тоже пустая",
			inputM1:        map[string]int{},
			inputM2:        nil,
			expectedOutput: nil,
			expectPanic:    false,
		},
		{
			name:           "m2 изначально nil и m1 имеет элементы",
			inputM1:        map[string]int{"a": 5},
			inputM2:        nil,
			expectedOutput: map[string]int{"a": 5},
			expectPanic:    true, // Поскольку m2 nil и пытаемся записать в него
		},
		{
			name: "Общие ключи с суммой ровно 10",
			inputM1: map[string]int{
				"a": 5,
				"b": 5,
			},
			inputM2: map[string]int{
				"a": 5,
				"b": 5,
			},
			expectedOutput: map[string]int{},
			expectPanic:    false,
		},
		{
			name: "Разные типы значений",
			inputM1: map[string]int{
				"a": 15,
				"b": 20,
			},
			inputM2: map[string]int{
				"a": 10,
				"c": 25,
			},
			expectedOutput: map[string]int{
				"a": 25, // 15+10=25 >10
				"b": 20,
				"c": 25,
			},
			expectPanic: false,
		},
		{
			name: "Ключи с нулевыми значениями",
			inputM1: map[string]int{
				"a": 0,
				"b": 11,
			},
			inputM2: map[string]int{
				"a": 0,
				"b": 0,
			},
			expectedOutput: map[string]int{
				"b": 11, // "a" удаляется, "b" обновляется до 11
			},
			expectPanic: false,
		},
		{
			name: "Ключи только в m1",
			inputM1: map[string]int{
				"a": 12,
				"b": 15,
			},
			inputM2: map[string]int{
				"c": 5,
				"d": 8,
			},
			expectedOutput: map[string]int{
				"c": 5,
				"d": 8,
				"a": 12,
				"b": 15,
			},
			expectPanic: false,
		},
		{
			name:    "Ключи только в m2",
			inputM1: map[string]int{},
			inputM2: map[string]int{
				"c": 5,
				"d": 8,
			},
			expectedOutput: map[string]int{
				"c": 5,
				"d": 8,
			},
			expectPanic: false,
		},
		{
			name: "Большие значения",
			inputM1: map[string]int{
				"a": 1000,
				"b": 2000,
			},
			inputM2: map[string]int{
				"a": 5000,
				"b": 8000,
			},
			expectedOutput: map[string]int{
				"a": 6000,  // 1000 +5000=6000 >10
				"b": 10000, //2000 +8000=10000 >10
			},
			expectPanic: false,
		},
	}

	for _, tt := range tests {
		tt := tt // захват переменной цикла
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if !tt.expectPanic {
						t.Errorf("MergeHashTable(%v, %v) panicked unexpectedly: %v", tt.inputM1, tt.inputM2, r)
					}
				}
			}()

			output := MergeHashTable(tt.inputM1, tt.inputM2)

			if !tt.expectPanic {
				// Для корректного сравнения карт используем reflect.DeepEqual
				if !reflect.DeepEqual(output, tt.expectedOutput) {
					t.Errorf("MergeHashTable(%v, %v) = %v; want %v", tt.inputM1, tt.inputM2, output, tt.expectedOutput)
				}
			}
		})
	}
}
