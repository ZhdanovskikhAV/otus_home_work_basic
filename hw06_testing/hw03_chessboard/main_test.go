package main

import (
	"testing"
)

func TestGenerateRow(t *testing.T) {
	tests := []struct {
		row      int
		size     int
		expected string
	}{
		{0, 5, "# # #"},
		{1, 5, " # # "},
		{2, 4, "# # "},
		{3, 4, " # #"},
	}

	for _, test := range tests {
		result := generateRow(test.row, test.size)
		if result != test.expected {
			t.Errorf("generateRow(%d, %d) = %q; want %q", test.row, test.size, result, test.expected)
		}
	}
}
