package main

import (
	"strings"
	"testing"
)

// Функция для генерации доски в виде строки
func generateBoard(size int) string {
	var sb strings.Builder
	for row := 0; row < size; row++ {
		sb.WriteString(generateRow(row, size))
		sb.WriteString("\n") // Добавьте новую строку между рядами
	}
	return sb.String()
}

// Тестирование функции генерации доски
func TestGenerateBoard(t *testing.T) {
	// Ожидаемая доска для размера 3
	expectedBoard := "# #\n # \n# #\n"

	// Генерируем доску
	board := generateBoard(3)

	// Сравниваем результат
	if board != expectedBoard {
		t.Errorf("Expected:\n%sGot:\n%s", expectedBoard, board)
	}

	// Ожидаемая доска для размера 4
	expectedBoard = "# # \n # #\n# # \n # #\n"
	board = generateBoard(4)

	// Сравниваем результат
	if board != expectedBoard {
		t.Errorf("Expected:\n%sGot:\n%s", expectedBoard, board)
	}
}
