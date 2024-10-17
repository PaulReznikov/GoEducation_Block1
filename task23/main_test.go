package main

import (
	"testing"
)

// TestReverseWords проверяет корректность работы функции ReverseWords.
func TestReverseWords(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput string
	}{
		{
			name:           "Пустая строка",
			input:          "",
			expectedOutput: "",
		},
		{
			name:           "Одно слово",
			input:          "golang",
			expectedOutput: "gnalog",
		},
		{
			name:           "Два слова",
			input:          "hello world",
			expectedOutput: "dlrow olleh",
		},
		{
			name:           "Несколько слов с пробелами",
			input:          "The golang code",
			expectedOutput: "edoc gnalog ehT",
		},
		{
			name:           "Строка с несколькими пробелами между словами",
			input:          "hello    world golang",
			expectedOutput: "gnalog dlrow    olleh",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := ReverseWords(tt.input)
			if output != tt.expectedOutput {
				t.Errorf("ReverseWords(%q) = %q; want %q", tt.input, output, tt.expectedOutput)
			}
		})
	}
}
