package main

import "fmt"

func main() {
	var size int
	fmt.Println("Введите число клеток в ряду: ")
	_, err := fmt.Scanf("%d\n", &size)
	if err != nil {
		fmt.Println("Error:", err)
	}
	for row := 0; row < size; row++ {
		str, a, b := findData(row)
		for column := 0; column < size; column++ {
			if column%2 == 0 {
				str += a
			} else {
				str += b
			}
		}
		fmt.Println(str)
	}
}

func findData(row int) (string, string, string) {
	var (
		str = ""
		a   string
		b   string
	)
	if row%2 == 0 {
		a = "#"
		b = " "
	} else {
		a = " "
		b = "#"
	}
	return str, a, b
}
