package main

import (
	"reflect"
	"testing"
)

// TestDeletingAndConvertingElements проверяет функцию DeletingAndConvertingElements на различных наборах данных.
func TestDeletingAndConvertingElements(t *testing.T) {
	tests := []struct {
		name           string
		input          []string
		expectedOutput []string
	}{
		{
			name:           "Пустой слайс", //???
			input:          []string{},
			expectedOutput: []string{},
		},
		{
			name:           "Все строки короче 5 символов", //???
			input:          []string{"a", "bb", "ccc", "dddd", "ffff"},
			expectedOutput: []string{},
		},
		{
			name:           "Строки длиннее 5 символов",
			input:          []string{"aaa", "bbbbbbbb", "c", "ddddd", "eeeeeeeeeeeeeeeee", "ffff"},
			expectedOutput: []string{"EEEEEEEEEEEEEEEEE", "DDDDD", "BBBBBBBB"},
		},
		{
			name:           "Только одна строка длиннее 5 символов",
			input:          []string{"short", "longerthanfive"},
			expectedOutput: []string{"LONGERTHANFIVE", "SHORT"},
		},
		{
			name:           "Все строки длиннее 5 символов",
			input:          []string{"sixsix", "sevensven", "eightsateight"},
			expectedOutput: []string{"EIGHTSATEIGHT", "SEVENSVEN", "SIXSIX"},
		},
		{
			name:           "Смешанные длины строк и разный регистр",
			input:          []string{"short", "ALREADYUPPER", "lowercase", "MixedCase", "tiny"},
			expectedOutput: []string{"MIXEDCASE", "LOWERCASE", "ALREADYUPPER", "SHORT"},
		},
	}

	for _, tt := range tests {
		tt := tt // захват переменной цикла
		t.Run(tt.name, func(t *testing.T) {
			output := DeletingAndConvertingElements(tt.input)
			if !reflect.DeepEqual(output, tt.expectedOutput) {
				t.Errorf("DeletingAndConvertingElements(%v) = %v; want %v", tt.input, output, tt.expectedOutput)
			}
		})
	}
}
