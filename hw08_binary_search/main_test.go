package main

import (
	"sort"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		arr    []int
		target int
		want   int
	}{
		{[]int{3, 1, 4, 5, 2}, 3, 2},  // После сортировки массив будет [1, 2, 3, 4, 5].
		{[]int{3, 1, 4, 5, 2}, 1, 0},  // Элемент 1 будет в индексе 0.
		{[]int{3, 1, 4, 5, 2}, 5, 4},  // Элемент 5 будет в индексе 4.
		{[]int{3, 1, 4, 5, 2}, 6, -1}, // Элемент 6 не найден.
		{[]int{}, 1, -1},              // Пустой массив.
	}

	for _, tt := range tests {
		// Сортировка массива перед тестом.
		sort.Ints(tt.arr)
		got := BinarySearch(tt.arr, tt.target)
		if got != tt.want {
			t.Errorf("BinarySearch(%v, %d) = %d; want %d", tt.arr, tt.target, got, tt.want)
		}
	}
}
