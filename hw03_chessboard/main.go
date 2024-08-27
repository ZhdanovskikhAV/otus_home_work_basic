package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var size int
	fmt.Println("Введите число клеток в ряду: ")
	_, err := fmt.Scanf("%d\n", &size)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(0)
	}
	for row := 0; row < size; row++ {
		var sb strings.Builder
		a, b := findData(row)
		for column := 0; column < size; column++ {
			if column%2 == 0 {
				sb.WriteRune(a)
			} else {
				sb.WriteRune(b)
			}
		}
		fmt.Println(sb.String())
	}
}

func findData(row int) (rune, rune) {
	if row%2 == 0 {
		return '#', ' '
	}
	return ' ', '#'
}
