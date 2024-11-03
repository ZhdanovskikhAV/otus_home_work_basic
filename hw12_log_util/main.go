package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	LogFile    string
	LogLevel   string
	OutputFile string
}

func parseFlags() Config {
	var config Config

	// Определение флагов командной строки.
	flag.StringVar(&config.LogFile, "file", "", "Path to the log file (required).")
	flag.StringVar(&config.LogLevel, "level", "info", "Log level to analyze (default: \"info\").")
	flag.StringVar(&config.OutputFile, "output", "", "Path to output statistics file.")

	flag.Parse()

	// Проверка наличия обязательного флага.
	if config.LogFile == "" {
		config.LogFile = os.Getenv("LOG_ANALYZER_FILE")
	}

	// Обработка переменных окружения.
	if level := os.Getenv("LOG_ANALYZER_LEVEL"); level != "" {
		config.LogLevel = level
	}
	if output := os.Getenv("LOG_ANALYZER_OUTPUT"); output != "" {
		config.OutputFile = output
	}

	// // Вывод считанных флагов.
	// fmt.Println("Parsed configuration:")
	// fmt.Printf("LogFile: %s\n", config.LogFile)
	// fmt.Printf("LogLevel: %s\n", config.LogLevel)
	// fmt.Printf("OutputFile: %s\n", config.OutputFile)

	return config
}

func parseLogFile(filePath string, logLevel string) map[string]int {
	logCounter := make(map[string]int)

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening log file: %v\n", err)
		return logCounter
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Проверяем, соответствует ли строка заданному уровню логирования.
		if strings.Contains(line, logLevel) {
			logCounter[logLevel]++
		}
	}

	if scanErr := scanner.Err(); scanErr != nil {
		fmt.Printf("Error reading log file '%s': %v\n", filePath, scanErr)
	}

	return logCounter
}

func writeOutput(outputPath string, statistics map[string]int) {
	var output *os.File
	var err error

	if outputPath == "" {
		output = os.Stdout
	} else {
		output, err = os.Create(outputPath)
		if err != nil {
			fmt.Printf("Error creating output file: %v\n", err)
			return
		}
		defer output.Close()
	}

	for level, count := range statistics {
		fmt.Fprintf(output, "Log Level: '%s' - Number of occurrences: %d\n", level, count)
	}
}

func main() {
	// Загружаем переменные окружения из файла .env.
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file.")
		return
	}

	config := parseFlags()

	fmt.Printf("LogFile: %s\n", config.LogFile)
	// Проверка наличия лог-файла после парсинга флагов.
	if config.LogFile == "" {
		fmt.Println("Error: Log file path must be specified.")
		os.Exit(1)
	}

	statistics := parseLogFile(config.LogFile, config.LogLevel)
	writeOutput(config.OutputFile, statistics)
}
