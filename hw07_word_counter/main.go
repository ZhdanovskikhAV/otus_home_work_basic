package main

import (
	"fmt"
	"regexp"
	"strings"
)

func countWords(text string) map[string]int {
	text = strings.ToLower(text)

	re := regexp.MustCompile(`[^\p{L}\p{N}\s]+`)
	text = re.ReplaceAllString(text, " ")

	words := strings.Fields(text)

	wordCount := make(map[string]int)

	for _, word := range words {
		wordCount[word]++
	}

	return wordCount
}

func main() {
	text := "Привет, мир! Привет всем. Мир! Это тест, это тест! k8s, log4j, ta4ka/ K8S"
	wordCount := countWords(text)
	fmt.Println(wordCount)
}
