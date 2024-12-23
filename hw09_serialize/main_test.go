package main

import (
	"reflect"
	"testing"

	"github.com/tmpcaot/home_work_basic/hw09_serialize/protofile"
	"google.golang.org/protobuf/proto"
)

func TestSerializeBooksProto(t *testing.T) {
	books := []*protofile.Book{
		{Id: 1, Title: "Go Programming", Author: "John Doe", Year: 2023, Size: 300, Rate: 4.5},
		{Id: 2, Title: "Learning Go", Author: "Jane Smith", Year: 2022, Size: 250, Rate: 4.0},
	}

	data, err := serializeBooksProto(books)
	if err != nil {
		t.Fatalf("Failed to serialize books (protobuf): %v", err)
	}

	if len(data) == 0 {
		t.Error("Serialized data (protobuf) is empty")
	}
}

func TestDeserializeBooksProto(t *testing.T) {
	originalBooks := []*protofile.Book{
		{Id: 1, Title: "Go Programming", Author: "John Doe", Year: 2023, Size: 300, Rate: 4.5},
		{Id: 2, Title: "Learning Go", Author: "Jane Smith", Year: 2022, Size: 250, Rate: 4.0},
	}

	data, err := serializeBooksProto(originalBooks)
	if err != nil {
		t.Fatalf("Failed to serialize books (protobuf): %v", err)
	}

	deserializedBooks, err := deserializeBooksProto(data)
	if err != nil {
		t.Fatalf("Failed to deserialize books (protobuf): %v", err)
	}

	if len(deserializedBooks) != len(originalBooks) {
		t.Errorf("Expected %d books (protobuf), got %d", len(originalBooks), len(deserializedBooks))
	}

	for i, book := range deserializedBooks {
		if !proto.Equal(book, originalBooks[i]) {
			t.Errorf("Book %d (protobuf) does not match. Expected %+v, got %+v", i, originalBooks[i], book)
		}
	}
}

func TestDeserializeBooksInvalidDataProto(t *testing.T) {
	data := []byte{0x00, 0x01, 0x02, 0x03}

	_, err := deserializeBooksProto(data)
	if err == nil {
		t.Error("Expected error during deserialization (protobuf), got nil")
	}
}

func TestSerializeBooksJSON(t *testing.T) {
	books := []*protofile.Book{
		{Id: 1, Title: "Go Programming", Author: "John Doe", Year: 2023, Size: 300, Rate: 4.5},
		{Id: 2, Title: "Learning Go", Author: "Jane Smith", Year: 2022, Size: 250, Rate: 4.0},
	}

	data, err := serializeBooksJSON(books)
	if err != nil {
		t.Fatalf("Failed to serialize books (JSON): %v", err)
	}

	if len(data) == 0 {
		t.Error("Serialized data (JSON) is empty")
	}
}

func TestDeserializeBooksJSON(t *testing.T) {
	originalBooks := []*protofile.Book{
		{Id: 1, Title: "Go Programming", Author: "John Doe", Year: 2023, Size: 300, Rate: 4.5},
		{Id: 2, Title: "Learning Go", Author: "Jane Smith", Year: 2022, Size: 250, Rate: 4.0},
	}

	data, err := serializeBooksJSON(originalBooks)
	if err != nil {
		t.Fatalf("Failed to serialize books (JSON): %v", err)
	}

	deserializedBooks, err := deserializeBooksJSON(data)
	if err != nil {
		t.Fatalf("Failed to deserialize books (JSON): %v", err)
	}

	if len(deserializedBooks) != len(originalBooks) {
		t.Errorf("Expected %d books (JSON), got %d", len(originalBooks), len(deserializedBooks))
	}

	for i, book := range deserializedBooks {
		if !reflect.DeepEqual(book, originalBooks[i]) {
			t.Errorf("Book %d (JSON) does not match. Expected %+v, got %+v", i, originalBooks[i], book)
		}
	}
}

func TestDeserializeBooksInvalidDataJSON(t *testing.T) {
	data := []byte{0x00, 0x01, 0x02, 0x03}

	_, err := deserializeBooksJSON(data)
	if err == nil {
		t.Error("Expected error during deserialization (JSON), got nil")
	}
}
