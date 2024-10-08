package main

import (
	"fmt"
)

func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2
		midValue := arr[mid]

		switch {
		case midValue == target:
			return mid
		case midValue < target:
			left = mid + 1
		case midValue > target:
			right = mid - 1
		}
	}

	return -1
}

func main() {
	data := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	target := 20
	// Проверка, что элемент больше первого и меньше последнего в массиве
	if target > data[0] && target < data[len(data)-1] {
		result := binarySearch(data, target)

		if result != -1 {
			fmt.Printf("Элемент %d найден по индексу %d.\n", target, result)
		} else {
			fmt.Printf("Элемент %d не найден.\n", target)
		}
	} else {
		fmt.Printf("Элемент %d должен быть больше %d и меньше %d.\n", target, data[0], data[len(data)-1])
	}
}
