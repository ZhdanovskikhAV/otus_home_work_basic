package main

import (
	"strings"
	"testing"
)

func generateBoard(size int) string {
	var sb strings.Builder
	for row := 0; row < size; row++ {
		sb.WriteString(generateRow(row, size))
		sb.WriteString("\n") // Добавьте новую строку между рядами
	}
	return sb.String()
}

func TestGenerateBoard(t *testing.T) {
	expectedBoard := "# #\n # \n# #\n"
	board := generateBoard(3)

	if board != expectedBoard {
		t.Errorf("Expected:\n%sGot:\n%s", expectedBoard, board)
	}

	expectedBoard = "# # \n # #\n# # \n # #\n"
	board = generateBoard(4)

	if board != expectedBoard {
		t.Errorf("Expected:\n%sGot:\n%s", expectedBoard, board)
	}
}
