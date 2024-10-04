package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Year   int     `json:"year"`
	Size   int     `json:"size"`
	Rate   float64 `json:"rate"`
}

func (b *Book) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":     b.ID,
		"title":  b.Title,
		"author": b.Author,
		"year":   b.Year,
		"size":   b.Size,
		"rate":   b.Rate,
	})
}

func (b *Book) UnmarshalJSON(data []byte) error {
	var aux struct {
		ID     int     `json:"id"`
		Title  string  `json:"title"`
		Author string  `json:"author"`
		Year   int     `json:"year"`
		Size   int     `json:"size"`
		Rate   float64 `json:"rate"`
	}

	// Десериализуем данные в вспомогательную структуру
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Присваивание значений полям структуры Book
	b.ID = aux.ID
	b.Title = aux.Title
	b.Author = aux.Author
	b.Year = aux.Year
	b.Size = aux.Size
	b.Rate = aux.Rate

	return nil
}

func main() {
	book := Book{
		ID:     1,
		Title:  "Книга1",
		Author: "Иванов А.",
		Year:   1982,
		Size:   142,
		Rate:   6.2,
	}

	// Сериализация в JSON
	jsonData, err := json.Marshal(book)
	if err != nil {
		fmt.Println("Ошибка сериализации:", err)
		return
	}
	fmt.Println("Сериализованный JSON:", string(jsonData))

	// Десериализация из JSON
	var newBook Book
	err = json.Unmarshal(jsonData, &newBook)
	if err != nil {
		fmt.Println("Ошибка десериализации:", err)
		return
	}
	fmt.Printf("Десериализованная книга: %+v\n", newBook)

	// Проверка десериализации
	if book == newBook {
		fmt.Println("Десериализация прошла успешно. Объекты совпадают!")
	} else {
		fmt.Println("Десериализация не удалась. Объекты различаются.")
	}
}
