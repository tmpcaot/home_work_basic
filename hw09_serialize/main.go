package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/tmpcaot/home_work_basic/hw09_serialize/protofile"
	"google.golang.org/protobuf/proto"
)

// serializeBooksProto сериализует список книг в формат protobuf.
func serializeBooksProto(books []*protofile.Book) ([]byte, error) {
	bookList := &protofile.BookList{Books: books}
	return proto.Marshal(bookList)
}

// deserializeBooksProto десериализует байты обратно в список книг.
func deserializeBooksProto(data []byte) ([]*protofile.Book, error) {
	var bookList protofile.BookList
	err := proto.Unmarshal(data, &bookList)
	if err != nil {
		return nil, err
	}
	return bookList.Books, nil
}

// serializeBooksJSON сериализует список книг в формат JSON.
func serializeBooksJSON(books []*protofile.Book) ([]byte, error) {
	return json.Marshal(books)
}

// deserializeBooksJSON десериализует байты обратно в список книг.
func deserializeBooksJSON(data []byte) ([]*protofile.Book, error) {
	var books []*protofile.Book
	err := json.Unmarshal(data, &books)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func main() {
	books := []*protofile.Book{
		{Id: 1, Title: "Book 1", Author: "John Doe", Year: 2023, Size: 300, Rate: 4.5},
		{Id: 2, Title: "LBook 2", Author: "Jane Smith", Year: 2022, Size: 250, Rate: 4.0},
	}

	dataProto, err := serializeBooksProto(books)
	if err != nil {
		log.Fatalf("Failed to serialize books (Protobuf): %v", err)
	}

	deserializedBooksProto, err := deserializeBooksProto(dataProto)
	if err != nil {
		log.Fatalf("Failed to deserialize books (Protobuf): %v", err)
	}

	fmt.Println("Deserialized books (Protobuf):")
	for _, book := range deserializedBooksProto {
		fmt.Printf("%+v\n", book)
	}

	dataJSON, err := serializeBooksJSON(books)
	if err != nil {
		log.Fatalf("Failed to serialize books (JSON): %v", err)
	}

	deserializedBooksJSON, err := deserializeBooksJSON(dataJSON)
	if err != nil {
		log.Fatalf("Failed to deserialize books (JSON): %v", err)
	}

	fmt.Println("\nDeserialized books (JSON):")
	for _, book := range deserializedBooksJSON {
		fmt.Printf("%+v\n", book)
	}
}
