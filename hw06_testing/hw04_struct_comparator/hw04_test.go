package main

import (
	"testing"
)

func TestBookComparator(t *testing.T) {

	book1 := &Book{
		id:     "123-456",
		title:  "Книга 1",
		author: "Автор 1",
		year:   2023,
		size:   301,
		rate:   7.2,
	}
	book2 := &Book{
		id:     "142-365",
		title:  "Книга 2",
		author: "Автор 2",
		year:   2024,
		size:   300,
		rate:   5.2,
	}
	book3 := &Book{
		id:     "543-210",
		title:  "Книга 3",
		author: "Автор 3",
		year:   2023,
		size:   350,
		rate:   8.5,
	}

	tests := []struct {
		name      string
		cmpMethod CompareMethod
		bookA     *Book
		bookB     *Book
		expected  bool
	}{
		{"Сравнение по году: book2 > book1", Year, book2, book1, false},
		{"Сравнение по году: book1 > book2", Year, book1, book2, true},
		{"Сравнение по размеру: book1 > book2", Size, book1, book2, true},
		{"Сравнение по размеру: book2 > book1", Size, book2, book1, false},
		{"Сравнение по рейтингу: book1 > book2", Rate, book1, book2, true},
		{"Сравнение по рейтингу: book2 > book1", Rate, book2, book1, false},
		{"Сравнение по размеру: book3 > book1", Size, book3, book1, true},
		{"Сравнение по году: book3 > book1", Year, book3, book1, false},
		{"Сравнение по рейтингу: book3 > book1", Rate, book3, book1, true},
		{"Сравнение по году: book1 > book3", Year, book1, book3, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comparator := NewComparator(tt.cmpMethod)
			result := comparator.Compare(tt.bookA, tt.bookB)
			if result != tt.expected {
				t.Errorf("Compare(%v, %v) = %v; want %v", tt.bookA, tt.bookB, result, tt.expected)
			}
		})
	}
}
