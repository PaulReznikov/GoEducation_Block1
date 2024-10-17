package main

import (
	"reflect"
	"testing"
)

// TestWorkingWithAnagrams проверяет функцию WorkingWithAnagrams на различных наборах входных данных.
func TestWorkingWithAnagrams(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput map[string][]string
	}{
		{
			name:           "Пустая строка",
			input:          "",
			expectedOutput: map[string][]string{},
		},
		{
			name:           "Все слова с не-буквенными символами",
			input:          "hello! world@ 123",
			expectedOutput: map[string][]string{},
		},
		{
			name:  "Нет анаграмм, все уникальные слова",
			input: "cat dog bird fish",
			expectedOutput: map[string][]string{
				"act":  {"cat"},
				"dgo":  {"dog"},
				"bdir": {"bird"},
				"fhis": {"fish"},
			},
		},
		{
			name:  "Существуют анаграммы",
			input: "listen silent enlist inlets google gogole",
			expectedOutput: map[string][]string{
				"eilnst": {"listen", "silent", "enlist", "inlets"},
				"eggloo": {"google", "gogole"},
			},
		},
		{
			name:  "Дублирующиеся слова",
			input: "listen silent enlist inlets listen",
			expectedOutput: map[string][]string{
				"eilnst": {"listen", "silent", "enlist", "inlets"},
			},
		},
		{
			name:  "Слова с разным регистром",
			input: "Listen Silent Enlist Inlets listen",
			expectedOutput: map[string][]string{
				"eilnst": {"listen", "silent", "enlist", "inlets"},
			},
		},
		{
			name:  "Смешанные валидные и невалидные слова",
			input: "listen silent! enlist inlets gogole google123",
			expectedOutput: map[string][]string{
				"eggloo": {"gogole"},
				"eilnst": {"listen", "enlist", "inlets"},
			},
		},
		{
			name:  "Множество групп анаграмм",
			input: "rat tar art star tars",
			expectedOutput: map[string][]string{
				"art":  {"rat", "tar", "art"},
				"arst": {"star", "tars"},
			},
		},
		{
			name:  "Слова с одинаковыми буквами в разном порядке и регистре",
			input: "Dormitory Dirtyroom",
			expectedOutput: map[string][]string{
				"dimoorrty": {"dormitory", "dirtyroom"},
			},
		},
		{
			name:           "Слова с не-ASCII символами",
			input:          "café café café",
			expectedOutput: map[string][]string{}, // 'café' содержит 'é', не проходит regex
		},
		{
			name:  "Множество дубликатов и анаграмм",
			input: "listen silent enlist inlets listen silent",
			expectedOutput: map[string][]string{
				"eilnst": {"listen", "silent", "enlist", "inlets"},
			},
		},
		{
			name:  "Все слова являются анаграммами",
			input: "evil vile live veil",
			expectedOutput: map[string][]string{
				"eilv": {"evil", "vile", "live", "veil"},
			},
		},
		{
			name:  "Слова с внутренними символами и пунктуацией",
			input: "listen, silent! enlist; inlets gogole.",
			expectedOutput: map[string][]string{
				"eilnst": {"inlets"},
			},
		},
		{
			name:  "Слова с смешанными буквами и символами",
			input: "LiStEn SiLeNt EnLiSt InLeTs GogOlE",
			expectedOutput: map[string][]string{
				"eilnst": {"listen", "silent", "enlist", "inlets"},
				"eggloo": {"gogole"},
			},
		},
		{
			name:           "Строка содержит только пробелы",
			input:          "   ",
			expectedOutput: map[string][]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := WorkingWithAnagrams(tt.input)

			// Проверка на соответствие ожидаемого и фактического результата
			if !reflect.DeepEqual(output, tt.expectedOutput) {
				t.Errorf("WorkingWithAnagrams(%q) = %v; want %v", tt.input, output, tt.expectedOutput)
			}
		})
	}
}
