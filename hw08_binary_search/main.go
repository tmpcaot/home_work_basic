package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// BinarySearch выполняет бинарный поиск в отсортированном массиве.
func BinarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		switch {
		case arr[mid] == target:
			return mid // Возвращаем индекс найденного элемента.
		case arr[mid] < target:
			left = mid + 1
		default:
			right = mid - 1
		}
	}
	return -1 // Если элемент не найден, возвращаем -1.
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Ввод массива.
	fmt.Println("Введите элементы массива через пробел:")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	strArr := strings.Fields(input)

	// Преобразование строкового массива в массив целых чисел.
	arr := make([]int, len(strArr))
	for i, str := range strArr {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Ошибка преобразования:", err)
			return
		}
		arr[i] = num
	}

	// Сортировка массива.
	sort.Ints(arr)

	// Ввод элемента для поиска.
	fmt.Println("Введите элемент для поиска:")
	targetInput, _ := reader.ReadString('\n')
	targetInput = strings.TrimSpace(targetInput)
	target, err := strconv.Atoi(targetInput)
	if err != nil {
		fmt.Println("Ошибка преобразования:", err)
		return
	}

	// Поиск элемента.
	index := BinarySearch(arr, target)
	if index != -1 {
		fmt.Printf("Элемент %d найден на индексе %d\n", target, index)
	} else {
		fmt.Printf("Элемент %d не найден\n", target)
	}
}
