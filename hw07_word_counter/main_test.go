package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountWords(t *testing.T) {
	tests := []struct {
		input    string
		expected map[string]int
	}{
		{"Hello, world! Hello everyone. This is a test.", map[string]int{"hello": 2, "world": 1, "everyone": 1, "this": 1, "is": 1, "a": 1, "test": 1}},
		{"Go is great! Go is fast.", map[string]int{"go": 2, "is": 2, "great": 1, "fast": 1}},
		{"", map[string]int{}},
		{"Go, Python, and Java are programming languages.", map[string]int{"go": 1, "python": 1, "and": 1, "java": 1, "are": 1, "programming": 1, "languages": 1}},
	}

	for _, test := range tests {
		output := countWords(test.input)
		assert.Equal(t, test.expected, output, "For input '%s', expected %v, but got %v", test.input, test.expected, output)
	}
}
