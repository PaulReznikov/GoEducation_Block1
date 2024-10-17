package main

import (
	"reflect"
	"testing"
)

// TestGroupAnagrams проверяет функцию GroupAnagrams на различных наборах входных данных.
func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected map[string][]string
	}{
		{
			name:  "Basic case with multiple anagrams",
			input: []string{"listen", "silent", "enlist", "inlets", "google", "gogole"},
			expected: map[string][]string{
				"eilnst": {"listen", "silent", "enlist", "inlets"},
				"eggloo": {"google", "gogole"},
			},
		},
		{
			name:  "Case with invalid strings and duplicates",
			input: []string{"listen", "siL11ent", "enlist", "inlets", "goOgle%", "gogole", "listen"},
			expected: map[string][]string{
				"eilnst": {"listen", "enlist", "inlets"},
				"eggloo": {"gogole"},
			},
		},
		{
			name:     "Empty input slice",
			input:    []string{},
			expected: map[string][]string{},
		},
		{
			name:  "All unique words with no anagrams",
			input: []string{"apple", "banana", "carrot"},
			expected: map[string][]string{
				"aelpp":  {"apple"},
				"aaabnn": {"banana"},
				"acorrt": {"carrot"},
			},
		},
		{
			name:  "All words are anagrams of each other",
			input: []string{"abc", "bca", "cab", "cba", "bac", "acb"},
			expected: map[string][]string{
				"abc": {"abc", "bca", "cab", "cba", "bac", "acb"},
			},
		},
		{
			name:  "Mixed case sensitivity",
			input: []string{"Listen", "Silent", "Enlist", "Inlets", "Google", "Gogole"},
			expected: map[string][]string{
				"eilnst": {"listen", "silent", "enlist", "inlets"},
				"eggloo": {"google", "gogole"},
			},
		},
		{
			name:     "Strings with non-alphabetic characters only",
			input:    []string{"123", "!!!", "@@@"},
			expected: map[string][]string{},
		},
		{
			name:  "Mixed valid and invalid strings",
			input: []string{"race", "care", "acre", "race1", "care!", "acer"},
			expected: map[string][]string{
				"acer": {"race", "care", "acre", "acer"},
			},
		},
		{
			name:  "Strings with empty strings",
			input: []string{"", "a", "b", ""},
			expected: map[string][]string{
				"a": {"a"},
				"b": {"b"},
			},
		},
		{
			name:  "Strings with uppercase and lowercase duplicates",
			input: []string{"Listen", "Silent", "LISTEN", "silent"},
			expected: map[string][]string{
				"eilnst": {"listen", "silent"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GroupAnagrams(tt.input)

			// Для корректного сравнения карт используем reflect.DeepEqual
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("GroupAnagrams(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}
