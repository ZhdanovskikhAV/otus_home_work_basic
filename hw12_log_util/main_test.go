package main

import (
	"os"
	"testing"
)

// Тест для функции parseLogFile.
func TestParseLogFile(t *testing.T) {
	// Создаем временный лог-файл для теста.
	logFile, err := os.CreateTemp("", "test_log_*.txt")
	if err != nil {
		t.Fatalf("Error creating temp log file: %v", err)
	}
	defer os.Remove(logFile.Name()) // Удаляем файл после теста.

	// Записываем тестовые данные в лог-файл.
	testData := "INFO: This is an info message\nERROR: This is an error message\nINFO: Another info message\n"
	if _, writeErr := logFile.WriteString(testData); writeErr != nil {
		t.Fatalf("Error writing to temp log file: %v", writeErr)
	}
	logFile.Close() // Закрываем файл для чтения.

	// Тестируем функцию parseLogFile.
	expectedCount := 2
	logLevel := "INFO"
	result := parseLogFile(logFile.Name(), logLevel)

	if count, ok := result[logLevel]; !ok || count != expectedCount {
		t.Errorf("Expected %d occurrences for log level '%s', got %d", expectedCount, logLevel, count)
	}
}

// Тест для функции writeOutput.
func TestWriteOutput(t *testing.T) {
	// Создаем временный файл для вывода.
	outputFile, err := os.CreateTemp("", "test_output_*.txt")
	if err != nil {
		t.Fatalf("Error creating temp output file: %v", err)
	}
	defer os.Remove(outputFile.Name()) // Удаляем файл после теста.
	outputFile.Close()                 // Закрываем файл для записи.

	statistics := map[string]int{
		"INFO":  2,
		"ERROR": 1,
	}

	// Тестируем функцию writeOutput.
	writeOutput(outputFile.Name(), statistics)

	// Читаем содержимое выходного файла.
	data, err := os.ReadFile(outputFile.Name())
	if err != nil {
		t.Fatalf("Error reading output file: %v", err)
	}

	expectedOutput := "Log Level: 'INFO' - Number of occurrences: 2\nLog Level: 'ERROR' - Number of occurrences: 1\n"
	if string(data) != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot:\n%s", expectedOutput, string(data))
	}
}
