package main

import (
	"encoding/json"
	"fmt"
	"log"

	pb "github.com/ZhdanovskikhAV/otus_home_work_basic/hw09_serialize/bookproto"
	"google.golang.org/protobuf/proto"
)

type Book struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Year   int     `json:"year"`
	Size   int     `json:"size"`
	Rate   float64 `json:"rate"`
}

// Реализация метода String для структуры Book.
func (b *Book) String() string {
	return fmt.Sprintf("ID: %d, Title: %s, Author: %s, Year: %d, Size: %d, Rate: %.2f",
		b.ID, b.Title, b.Author, b.Year, b.Size, b.Rate)
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

	// Десериализуем данные в вспомогательную структуру.
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Присваивание значений полям структуры Book.
	b.ID = aux.ID
	b.Title = aux.Title
	b.Author = aux.Author
	b.Year = aux.Year
	b.Size = aux.Size
	b.Rate = aux.Rate

	return nil
}

// Сериализация слайса книг.
func SerializeBooks(books []*pb.Book) ([]byte, error) {
	bookList := &pb.BookList{Books: books}
	return proto.Marshal(bookList)
}

// Десериализация слайса книг.
func DeserializeBooks(data []byte) ([]*pb.Book, error) {
	bookList := &pb.BookList{}
	if err := proto.Unmarshal(data, bookList); err != nil {
		return nil, err
	}
	return bookList.Books, nil
}

func main() {
	// Создаем экземпляры книг
	book := &pb.Book{
		Id:     1,
		Title:  "Книга 1",
		Author: "Автор 1",
		Year:   1982,
		Size:   142,
		Rate:   8.5,
	}

	book2 := &pb.Book{
		Id:     2,
		Title:  "Книга 2",
		Author: "Автор 2",
		Year:   1990,
		Size:   250,
		Rate:   9.0,
	}

	// Сериализация в JSON
	jsonData, err := json.Marshal(book)
	if err != nil {
		log.Fatalf("Ошибка сериализации JSON: %v", err)
	}
	fmt.Println("Сериализованный JSON:", string(jsonData))

	// Десериализация из JSON
	newBook := &pb.Book{}
	err = json.Unmarshal(jsonData, newBook)
	if err != nil {
		log.Fatalf("Ошибка десериализации JSON: %v", err)
	}
	fmt.Printf("Десериализованная книга: %+v\n", newBook)

	// Проверяем, реализует ли book интерфейс proto.Message
	var msg proto.Message = book // Это безопасно, так как Book удовлетворяет интерфейсу

	// Сериализация с использованием Protobuf
	data, err := proto.Marshal(msg)
	if err != nil {
		log.Fatalf("Ошибка сериализации Protobuf: %v", err)
	}
	fmt.Println("Сериализованные данные Protobuf:", data)

	// Десериализация из Protobuf
	newBookProto := &pb.Book{}
	if unmarshalErr := proto.Unmarshal(data, newBookProto); unmarshalErr != nil {
		err = unmarshalErr
		log.Fatalf("Ошибка десериализации Protobuf: %v", err)
	}
	fmt.Printf("Десериализованная книга из Protobuf: %+v\n", newBookProto)

	// Создаем и сериализуем список книг
	bookList := &pb.BookList{
		Books: []*pb.Book{book, newBookProto},
	}

	listData, err := proto.Marshal(bookList)
	if err != nil {
		log.Fatalf("Ошибка сериализации списка книг: %v", err)
	}
	fmt.Println("Сериализованные данные списка книг:", listData)

	// Десериализация списка книг
	newBookList := &pb.BookList{}
	if unmarshalErr := proto.Unmarshal(listData, newBookList); unmarshalErr != nil {
		err = unmarshalErr
		log.Fatalf("Ошибка десериализации списка книг: %v", err)
	}
	fmt.Printf("Десериализованный список книг: %+v\n", newBookList)

	// Сериализация слайса книг
	books := []*pb.Book{book, book2}
	data, err = SerializeBooks(books)
	if err != nil {
		log.Fatalf("Ошибка сериализации: %v", err)
	}
	log.Println("Сериализованные данные:", data)

	// Десериализация слайса книг
	newBooks, err := DeserializeBooks(data)
	if err != nil {
		log.Fatalf("Ошибка десериализации: %v", err)
	}
	for _, b := range newBooks {
		log.Printf("Десериализованная книга: %+v\n", b)
	}
}
