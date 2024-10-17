package main

import (
	"testing"
)

// TestExtractingCoutingChar проверяет функцию ExtractingCoutingChar на различных наборах входных данных.
func TestExtractingCoutingChar(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput int
	}{
		{
			name:           "Пустая строка",
			input:          "",
			expectedOutput: 0,
		},
		{
			name:           "Строка без цифр",
			input:          "abcdefg",
			expectedOutput: 0,
		},
		{
			name:           "Строка только с цифрами",
			input:          "12345",
			expectedOutput: 15,
		},
		{
			name:           "Смешанная строка с цифрами и буквами",
			input:          "a1b2c3",
			expectedOutput: 6,
		},
		{
			name:           "Строка с ведущими нулями",
			input:          "00123",
			expectedOutput: 6,
		},
		{
			name:           "Строка с дубликатами цифр",
			input:          "112233",
			expectedOutput: 12,
		},
		{
			name:           "Строка с пробелами и цифрами",
			input:          "a 1 b 2 c 3",
			expectedOutput: 6,
		},
		{
			name:           "Строка с специальными символами и цифрами",
			input:          "!@#1$%^2&*()3",
			expectedOutput: 6,
		},
		{
			name:           "Строка с Unicode цифрами (арабские индикаторы)",
			input:          "a١b٢c٣", // '١', '٢', '٣' не соответствуют regex `^[0-9]$`
			expectedOutput: 0,
		},
		{
			name:           "Строка с разными регистрами букв и цифрами",
			input:          "A1a2B3b4",
			expectedOutput: 10,
		},
		{
			name:           "Строка с цифрами и пробелами",
			input:          "1 2 3 4 5",
			expectedOutput: 15,
		},
		{
			name:           "Строка с цифрами и новыми строками",
			input:          "1\n2\t3\r4",
			expectedOutput: 10,
		},
		{
			name:           "Строка с очень большими цифрами",
			input:          "99999",
			expectedOutput: 45,
		},
		{
			name:           "Строка с нулями",
			input:          "0a0b0c",
			expectedOutput: 0,
		},
		{
			name:           "Строка с нецифровыми числами",
			input:          "a12b34c56", // Функция считает только одиночные цифры '1','2','3','4','5','6'
			expectedOutput: 21,
		},
		{
			name:           "Строка с смешанными цифрами и Unicode символами",
			input:          "Go1😊Go2🚀",
			expectedOutput: 3,
		},
	}

	for _, tt := range tests {
		tt := tt // захват переменной цикла
		t.Run(tt.name, func(t *testing.T) {
			output := ExtractingCoutingChar(tt.input)
			if output != tt.expectedOutput {
				t.Errorf("ExtractingCoutingChar(%q) = %d; want %d", tt.input, output, tt.expectedOutput)
			}
		})
	}
}
