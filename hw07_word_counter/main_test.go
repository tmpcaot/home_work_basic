package main

import (
	"reflect"
	"testing"
)

func TestCountWords(t *testing.T) {
	text := "Hello, world! Hello, Go. Go is great; Go is fun."
	expected := map[string]int{
		"hello": 2,
		"world": 1,
		"go":    3,
		"is":    2,
		"great": 1,
		"fun":   1,
	}
	result := countWords(text)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
