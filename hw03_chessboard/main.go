package main

import "fmt"

func main() {
	var size int
	fmt.Print("Введите размер шахматной доски: ")

	// Сканируем ввод пользователя.
	_, err := fmt.Scan(&size)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	// Проверяем, чтобы число было больше 0.
	if size <= 0 {
		fmt.Println("Размер должен быть больше 0.")
		return
	}

	board := createChessBoard(size)
	fmt.Println(board)
}

func createChessBoard(size int) string {
	var board string
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if (i+j)%2 == 0 {
				board += "#"
			} else {
				board += " "
			}
		}
		board += "\n"
	}
	return board
}
