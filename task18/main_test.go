package main

import (
	"reflect"
	"testing"
)

func TestExtractionRevers(t *testing.T) {
	tests := []struct {
		testName string
		input    string
		expected string
	}{
		{
			testName: "Проверка на пустую строку",
			input:    "",
			expected: "",
		},
		{
			testName: "Проверка строки без уникальных символов",
			input:    "aabbccdd",
			expected: "ddccbbaa",
		},
		{
			testName: "Проверка строки из уникальных символов",
			input:    "abcd",
			expected: "",
		},
		{
			testName: "Проверка строки с разным регистром букв",
			input:    "AbAccB",
			expected: "ccAA",
		},
		{
			testName: "Проверка строки с пробелами",
			input:    "aa bb CC dd DD",
			expected: "DD dd CC bb aa",
		},
		{
			testName: "Проверка на цифры и нестандартные символы",
			input:    "23!**%((3%",
			expected: "%3((%**3",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.testName, func(t *testing.T) {
			result := ExtractionRevers(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ExtractionRevers(%v) = %v, want = %v", tt.input, result, tt.expected)
			}
		})

	}

}
