package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// countWords принимает строку текста и возвращает мапу с количеством упоминаний каждого слова.

func countWords(text string) map[string]int {
	// Создаем мапу для хранения частоты слов.
	wordCount := make(map[string]int)

	// Разделяем текст на слова
	words := strings.Fields(text)
	for _, word := range words {
		// Очищаем слово от пунктуации
		cleanedWord := cleanWord(word)
		if cleanedWord != "" {
			wordCount[cleanedWord]++
		}
	}
	return wordCount
}

// cleanWord удаляет пунктуацию и приводит слово к нижнему регистру.
func cleanWord(word string) string {
	var cleaned strings.Builder
	for _, char := range word {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			cleaned.WriteRune(unicode.ToLower(char))
		}
	}
	return cleaned.String()
}

// Пример использования функции.
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите текст:")
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	wordCount := countWords(text)
	fmt.Println(wordCount)
}
