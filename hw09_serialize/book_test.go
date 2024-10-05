package main

import (
	"encoding/json"
	"testing"

	pb "github.com/ZhdanovskikhAV/otus_home_work_basic/hw09_serialize/bookproto"
	"google.golang.org/protobuf/proto"
)

// Создание тестовых книг.
func createTestBooks() []*pb.Book {
	return []*pb.Book{
		{
			Id:     1,
			Title:  "Книга 1",
			Author: "Автор 1",
			Year:   1982,
			Size:   142,
			Rate:   8.5,
		},
		{
			Id:     2,
			Title:  "Книга 2",
			Author: "Автор 2",
			Year:   1990,
			Size:   250,
			Rate:   9.0,
		},
	}
}

// Тест для функции SerializeBooks.
func TestSerializeBooks(t *testing.T) {
	books := createTestBooks()

	data, err := SerializeBooks(books)
	if err != nil {
		t.Fatalf("Ошибка сериализации: %v", err)
	}

	// Проверяем, что данные сериализуются
	if len(data) == 0 {
		t.Fatal("Сериализованные данные пусты")
	}
}

// Тест для функции DeserializeBooks.
func TestDeserializeBooks(t *testing.T) {
	books := createTestBooks()

	// Сначала сериализуем книги.
	data, err := SerializeBooks(books)
	if err != nil {
		t.Fatalf("Ошибка сериализации: %v", err)
	}

	// Теперь десериализуем.
	newBooks, err := DeserializeBooks(data)
	if err != nil {
		t.Fatalf("Ошибка десериализации: %v", err)
	}

	// Проверяем, что мы получили правильное количество книг.
	if len(newBooks) != len(books) {
		t.Fatalf("Ожидалось %d книг, но получено %d", len(books), len(newBooks))
	}

	// Проверяем содержимое.
	for i, book := range newBooks {
		if book.Title != books[i].Title {
			t.Errorf("Ожидалось название '%s', но получено '%s'", books[i].Title, book.Title)
		}
		if book.Author != books[i].Author {
			t.Errorf("Ожидался автор '%s', но получено '%s'", books[i].Author, book.Author)
		}
	}
}

// Тест для сериализации и десериализации JSON.
func TestJSONSerialization(t *testing.T) {
	books := createTestBooks()

	// Сериализация в JSON.
	jsonData, err := json.Marshal(books)
	if err != nil {
		t.Fatalf("Ошибка сериализации JSON: %v", err)
	}

	// Десериализация из JSON.
	var newBooks []*pb.Book
	if unmarshalErr := json.Unmarshal(jsonData, &newBooks); unmarshalErr != nil {
		err = unmarshalErr
		t.Fatalf("Ошибка десериализации JSON: %v", err)
	}

	// Проверяем, что мы получили правильное количество книг.
	if len(newBooks) != len(books) {
		t.Fatalf("Ожидалось %d книг, но получено %d", len(books), len(newBooks))
	}

	// Проверяем содержимое.
	for i, book := range newBooks {
		if book.Title != books[i].Title {
			t.Errorf("Ожидалось название '%s', но получено '%s'", books[i].Title, book.Title)
		}
		if book.Author != books[i].Author {
			t.Errorf("Ожидался автор '%s', но получено '%s'", books[i].Author, book.Author)
		}
	}
}

// Тест для сериализации и десериализации Protobuf.
func TestProtobufSerialization(t *testing.T) {
	books := createTestBooks()

	// Сериализация Protobuf.
	data, err := SerializeBooks(books)
	if err != nil {
		t.Fatalf("Ошибка сериализации Protobuf: %v", err)
	}

	// Десериализация Protobuf.
	newBooks, err := DeserializeBooks(data)
	if err != nil {
		t.Fatalf("Ошибка десериализации Protobuf: %v", err)
	}

	// Проверяем, что мы получили правильное количество книг.
	if len(newBooks) != len(books) {
		t.Fatalf("Ожидалось %d книг, но получено %d", len(books), len(newBooks))
	}

	// Проверяем содержимое.
	for i, book := range newBooks {
		if book.Title != books[i].Title {
			t.Errorf("Ожидалось название '%s', но получено '%s'", books[i].Title, book.Title)
		}
		if book.Author != books[i].Author {
			t.Errorf("Ожидался автор '%s', но получено '%s'", books[i].Author, book.Author)
		}
	}
}

// Тест для сериализации и десериализации списка книг.
func TestBookListSerialization(t *testing.T) {
	books := createTestBooks()

	// Создаем и сериализуем список книг
	bookList := &pb.BookList{
		Books: books,
	}

	listData, err := proto.Marshal(bookList)
	if err != nil {
		t.Fatalf("Ошибка сериализации списка книг: %v", err)
	}

	// Десериализация списка книг
	newBookList := &pb.BookList{}
	if unmarshalErr := proto.Unmarshal(listData, newBookList); unmarshalErr != nil {
		err = unmarshalErr
		t.Fatalf("Ошибка десериализации списка книг: %v", err)
	}

	// Проверяем, что мы получили правильное количество книг
	if len(newBookList.Books) != len(bookList.Books) {
		t.Fatalf("Ожидалось %d книг, но получено %d", len(bookList.Books), len(newBookList.Books))
	}

	// Проверяем содержимое
	for i, book := range newBookList.Books {
		if book.Title != bookList.Books[i].Title {
			t.Errorf("Ожидалось название '%s', но получено '%s'", bookList.Books[i].Title, book.Title)
		}
		if book.Author != bookList.Books[i].Author {
			t.Errorf("Ожидался автор '%s', но получено '%s'", bookList.Books[i].Author, book.Author)
		}
	}
}
