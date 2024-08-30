package main

import (
	"fmt"
)

type Book struct {
	id     string
	title  string
	author string
	year   int
	size   int
	rate   float64
}

func (b *Book) SetID(id string) {
	b.id = id
}

func (b *Book) ID() string {
	return b.id
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) Title() string {
	return b.title
}

func (b *Book) SetAuthor(author string) {
	b.author = author
}

func (b *Book) Author() string {
	return b.author
}

func (b *Book) SetYear(year int) {
	b.year = year
}

func (b *Book) Year() int {
	return b.year
}

func (b *Book) SetSize(size int) {
	b.size = size
}

func (b *Book) Size() int {
	return b.size
}

func (b *Book) SetRate(rate float64) {
	b.rate = rate
}

func (b *Book) Rate() float64 {
	return b.rate
}

type CompareMethod int

const (
	Year CompareMethod = 0
	Size CompareMethod = 1
	Rate CompareMethod = 2
)

type BookComparator struct {
	compareMethod CompareMethod
}

func NewComparator(cm CompareMethod) *BookComparator {
	return &BookComparator{compareMethod: cm}
}

func (c BookComparator) Compare(book1, book2 *Book) bool {
	switch {
	case c.compareMethod == Year:
		return book1.Year() > book2.Year()
	case c.compareMethod == Size:
		return book1.Size() > book2.Size()
	case c.compareMethod == Rate:
		return book1.Rate() > book2.Rate()
	default:
		return false
	}
}

func main() {
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
	year := NewComparator(Year)
	size := NewComparator(Size)
	rate := NewComparator(Rate)

	fmt.Println("Первая книга", book1)
	fmt.Println("Вторая книга", book2)
	fmt.Println("compare", year.compareMethod, "(Год книги1 больше книги2):", year.Compare(book1, book2))
	fmt.Println("compare", size.compareMethod, "(Размер книги1 больше книги2):", size.Compare(book1, book2))
	fmt.Println("compare", rate.compareMethod, "(Рейтинг книги1 больше книги2):", rate.Compare(book1, book2))
}
