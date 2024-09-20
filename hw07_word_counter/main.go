package main

import (
	"fmt"
	"regexp"
	"strings"
)

func countWords(text string) map[string]int {
	text = strings.ToLower(text)

	re := regexp.MustCompile(`[^\p{L}\s]+`)
	text = re.ReplaceAllString(text, " ")

	words := strings.Fields(text)

	wordCount := make(map[string]int)

	for _, word := range words {
		wordCount[word]++
	}

	return wordCount
}

func main() {
	text := "Привет, мир! Привет всем. Мир! Это тест, это тест!"
	wordCount := countWords(text)
	fmt.Println(wordCount)
}
