package main

import (
	"testing"
)

// TestCompressString проверяет функцию CompressString на различных наборах входных данных.
func TestCompressString(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput string
	}{
		{
			name:           "Empty string",
			input:          "",
			expectedOutput: "",
		},
		{
			name:           "Single character",
			input:          "a",
			expectedOutput: "a",
		},
		{
			name:           "No repeating characters",
			input:          "abcd",
			expectedOutput: "abcd",
		},
		{
			name:           "All characters the same",
			input:          "aaaaa",
			expectedOutput: "a5",
		},
		{
			name:           "Repeating characters with counts",
			input:          "aaabbcccc",
			expectedOutput: "a3b2c4",
		},
		{
			name:           "Repeating characters at the end",
			input:          "aabbbb",
			expectedOutput: "a2b4",
		},
		{
			name:           "Repeating characters not at the end",
			input:          "aabcccd",
			expectedOutput: "a2bc3d",
		},
		{
			name:           "Unicode characters",
			input:          "hééllo",
			expectedOutput: "hé2l2o",
		},
		{
			name:           "Compressed string not shorter",
			input:          "abcd",
			expectedOutput: "abcd",
		},
		{
			name:           "Compressed string equal in length",
			input:          "aabb",
			expectedOutput: "a2b2",
		},
		{
			name:           "Compressed string longer than original",
			input:          "abc",
			expectedOutput: "abc",
		},
		{
			name:           "Alternating repeating characters",
			input:          "ababab",
			expectedOutput: "ababab",
		},
		{
			name:           "Mixed repeating and non-repeating characters",
			input:          "aaabbaa",
			expectedOutput: "a3b2a2",
		},
		{
			name:           "Single repeating character",
			input:          "bb",
			expectedOutput: "b2",
		},
		{
			name:           "Multiple single characters",
			input:          "aabccbaa",
			expectedOutput: "a2bc2ba2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := CompressString(tt.input)
			if output != tt.expectedOutput {
				t.Errorf("CompressString(%q) = %q; want %q", tt.input, output, tt.expectedOutput)
			}
		})
	}
}
