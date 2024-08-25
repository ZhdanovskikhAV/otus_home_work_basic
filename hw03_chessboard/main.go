package main

import "fmt"

func main() {
	var size int
	fmt.Println("Введите число клеток в ряду: ")
	fmt.Scanf("%d\n", &size)
	for row := 0; row < size; row++ {
		var (
			str = ""
			a   = ""
			b   = ""
		)

		if row%2 == 0 {
			a = "#"
			b = " "
		} else {
			a = " "
			b = "#"
		}

		for column := 0; column < size; column++ {
			if column%2 == 0 {
				str = str + a

			} else {
				str = str + b

			}

		}
		fmt.Printf("%s\n", str)

	}
}
